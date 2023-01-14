package domain

type OpenAI interface {
	Chat(msg *MessageEvent) (*MessageEvent, error)
}

type Gpt3Balancer interface {
}
