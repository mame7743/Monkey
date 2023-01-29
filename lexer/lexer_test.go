package lexer

import (
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {

	input := `=+(){},;`
	tests := []struct {
		expectedType     token.TokenType
		expectedLitreral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := new(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedLitreral {
			t.Fatalf("test[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLitreral {
			t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLitreral, tok.Literal)
		}
	}

}
