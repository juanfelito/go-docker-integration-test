package mediators

import "docker-example/internal/database"

type Message interface {
	GetMessage(id string) string
}

type MessageOption func(a *message)

func NewMessageMediator(opts ...MessageOption) Message {
	m := &message{
		db: database.Noop,
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

type message struct {
	db database.DB
}

func (m *message) GetMessage(id string) string {
	return m.db.ReadMessage(id)
}

func WithDatabase(db database.DB) MessageOption {
	return func(a *message) {
		a.db = db
	}
}
