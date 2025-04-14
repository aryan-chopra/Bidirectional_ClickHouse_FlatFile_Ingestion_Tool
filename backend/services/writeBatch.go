package services

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"
	"zeotap/models"
	"zeotap/utils"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func connect(connectionInfo models.ConnectionInfo) (clickhouse.Conn, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", connectionInfo.Host, connectionInfo.Port)},
		Auth: clickhouse.Auth{
			Database: connectionInfo.Database,
			Username: connectionInfo.Username,
			Password: connectionInfo.Password,
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
		},
	})

	if err != nil {
		return nil, err
	}

	v, err := conn.ServerVersion()

	if err != nil {
		return nil, err
	}

	fmt.Println(v.String())

	return conn, nil
}

func getTableString(batch models.Batch, types []string) string {
	var columns []string
	for index, _ := range batch.ColumnNames {
		column := fmt.Sprintf("`%s` %s", batch.ColumnNames[index], types[index])
		columns = append(columns, column)
	}

	return fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s) ENGINE = Memory`,
		batch.TableName, strings.Join(columns, ", ")) 
}

func sendBatch(conn clickhouse.Conn, batch models.Batch) error {
	ctx := context.Background()
	types := utils.InferTypes(batch.Rows[0])
	tableString := getTableString(batch, types)

	err := conn.Exec(ctx, tableString)

	if err != nil {
		return err
	}

	insertQuery := fmt.Sprintf("INSERT INTO `%s`", batch.TableName)

	batchToPush, err := conn.PrepareBatch(ctx, insertQuery)

	if err != nil {
		return err
	}

	for _, row := range batch.Rows {
		err := batchToPush.Append(row...)
		
		if err != nil {
			return err
		}
	}
	
	return batchToPush.Send()
}

func WriteBatch(batch models.Batch) error {
	fmt.Println("In service to write")
	connectionInfo := batch.ConnectionInfo
	conn, err := connect(connectionInfo)
	fmt.Println("Executed connect")
	fmt.Println(conn, err)
	
	if err != nil {
		return err
	}

	fmt.Println("Connected")
	
	return sendBatch(conn, batch)
}
