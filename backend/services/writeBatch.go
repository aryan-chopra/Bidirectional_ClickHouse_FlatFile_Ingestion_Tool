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

func getTableString(batch models.Batch, types []string) string {
	var columns []string
	for index, _ := range batch.ColumnNames {
		column := fmt.Sprintf("`%s` %s", batch.ColumnNames[index], types[index])
		columns = append(columns, column)
	}

	return fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s) ENGINE = Memory",
		batch.TableName, strings.Join(columns, ", "))
}

func sendBatch(conn clickhouse.Conn, batch models.Batch) (int, error) {
	fmt.Println("uploading....")

	ctx := context.Background()
	insertQuery := fmt.Sprintf("INSERT INTO `%s`", batch.TableName)

	batchToPush, err := conn.PrepareBatch(ctx, insertQuery)

	if err != nil {
		errorMessage := fmt.Sprintf("Failed to prepare batch: %s", err.Error())
		return 0, errors.MakeInternalServerError(errorMessage)
	}

	var rowsProcessed int

	fmt.Println("Processing rows")

	for _, row := range batch.Rows {
		err := batchToPush.Append(row...)

		if err != nil {
			errorMessage := fmt.Sprintf("Failed to push row to batch: %s", err.Error())
			return 0, errors.MakeInternalServerError(errorMessage)
		}

		rowsProcessed++
	}

	fmt.Println("Processed rows")

	batchToPush.Send()

	fmt.Println("Uploaded")

	return rowsProcessed, nil
}

func splitAndInsertBatch(conn clickhouse.Conn, batch models.Batch) (int32, error) {
	batchSize := 15000
	batchCount := (len(batch.Rows) + batchSize - 1) / batchSize

	var processedRows int32

	ctx := context.Background()
	types := utils.InferTypes(batch.Rows[0])
	tableString := getTableString(batch, types)

	err := conn.Exec(ctx, tableString)

	if err != nil {
		errorMessage := fmt.Sprintf("Could not make new table: %s", err.Error())
		return 0, errors.MakeInternalServerError(errorMessage)
	}

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

		rows, err := sendBatch(conn, subBatch)
		if err != nil {
			return processedRows, err
		}
		processedRows += int32(rows)
	}

	return processedRows, nil
}

func WriteBatch(batch models.Batch) (int32, error) {
	fmt.Println("In service to write")
	connectionInfo := batch.ConnectionInfo
	conn, err := Connect(connectionInfo)
	fmt.Println("Executed connect")
	fmt.Println(conn, err)

	if err != nil {
		return 0, err
	}

	fmt.Println("Connected")

	return splitAndInsertBatch(conn, batch)
}
