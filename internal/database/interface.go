package database

type DB interface {
	ReadMessage(id string) string
}

type Message struct {
	ID      string `db:"id"`
	Message string `db:"content"`
}

var Noop DB = noop{}

type noop struct{}

func (n noop) ReadMessage(id string) string {
	return ""
}
