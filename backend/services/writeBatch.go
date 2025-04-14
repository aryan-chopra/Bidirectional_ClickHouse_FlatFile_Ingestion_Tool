package services

import (
	"crypto/tls"
	"fmt"
	"zeotap/models"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func connect(connectionInfo models.ConnectionInfo) error {
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
		return err
	}

	v, err := conn.ServerVersion()

	if err != nil {
		return err
	}

	fmt.Println(v.String())

	return nil
}

func sendBatch(conn clickhouse.Conn, )

func WriteBatch(batch models.Batch) error {
	connectionInfo := batch.ConnectionInfo
	err := connect(connectionInfo)

	if err != nil {
		return err
	}

}
