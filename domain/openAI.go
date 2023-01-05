package domain

type OpenAI interface {
	Chat(text string) (string, error)
}
