// Provides simple api to tokenize sentences or words.
package gotokenizer

import "regexp"

// Split text sentences using some regular expresions to replace special symbols.
// Receives text string and return list of strings witch contains sentences separated.
func Sentences(s string) (sentences []string) {
	var (
		// Patterns
		numP       = regexp.MustCompile(`([0-9]+)\.([0-9]+)`)
		quoutesP   = regexp.MustCompile(`("|'|“|”|’|«|»)`)
		pstopsP    = regexp.MustCompile(`"(.+)\.(.+)"`)
		revpstopsP = regexp.MustCompile(`{partial_stop}`)
		stopsP     = regexp.MustCompile(`[^..][!?.]\s`)
		resP       = regexp.MustCompile(`\*\|\*`)
		dotP       = regexp.MustCompile(`{stop}`)

		// Exchangers
		noNum     = numP.ReplaceAllString(s, `$1*|*$2`)
		noQuoutes = quoutesP.ReplaceAllString(noNum, `"`)
		noPstops  = pstopsP.ReplaceAllString(noQuoutes, "\"$1{partial_stop}$2\"")
		noStops   = stopsP.ReplaceAllString(noPstops, `$0{stop}`)
		text      = resP.ReplaceAllString(noStops, `.`)
		resText   = revpstopsP.ReplaceAllString(text, `.`)
	)

	sentences = dotP.Split(resText, -1)
	return sentences
}

// Tokenize sentence in individual entities such as words or punctuation symbols.
// Receives sentence string and returns list of string tokens.
func Words(s string) (tokens []string) {
	var (
		rgxS = regexp.MustCompile(`\s|\t`)
		rgxD = regexp.MustCompile(`("|\.\.\.|\.|,|:|\(|\)|\[|]|{|}|¿|\?|¡|!|[0-9]+\.[0-9]+)`)
		ws   = rgxS.Split(s, -1)
	)

	for _, w := range ws {
		if rgxD.MatchString(w) {
			processed := rgxD.ReplaceAllString(w, ` $1 `)
			temps := rgxS.Split(processed, -1)

			for _, t := range temps {
				if len(t) > 0 {
					tokens = append(tokens, t)
				}
			}
		} else if len(w) > 0 {
			tokens = append(tokens, w)
		}
	}

	return tokens
}
