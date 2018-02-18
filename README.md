[![GoDoc](https://godoc.org/github.com/lucasmenendez/gotokenizer?status.svg)](https://godoc.org/github.com/lucasmenendez/gotokenizer)
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
    var input string = `Go (often referred to as golang) is a programming language created at Google[12] in 2.009 by Robert Griesemer, Rob Pike, and Ken Thompson[10]. It is a compiled, statically typed language in the tradition of Algol and C, with garbage collection, limited structural typing[3], memory safety features and CSP-style concurrent programming features added.`
    
    var sentences []string = gotokenizer.Sentences(input)
    for _, s := range sentences {
        fmt.Printf("%q\n", gotokenizer.Words(s))
    }
}
````