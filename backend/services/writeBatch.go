package services

import (
	"context"
	"fmt"
	"strings"
	"zeotap/errors"
	"zeotap/models"
	"zeotap/utils"

	"github.com/ClickHouse/clickhouse-go/v2"
)

// getTableString generates the SQL string required to create a table in ClickHouse.
// The table is created using the `Memory` engine, and the columns are defined based on the batch's column names and their respective types.
//
// Parameters:
// - batch: The batch of data that contains column names and data types.
// - types: A slice of strings representing the data types for each column.
//
// Returns:
// - string: The SQL query string to create the table.
func getTableString(batch models.Batch, types []string) string {
	var columns []string

	// Loop over each column in the batch and create a column definition string
	for index, _ := range batch.ColumnNames {
		column := fmt.Sprintf("`%s` %s", batch.ColumnNames[index], types[index])
		columns = append(columns, column)
	}

	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s) ENGINE = MergeTree ORDER BY tuple()",
		batch.TableName, strings.Join(columns, ", "))
}

// sendBatch inserts a batch of data into the ClickHouse table using a prepared statement.
//
// Parameters:
// - conn: The established connection to ClickHouse.
// - batch: The batch of data to be inserted, including the table name, column names, and the rows to insert.
//
// Returns:
// - int: The number of rows successfully inserted.
// - error: An error if the insertion fails, or nil if successful.
func sendBatch(conn clickhouse.Conn, batch models.Batch) (int, error) {
	fmt.Println("uploading....")

	ctx := context.Background()

	// Prepare the SQL INSERT query for the table
	insertQuery := fmt.Sprintf("INSERT INTO `%s`", batch.TableName)

	// Prepare the batch for insertion
	batchToPush, err := conn.PrepareBatch(ctx, insertQuery)
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to prepare batch: %s", err.Error())
		return 0, errors.MakeInternalServerError(errorMessage)
	}

	var rowsProcessed int

	fmt.Println("Processing rows")

	// Loop through each row in the batch
	for _, row := range batch.Rows {
		// Append the row to the batch for insertion
		err := batchToPush.Append(row...)

		if err != nil {
			errorMessage := fmt.Sprintf("Failed to push row to batch: %s", err.Error())
			return 0, errors.MakeInternalServerError(errorMessage)
		}

		rowsProcessed++
	}

	fmt.Println("Processed rows")

	// Send the batch to ClickHouse
	batchToPush.Send()

	fmt.Println("Uploaded")

	return rowsProcessed, nil
}

// splitAndInsertBatch splits a large batch of rows into smaller sub-batches and inserts them into ClickHouse.
//
// Parameters:
// - conn: The established connection to ClickHouse.
// - batch: The batch of data that needs to be inserted, including the table name, column names, and rows.
//
// Returns:
// - int32: The total number of rows successfully inserted.
// - error: An error if the batch insertion fails, or nil if successful.
func splitAndInsertBatch(conn clickhouse.Conn, batch models.Batch) (int32, error) {
	// Define the maximum size for each batch, and the number of batches
	batchSize := 15000
	batchCount := (len(batch.Rows) + batchSize - 1) / batchSize

	var processedRows int32

	ctx := context.Background()
	types := utils.InferTypes(batch.Rows[0])
	tableString := getTableString(batch, types)

	// Execute the CREATE TABLE query to create the table
	err := conn.Exec(ctx, tableString)
	if err != nil {
		errorMessage := fmt.Sprintf("Could not make new table: %s", err.Error())
		return 0, errors.MakeInternalServerError(errorMessage)
	}

	// Loop through the batch to split it into sub-batches and insert them to DB
	for i := 0; i < batchCount; i++ {
		start := i * batchSize
		end := (i + 1) * batchSize
		if end > len(batch.Rows) {
			end = len(batch.Rows)
		}

		subBatch := models.Batch{
			Rows:           batch.Rows[start:end],
			TableName:      batch.TableName,
			ColumnNames:    batch.ColumnNames,
			ConnectionInfo: batch.ConnectionInfo,
		}

		// Insert the sub-batch and track the number of rows inserted
		rows, err := sendBatch(conn, subBatch)
		if err != nil {
			return processedRows, err
		}
		processedRows += int32(rows)
	}

	return processedRows, nil
}

// WriteBatch writes a batch of rows to a ClickHouse database by establishing a connection
// and splitting the batch into smaller parts for efficient insertion.
//
// Parameters:
// - batch: The batch of data to be inserted, including table name, column names, and rows.
//
// Returns:
// - int32: The total number of rows successfully inserted.
// - error: An error if the insertion fails, or nil if successful.
func WriteBatch(batch models.Batch) (int32, error) {
	fmt.Println("In service to write")
	connectionInfo := batch.ConnectionInfo
	
	//Establish connection to DB
	conn, err := Connect(connectionInfo)
	fmt.Println("Executed connect")
	fmt.Println(conn, err)
	if err != nil {
		return 0, err
	}

	fmt.Println("Connected")

	return splitAndInsertBatch(conn, batch)
}
