package models

type Batch struct {
	ConnectionInfo ConnectionInfo
	TableName      string
	Headers        []Column
	Rows           [][]string
}
