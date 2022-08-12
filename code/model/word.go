package model

var TableName string

type Word struct {
	Chinese string
	English string
}

func SetName(name string) {
	TableName = name
}
func (Word) TableName() string {
	return TableName
}
