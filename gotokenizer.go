package gotokenizer

func NewTokenizer(a string) (t *Tokenizer, err error) {
	var m *model = initModel(a)

	return &Tokenizer{m: m}, nil
}