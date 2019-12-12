[![GoDoc](https://godoc.org/github.com/next-lucasmenendez/interpretext-tokenizer?status.svg)](https://godoc.org/github.com/next-lucasmenendez/interpretext-tokenizer)
[![Report](https://goreportcard.com/badge/github.com/next-lucasmenendez/interpretext-tokenizer)](https://goreportcard.com/report/github.com/next-lucasmenendez/interpretext-tokenizer)

# Interpretext tokenizer
Simple rule-based word/sentence tokenizer.

## Installation
```bash
go install github.com/next-lucasmenendez/interpretext-tokenizer
```

## Demo
````go
package main

import (
    "fmt"
    tokenizer "github.com/next-lucasmenendez/interpretext-tokenizer"
)

func main() {
    var input string = `Go (often referred to as golang) is a programming language created at Google[12] in 2.009 by Robert Griesemer, Rob Pike, and Ken Thompson[10]. It is a compiled, statically typed language in the tradition of Algol and C, with garbage collection, limited structural typing[3], memory safety features and CSP-style concurrent programming features added.`
    
    var sentences []string = tokenizer.Sentences(input)
    for _, s := range sentences {
        fmt.Printf("%q\n", tokenizer.Words(s))
    }
}
````