package services

import (
	"context"
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
		return nil, err
	}
	
	for row.Next() {
		var tableName string
	
		row.Scan(&tableName)
		
		tables = append(tables, tableName)
	}
	
	return tables, nil
}
