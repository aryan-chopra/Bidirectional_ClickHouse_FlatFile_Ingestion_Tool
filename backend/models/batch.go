package models

// Batch represents a data batch that will be written to a database table.
//
// This struct holds information about the database connection, the table to
// which data will be written, the column names, and the rows of data that
// will be inserted.
//
// Fields:
//   - ConnectionInfo: A `ConnectionInfo` struct that contains the necessary
//     information for establishing a connection to the database.
//   - TableName: A string that specifies the name of the table in the database
//     to which the batch will be written.
//   - ColumnNames: A slice of strings representing the names of the columns
//     that the data will be inserted into.
//   - Rows: A slice of slices (`[][]any`) representing the actual data rows
//     that will be inserted. Each inner slice represents a single row of data,
//     with each element corresponding to a value for a column in the table.
type Batch struct {
	ConnectionInfo ConnectionInfo
	TableName      string
	ColumnNames    []string
	Rows           [][]any
}
