package services

import (
	"context"
	"fmt"
	"zeotap/errors"
	"zeotap/models"
)

// FetchTables retrieves the list of table names from a specified ClickHouse database.
//
// This function connects to the ClickHouse database using the provided `ConnectionInfo` in `TableInfo` and executes
// a query to fetch the names of all available tables in the database. The table names are then returned as a slice of strings.
//
// Parameters:
// - tableInfo: An instance of `models.TableInfo` containing the connection information to the ClickHouse database.
//
// Returns:
// - []string: A slice containing the names of all the tables in the database.
// - error: Returns an error if there are issues with the connection, query execution, or any other failure in fetching the tables.
//
// Errors:
// - If there are issues with the connection, a relevant connection error is returned.
// - If the query fails, an InternalServerError is returned.
func FetchTables(tableInfo models.TableInfo) ([]string, error) {
	// Establish a connection to the database
	conn, err := Connect(tableInfo.ConnectionInfo)
	if err != nil {
		return nil, err
	}

	var tables []string
	row, err := conn.Query(context.Background(), "SHOW TABLES")
	if err != nil {
		errorMessage := fmt.Sprintf("Could not fetch tables: %s", err.Error())
		return nil, errors.MakeInternalServerError(errorMessage)
	}

	// Iterate over the query result set and append the table names to the list
	for row.Next() {
		var tableName string

		row.Scan(&tableName)

		tables = append(tables, tableName)
	}

	return tables, nil
}
