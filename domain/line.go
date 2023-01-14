package domain

import "net/http"

type Line interface {
	GetMessage(r *http.Request) (*MessageEvent, error)
	ReplyMessage(reply *MessageEvent) error
	SendDevMessage(msg string) error
}

type MessageEvent struct {
	User  string
	Text  string
	Token string
}
