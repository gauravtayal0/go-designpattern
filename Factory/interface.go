package Factory

type Database interface {
	GetData(string) string
	PutData(string, string)
}