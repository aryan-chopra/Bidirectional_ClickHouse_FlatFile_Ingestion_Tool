package services

import (
	"crypto/tls"
	"fmt"
	"strings"
	"zeotap/errors"
	"zeotap/models"

	"github.com/ClickHouse/clickhouse-go/v2"
)

// Connect establishes a connection to a ClickHouse database using the provided connection information.
//
// This function uses the details from the `ConnectionInfo` struct (e.g., host, port, username, password, and database)
// to create a connection to the ClickHouse server. It also performs authentication and handles potential errors.
//
// Parameters:
// - connectionInfo: An instance of models.ConnectionInfo containing the connection details to the ClickHouse server.
//
// Returns:
// - clickhouse.Conn: A connection object to interact with the ClickHouse server, if the connection is successful.
// - error: Returns an error if the connection fails or authentication is incorrect.
//
// Errors:
// - If authentication fails, an AuthError with the message "Wrong username or password" is returned.
// - If the connection fails, a ConnectionError with the relevant error message is returned.
func Connect(connectionInfo models.ConnectionInfo) (clickhouse.Conn, error) {
	fmt.Println("Connecting")
	fmt.Println(connectionInfo)

	// Attempt to establish a connection to the ClickHouse server.
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
		// If authentication error, return specific error
		if strings.Contains(err.Error(), "code: 516") {
			return nil, errors.MakeAuthError("Wrong username or password")
		}

		return nil, errors.MakeConnectionError(err.Error())
	}

	// Verify the server version after connection
	v, err := conn.ServerVersion()
	if err != nil {
		// If authentication error, return specific error
		if strings.Contains(err.Error(), "code: 516") {
			return nil, errors.MakeAuthError("Wrong username or password")
		}

		return nil, errors.MakeConnectionError(err.Error())
	}

	fmt.Println(v.String())

	return conn, nil
}
