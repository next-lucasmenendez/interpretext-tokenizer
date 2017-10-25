package gotokenizer

import (
	"regexp"
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

	var o = 0
	for _, c := range b {
		var f *token = &token{raw: c, order: o, prev: s.get(o)}
		s = append(s, f)
		s.get(o).next = f
		o++
	}

	if !t.m.trained {
		return s
	}

	return s
}

func (t *Tokenizer) regex(s string) (tokens []string) {
	var ws *regexp.Regexp = regexp.MustCompile(`\s|\t`)
	var temp []string = ws.Split(s, -1)

	var cr *regexp.Regexp = regexp.MustCompile(`\s|\t|"|'|\.|,|\:`)
	for _, w := range temp {
		var _w string = cr.ReplaceAllLiteralString(w, "")
		if _w != "" {
			tokens = append(tokens, w)
		}
	}

	return tokens
}