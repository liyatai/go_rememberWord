package modeladmin

type Record struct {
	Name  string
	Time  string
	Level string
}

func (Record) TableName() string {
	return "record"
}
