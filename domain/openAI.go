package domain

type OpenAI interface {
	Chat(msg *MessageEvent) (*MessageEvent, error)
}
