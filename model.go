package gotokenizer

import (
	"regexp"
	"fmt"
	"strings"
)

type model struct {
	alias string
	trained bool
}

func initModel(a string) *model {
	return &model{alias: a, trained: false}
}

func (m *model) train(wl []string, c string) {
	var wr *regexp.Regexp = regexp.MustCompile(`\s`)

	var tl map[string][]string = make(map[string][]string, len(wl))
	for _, w := range wl {
		var t []string = wr.Split(w, -1)
		var k string = t[0]
		var f []string = t[1:]
		for _, v := range f {
			var tm string = strings.TrimSpace(v)
			if len(tm) > 0 {
				tl[k] = append(tl[k], tm)
			}
		}
	}

	for k, f := range tl {
		fmt.Println(k, len(f), f)
	}

	return
}

func (m *model) load() bool {
	return false
}
