package openAI

func (o openAI) NextClient(i int) error {
	i = i % o.index
	o.gpt3 = (*o.gpt3Clients)[i]
	return nil
}
