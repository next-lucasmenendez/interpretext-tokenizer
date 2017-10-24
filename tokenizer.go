package gotokenizer

import (
	"regexp"
	"strings"
)

const (
	Start string = "<s>"
	End string = "</s>"
)

type Tokenizer struct {
	m *model
}

func (t *Tokenizer) tokenizeSentences(s string) (tokens []sentence) {
	var numP *regexp.Regexp = regexp.MustCompile(`([0-9]+)\.([0-9]+)`)
	var stopsP *regexp.Regexp = regexp.MustCompile(`[^..][!?.]\s`)
	var resP *regexp.Regexp = regexp.MustCompile(`\*\|\*`)
	var dotP *regexp.Regexp = regexp.MustCompile(`{stop}`)

	var numN string = `$1*|*$2`
	var stopsN string = `$0{stop}`
	var resN = `.`

	var noNum string = numP.ReplaceAllString(s, numN)
	var noStops string = stopsP.ReplaceAllString(noNum, stopsN)
	var text string = resP.ReplaceAllString(noStops, resN)
	var sl []string = dotP.Split(text, -1)

	var res []sentence
	for _, s := range sl {
		res = append(res, t.tokenizeWords(s))
	}
	return res
}

func (t *Tokenizer) tokenizeWords(txt string) (s sentence) {
	var b []string = t.regex(txt)

	s = append(s, &token{raw: Start, order: 0, next: nil, prev: nil})
	for i, c := range b {
		if strings.TrimSpace(c) != "" {
			var f *token = &token{raw: c, order: i + 1, prev: s.get(i)}
			s = append(s, f)
			s.get(i).next = f
		}
	}
	s = append(s, &token{raw: End, order: len(s), next: nil, prev: s.get(len(s) - 1)})

	if !t.m.trained {
		return s
	}

	return nil
}

func (t *Tokenizer) regex(s string) (tokens []string) {
	var ws *regexp.Regexp = regexp.MustCompile(`\s`)

	return ws.Split(s, -1)
}