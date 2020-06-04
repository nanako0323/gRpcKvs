package models

//Dog dynamodb struct
type Dog struct {
	Name string
	Kind string
}

//TableName ... dynamodb table name
const TableName string = "Dog"
