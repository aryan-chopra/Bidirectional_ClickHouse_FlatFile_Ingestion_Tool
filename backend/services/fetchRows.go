package services

import (
	"context"
	"fmt"
	"reflect"
	"zeotap/models"
)

func FetchRows(rowInfo models.RowInfo) ([][]any, int, error) {
	conn, err := connect(rowInfo.ConnectionInfo)

	if err != nil {
		return nil, 0, err
	}

	var fetchedRows [][]any

	limit := 4

	query := fmt.Sprintf("SELECT * FROM %s LIMIT %d OFFSET %d", rowInfo.TableName, limit, rowInfo.Start)

	fmt.Println(query)

	rows, err := conn.Query(context.Background(), query)

	if err != nil {
		return nil, 0, err
	}

	var (
		columnTypes = rows.ColumnTypes()
		vars        = make([]any, len(columnTypes))
	)
	for i := range columnTypes {
		vars[i] = reflect.New(columnTypes[i].ScanType()).Interface()
	}

	for rows.Next() {
		err := rows.Scan(vars...)

		if err != nil {
			return nil, 0, err
		}

		scannedItems := make([]any, len(vars))
		for itemIndex, itemValue := range vars {
			scannedItems[itemIndex] = reflect.ValueOf(itemValue).Elem().Interface()
		}

		fetchedRows = append(fetchedRows, scannedItems)
	}

	return fetchedRows, limit, nil
}
