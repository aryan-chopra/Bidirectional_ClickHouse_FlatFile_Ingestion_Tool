package services

import (
	"crypto/tls"
	"fmt"
	"zeotap/models"

	"github.com/ClickHouse/clickhouse-go/v2"
)


func Connect(connectionInfo models.ConnectionInfo) (clickhouse.Conn, error) {
	fmt.Println("Connecting")
	fmt.Println(connectionInfo)
	
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
