package gotokenizer

import (
	"regexp"
)

type Tokenizer struct {
	m *model
}

func (t *Tokenizer) Train(wl []string, c string) {
	t.m.train(wl, c)
}

func (t *Tokenizer) Trained() (trained bool) {
	return t.m.trained
}

func (t *Tokenizer) Tokenize(s string) (tokens []string) {
	var b []string = t.regex(s)
	if !t.m.trained {
		return b
	}

	return []string{}
}

func (t *Tokenizer) regex(s string) (tokens []string) {
	var ws *regexp.Regexp = regexp.MustCompile(`\s`)

	return ws.Split(s, -1)
}