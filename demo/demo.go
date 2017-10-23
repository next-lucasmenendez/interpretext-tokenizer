package main

import (
	"fmt"
	"github.com/lucasmenendez/gotokenizer"
	"path/filepath"
	"io/ioutil"
	"regexp"
)

const (
	wordlist = "./demo/wordlist"
	corpus = "./demo/corpus"
)

func getWordlist() []string {
	var err error
	var root string
	if root, err = filepath.Abs(wordlist); err != nil {
		panic(err)
	}

	var raw []byte
	if raw, err = ioutil.ReadFile(root); err != nil {
		panic(err)
	}

	var lrgx *regexp.Regexp = regexp.MustCompile(`\n`)
	var content string = string(raw)
	return lrgx.Split(content, -1)
}

func getCorpus() string {
	var err error
	var root string
	if root, err = filepath.Abs(corpus); err != nil {
		panic(err)
	}

	var raw []byte
	if raw, err = ioutil.ReadFile(root); err != nil {
		panic(err)
	}

	return string(raw)
}

func main() {
	var input, model string
	input = "Apple makes the case that even its most banal features require a \"proficiency\" in machine learning"
	model = "technology_en"

	var err error
	var tokenizer *gotokenizer.Tokenizer
	if tokenizer, err = gotokenizer.NewTokenizer(model); err != nil {
		panic(err)
	}

	if !tokenizer.Trained() {
		var wl []string = getWordlist()
		var c string = getCorpus()
		tokenizer.Train(wl, c)
	}

	var tokens []string = tokenizer.Tokenize(input)
	fmt.Print("[")
	for _, t := range tokens {
		fmt.Printf("'%s', ", t)
	}
	fmt.Println("]")
}
