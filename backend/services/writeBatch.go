package services

import (
	"context"
	"fmt"
	"strings"
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

	return fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s) ENGINE = Memory`,
		batch.TableName, strings.Join(columns, ", ")) 
}

func sendBatch(conn clickhouse.Conn, batch models.Batch) (int, error) {
	ctx := context.Background()
	types := utils.InferTypes(batch.Rows[0])
	tableString := getTableString(batch, types)

	err := conn.Exec(ctx, tableString)

	if err != nil {
		return 0, err
	}

	insertQuery := fmt.Sprintf("INSERT INTO `%s`", batch.TableName)

	batchToPush, err := conn.PrepareBatch(ctx, insertQuery)

	if err != nil {
		return 0, err
	}

	var rowsProcessed int
	
	for _, row := range batch.Rows {
		err := batchToPush.Append(row...)
		
		if err != nil {
			return 0, err
		}
		
		rowsProcessed++
	}
	
	batchToPush.Send()
	
	return rowsProcessed, nil
}

func WriteBatch(batch models.Batch) (int, error) {
	fmt.Println("In service to write")
	connectionInfo := batch.ConnectionInfo
	conn, err := Connect(connectionInfo)
	fmt.Println("Executed connect")
	fmt.Println(conn, err)
	
	if err != nil {
		return 0, err
	}

	fmt.Println("Connected")
	
	return sendBatch(conn, batch)
}
