package movalid

type Method uint8

const (
	NOTNULL Method = iota
	EMAIL
	NUMBER
	PHONE
	FILE
	NULL
)

type Param struct {
	Name    string
	Methods []Method
}
