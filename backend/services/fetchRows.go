package services

import (
	"context"
	"fmt"
	"reflect"
	"zeotap/errors"
	"zeotap/models"
)

// FetchRows fetches rows from a specified table in the ClickHouse database, starting at a given offset.
//
// This function connects to the ClickHouse database using the provided `ConnectionInfo` in `RowInfo`, constructs a SQL
// query to fetch rows from the specified table, and processes the result set into a slice of rows. It also handles
// potential errors during the connection, query execution, and row processing.
//
// Parameters:
//   - rowInfo: An instance of `models.RowInfo` containing the connection information, table name, and starting row
//     offset for the query.
//
// Returns:
// - []string: A slice of column names for the queried table.
// - [][]any: A 2D slice containing the rows returned by the query. Each inner slice represents a single row.
// - int: The limit of rows fetched per query.
// - error: Returns an error if any issues occur during the connection, querying, or row processing stages.
//
// Errors:
// - If there are issues with the connection, a relevant connection error is returned.
// - If the SQL query execution fails, an InternalServerError is returned.
// - If row scanning fails, an InternalServerError is returned.
func FetchRows(rowInfo models.RowInfo) ([]string, [][]any, int, error) {
	// Establish a connection to the database
	conn, err := Connect(rowInfo.ConnectionInfo)
	if err != nil {
		return nil, nil, 0, err
	}

	var fetchedRows [][]any

	// Define the limit for the number of rows to fetch per query
	limit := 15000

	// Construct the SQL query with table name, limit, and offset
	query := fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", rowInfo.TableName, limit, rowInfo.Start)

	fmt.Println(query)

	rows, err := conn.Query(context.Background(), query)

	if err != nil {
		errorMessage := fmt.Sprintf("Failed to fetch rows: %s", err.Error())
		return nil, nil, 0, errors.MakeInternalServerError(errorMessage)
	}

	// Retrieve column types and initialize variables for scanning
	var (
		columnTypes = rows.ColumnTypes()
		vars        = make([]any, len(columnTypes))
	)
	// Initialize reflection-based variables for scanning rows
	for i := range columnTypes {
		vars[i] = reflect.New(columnTypes[i].ScanType()).Interface()
	}

	// Iterate over the result set and scan each row into `vars`
	for rows.Next() {
		err := rows.Scan(vars...)

		if err != nil {
			errorMessage := fmt.Sprintf("Failed to insert rows to struct: %s", err.Error())
			return nil, nil, 0, errors.MakeInternalServerError(errorMessage)
		}

		// Process the scanned values into a slice of values
		scannedItems := make([]any, len(vars))
		for itemIndex, itemValue := range vars {
			scannedItems[itemIndex] = reflect.ValueOf(itemValue).Elem().Interface()
		}

		// Append the processed row to the result set
		fetchedRows = append(fetchedRows, scannedItems)
	}

	return rows.Columns(), fetchedRows, limit, nil
}
