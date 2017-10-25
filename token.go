package gotokenizer

import (
	"strings"
)

type token struct {
	raw string
	order int
	prev, next *token
}

type sentence []*token

func (s sentence) Len() int {
	return len(s)
}

func (s sentence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sentence) Less(i, j int) bool {
	return s[i].order < s[j].order
}

func (s sentence) get(i int) (r *token) {
	for _, t := range s {
		if t.order == i {
			return t
		}
	}

	return nil
}

func (s sentence) clean() sentence {
	var n sentence

	for _, t := range s {
		if strings.TrimSpace(t.raw) != "" {
			n = append(n, t)
		}
	}

	return n
}