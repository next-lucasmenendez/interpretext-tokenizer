package gotokenizer

import (
	"path"
	"regexp"
	"strings"
	"io/ioutil"
	"path/filepath"

	"fmt"
)

const (
	modelPath string = "./models"
	stopwords string = "stopwords"
	corpus string = "corpus"
)

type model struct {
	alias string
	trained bool
	stopwords []string
}

func initModel(a string) (m *model) {
	m = &model{alias: a, trained: false}
	m.trained = m.load()

	return m
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
	var err error
	var root string
	if root, err = filepath.Abs(modelPath); err != nil {
		return false
	}

	var mPath string = path.Join(root, m.alias)
	var swPath string = path.Join(mPath, stopwords)

	var rawSw []byte
	if rawSw, err = ioutil.ReadFile(swPath); err != nil {
		return false
	}

	var lb *regexp.Regexp = regexp.MustCompile(`\n`)
	m.stopwords = lb.Split(string(rawSw), -1)

	return true
}

func (m *model) isStopword(t *token) bool {
	for _, w := range m.stopwords {
		if w == strings.ToLower(t.raw) {
			return true
		}
	}

	return false
}
