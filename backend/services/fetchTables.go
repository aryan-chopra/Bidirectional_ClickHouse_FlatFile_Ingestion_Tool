package services

import (
	"context"
	"fmt"
	"zeotap/errors"
	"zeotap/models"
)

func FetchTables(tableInfo models.TableInfo) ([]string, error) {
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
	
	for row.Next() {
		var tableName string
	
		row.Scan(&tableName)
		
		tables = append(tables, tableName)
	}
	
	return tables, nil
}
