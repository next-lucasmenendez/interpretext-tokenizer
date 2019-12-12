// gotokenizer package provides simple api to tokenize sentences or words.
package gotokenizer

import (
	"regexp"
	"strings"
)

// Sentences function splits text sentences using some regular expressions to
// replace special symbols. Receives text string and return list of strings
// witch contains sentences separated.
func Sentences(s string) (sentences []string) {
	var (
		// Patterns
		numP       = regexp.MustCompile(`([0-9]+)[\.|,|']([0-9]+)`)
		quoutesP   = regexp.MustCompile(`("|'|“|”|’|«|»)`)
		pstopsP    = regexp.MustCompile(`"(.+)\.(.+)"`)
		revpstopsP = regexp.MustCompile(`{p_stop}`)
		dotStop    = regexp.MustCompile(`(.+)\.\.\.\s?([A-Z].*)`)
		stopsP     = regexp.MustCompile(`[^..][!?.]\s`)
		resP       = regexp.MustCompile(`\*\|\*`)
		dotP       = regexp.MustCompile(`{stop}`)

		// Exchangers
		noNum     = numP.ReplaceAllString(s, `$1*|*$2`)
		noQuoutes = quoutesP.ReplaceAllString(noNum, `"`)
		noPstops  = pstopsP.ReplaceAllString(noQuoutes, "\"$1{p_stop}$2\"")
		noDstops  = dotStop.ReplaceAllString(noPstops, "\"$1...{stop}$2\"")
		noStops   = stopsP.ReplaceAllString(noDstops, `$0{stop}`)
		text      = resP.ReplaceAllString(noStops, `.`)
		resText   = revpstopsP.ReplaceAllString(text, `.`)
	)

	sentences = dotP.Split(resText, -1)
	for i, s := range sentences {
		sentences[i] = strings.TrimSpace(s)
	}
	return
}

// Words function tokenizes sentence in individual entities such as words or
// punctuation symbols. Receives sentence string and returns list of string
// tokens.
func Words(s string) (tokens []string) {
	var (
		sym string = `("|\.\.\.|\.|,|:|;|\(|\)|\[|]|{|}|¿|\?|¡|!|[0-9]+\.[0-9]+)`
		reS        = regexp.MustCompile(`\s|\t`)
		reD        = regexp.MustCompile(sym)
		ws         = reS.Split(s, -1)
	)

	for _, w := range ws {
		if reD.MatchString(w) {
			processed := reD.ReplaceAllString(w, ` $1 `)
			temps := reS.Split(processed, -1)

			for _, t := range temps {
				if len(t) > 0 {
					tokens = append(tokens, t)
				}
			}
		} else if len(w) > 0 {
			tokens = append(tokens, w)
		}
	}
	return
}
