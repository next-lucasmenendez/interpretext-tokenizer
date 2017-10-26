package gotokenizer

import "regexp"

func Sentences(s string) (sentences [][]string) {
	var (
		numP   = regexp.MustCompile(`([0-9]+)\.([0-9]+)`)
		stopsP = regexp.MustCompile(`[^..][!?.]\s`)
		resP   = regexp.MustCompile(`\*\|\*`)
		dotP   = regexp.MustCompile(`{stop}`)

		noNum   = numP.ReplaceAllString(s, `$1*|*$2`)
		noStops = stopsP.ReplaceAllString(noNum, `$0{stop}`)
		text    = resP.ReplaceAllString(noStops, `.`)
		sl      = dotP.Split(text, -1)
	)

	for _, s := range sl {
		sentences = append(sentences, Words(s))
	}

	return sentences
}

func Words(s string) (tokens []string) {
	var (
		ws = regexp.MustCompile(`\s|\t`)
		cr = regexp.MustCompile(`\s|\t|"|'|\.|,|\:`)

		temp = ws.Split(s, -1)
	)

	for _, w := range temp {
		var _w string = cr.ReplaceAllLiteralString(w, "")
		if _w != "" {
			tokens = append(tokens, w)
		}
	}

	return tokens
}
