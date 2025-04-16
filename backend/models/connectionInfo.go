package models

// ConnectionInfo holds the necessary details to establish a connection
// to a database.
//
// This struct is used to provide the connection parameters needed to
// connect to a specific database instance. It includes the host, port,
// database name, and authentication credentials (username and password).
//
// Fields:
// - Host: A string representing the hostname or IP address of the database server.
// - Port: An integer representing the port on which the database server is running.
// - Database: A string specifying the name of the database to connect to.
// - Username: A string representing the username used to authenticate with the database.
// - Password: A string representing the password associated with the username for authentication.
type ConnectionInfo struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}
