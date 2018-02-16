[![GoDoc](https://img.shields.io/badge/GoDoc-reference-5272B4.svg)](https://godoc.org/github.com/lucasmenendez/gotokenizer)
[![Build Status](https://travis-ci.org/lucasmenendez/gotokenizer.svg?branch=master)](https://travis-ci.org/lucasmenendez/gotokenizer)
[![Report](https://goreportcard.com/badge/github.com/lucasmenendez/gotokenizer)](https://goreportcard.com/report/github.com/lucasmenendez/gotokenizer)

# Gotokenizer
Simple rule-based word/sentence tokenizer.

## Installation
```bash
go install github.com/lucasmenendez/gotokenizer
```

## Demo
````go
package main

import (
    "fmt"
    "github.com/lucasmenendez/gotokenizer"
)

func main() {
    var input string = "LG Mobile has posted just one quarter of profitability over the last two years, that was six months ago during the first quarter of sales of its new flagship, the LG G6, when it eked out a $3.2 million profit. Previous to that, you have to go way back to Q1 2015 for a quarterly profit."
    
    var sentences []string = gotokenizer.Sentences(input)
    for _, s := range sentences {
        fmt.Printf("%q\n", gotokenizer.Words(s))
    }
}
````