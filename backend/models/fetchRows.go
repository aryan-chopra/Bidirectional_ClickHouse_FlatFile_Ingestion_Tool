package models

type FetchRows struct {
	ConnectionInfo ConnectionInfo
	TableName      string
	start          int
}
