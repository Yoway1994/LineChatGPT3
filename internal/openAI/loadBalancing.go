package openAI

import "go.uber.org/zap"

var index = 0

func (o openAI) NextClient(i int) error {
	zap.S().Info("current index: ", index)
	o.gpt3 = (*o.gpt3Clients)[(i % o.num)]
	return nil
}
