package openAI

func (o openAI) NextClient(i int) error {
	o.gpt3 = (*o.gpt3Clients)[i]
	return nil
}
