package gotokenizer

import "testing"

func TestSentences(t *testing.T) {
	var input string = `Go (often referred to as golang) is a programming language created at Google[12] in 2.009 by Robert Griesemer, Rob Pike, and Ken Thompson[10]. It is a compiled, statically typed language in the tradition of Algol and C, with garbage collection, limited structural typing,[3] memory safety features and CSP-style concurrent programming features added.`
	var expected []string = []string{
		`Go (often referred to as golang) is a programming language created at Google[12] in 2.009 by Robert Griesemer, Rob Pike, and Ken Thompson[10]. `,
		`It is a compiled, statically typed language in the tradition of Algol and C, with garbage collection, limited structural typing,[3] memory safety features and CSP-style concurrent programming features added.`,
	}

	var sentences []string = Sentences(input)
	if len(sentences) != len(expected) {
		t.Errorf("Expect %d, got %d", len(expected), len(sentences))
		return
	}

	for i, sentence := range sentences {
		if sentence != expected[i] {
			t.Errorf("Expected '%q', got '%q'", expected[i], sentence)
			return
		}
	}
}

func TestWords(t *testing.T) {
	var input string = `Go (often referred to as golang) is a programming language created at Google[12] in 2.009 by Robert Griesemer, Rob Pike, and Ken Thompson[10].`
	var expected []string = []string{ "Go", "(", "often", "referred", "to", "as", "golang", ")", "is", "a", "programming", "language", "created", "at", "Google", "[", "12", "]", "in", "2.009", "by", "Robert", "Griesemer", ",", "Rob", "Pike", ",", "and", "Ken", "Thompson", "[", "10", "]", ".", }

	var tokens []string = Words(input)
	if len(tokens) != len(expected) {
		t.Errorf("Expect %d, got %d", len(expected), len(tokens))
		return
	}

	for i, token := range tokens {
		if token != expected[i] {
			t.Errorf("Expected '%q', got '%q'", expected[i], token)
			return
		}
	}
}