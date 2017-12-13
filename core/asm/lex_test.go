package asm

import "testing"

func lexAll(src string) []token {
	ch := Lex("test.asm", []byte(src), false)

	var tokens []token
	for i := range ch {
		tokens = append(tokens, i)
	}
	return tokens
}

func TestComment(t *testing.T) {
	tokens := lexAll(";; this is a comment")
	if len(tokens) != 2 { // {new line, EOF}
		t.Error("expected no tokens")
	}
}
