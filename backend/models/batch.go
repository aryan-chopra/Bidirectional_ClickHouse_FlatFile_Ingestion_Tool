package models

type Batch struct {
	ConnectionInfo ConnectionInfo
	TableName      string
	ColumnNames    []string
	Rows           [][]any
}
