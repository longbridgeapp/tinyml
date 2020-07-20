package tinyml

import (
	"bytes"
	"io"
	"testing"

	"github.com/tdewolff/test"
)

type TTs []TokenType

func TestTokens(t *testing.T) {
	var tokenTests = []struct {
		css     string
		ttypes  []TokenType
		lexemes []string
	}{
		{"Hello world", TTs{TextToken}, []string{"Hello world"}},
		{"Hello world\n\nThis new next line", TTs{TextToken, NewLineToken, TextToken}, []string{"Hello world", "\n\n", "This new next line"}},
		{"\nHello world\n\nThis is new line", TTs{NewLineToken, TextToken, NewLineToken, TextToken}, []string{"\n", "Hello world", "\n\n", "This is new line"}},
		{"传言[st]ST/US/BABA#阿里巴巴.US[/st]将在港股上市", TTs{TextToken, StartTagToken, TextToken, EndTagToken, TextToken}, []string{"传言", "[st]", "ST/US/BABA#阿里巴巴.US", "[/st]", "将在港股上市"}},
		{"[ST]ST/US/BABA#阿里巴巴.US[/st] 将在港股上市", TTs{StartTagToken, TextToken, EndTagToken, TextToken}, []string{"[st]", "ST/US/BABA#阿里巴巴.US", "[/st]", " 将在港股上市"}},
	}

	for _, tt := range tokenTests {
		t.Run(tt.css, func(t *testing.T) {
			l := NewLexer(bytes.NewBufferString(tt.css))
			i := 0
			tokens := []TokenType{}
			lexemes := []string{}
			for {
				token, data := l.Next()
				// fmt.Println("token:", token, string(data))
				if token == ErrorToken {
					test.T(t, l.Err(), io.EOF)
					break
				}

				tokens = append(tokens, token)
				lexemes = append(lexemes, string(data))
				i++
			}

			test.T(t, tokens, tt.ttypes, "token types must match")
			test.T(t, lexemes, tt.lexemes, "token data must match")
		})
	}
}
