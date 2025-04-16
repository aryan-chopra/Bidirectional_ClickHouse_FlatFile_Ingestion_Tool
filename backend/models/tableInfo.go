package models

// TableInfo holds the connection details required to interact with the database
// and retrieve information about the available tables.
//
// This struct is used to specify the connection details needed to query metadata
// about the database tables.
//
// Fields:
// - ConnectionInfo: An instance of ConnectionInfo that contains the connection parameters for the database.
type TableInfo struct {
	ConnectionInfo ConnectionInfo
}
