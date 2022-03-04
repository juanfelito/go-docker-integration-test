package controllers

import (
	"docker-example/pkg/mediators"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func NewMessagesController(mediator mediators.Message) Messages {
	return &messagesController{
		mediator: mediator,
	}
}

type Messages interface {
	GetMessage(w http.ResponseWriter, r *http.Request)
}

type messagesController struct {
	mediator mediators.Message
}

func (m *messagesController) GetMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := m.mediator.GetMessage(vars["id"])
	fmt.Fprintf(w, "A message with the id %v: %v\n", vars["id"], message)
}
