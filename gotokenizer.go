package gotokenizer

import "sort"



func NewTokenizer(a string) (t *Tokenizer, err error) {
	var m *model = initModel(a)

	return &Tokenizer{m: m}, nil
}

func (t *Tokenizer) Train(wl []string, c string) {
	t.m.train(wl, c)
}

func (t *Tokenizer) Trained() (trained bool) {
	return t.m.trained
}

func (g *Tokenizer) Sentences(text string) ([][]string) {
	var sl []sentence = g.tokenizeSentences(text)

	var tt [][]string = make([][]string, len(sl))
	for i, s := range sl {
		sort.Sort(s)
		for _, t := range s {
			tt[i] = append(tt[i], t.raw)
		}
	}
	return tt
}

func (t *Tokenizer) Words(text string) ([]string) {
	var s sentence = t.tokenizeWords(text)
	var tl []string = make([]string, len(s))

	sort.Sort(s)
	for _, t := range s {
		tl = append(tl, t.raw)
	}
	return tl
}