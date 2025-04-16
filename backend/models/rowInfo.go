package models

// RowInfo holds the details necessary to retrieve rows from a specific database table.
//
// This struct is used to specify the connection details, target table, and the starting point
// (offset) for fetching rows from the database.
//
// Fields:
// - ConnectionInfo: An instance of ConnectionInfo that contains the connection parameters to the database.
// - TableName: A string representing the name of the database table from which rows need to be fetched.
// - Start: An integer indicating the starting row index (offset) for fetching rows. It is typically used for pagination.
type RowInfo struct {
	ConnectionInfo ConnectionInfo
	TableName      string
	Start          int
}
