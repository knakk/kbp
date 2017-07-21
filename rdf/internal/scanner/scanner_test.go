package scanner

import "testing"

func collectTokens(s *Scanner) []token {
	tokens := []token{}
	for {
		tk := s.Scan()
		if tk.Type == TokenEOF {
			break
		}
		tokens = append(tokens, token{tk.Type, tk.Text})

	}
	return tokens
}

func equalTokens(a, b []token) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if a[k].Type != b[k].Type {
			return false
		}
		if a[k].Text != b[k].Text {
			return false
		}
	}
	return true
}

func TestScanTokens(t *testing.T) {
	tests := []struct {
		input string
		want  []token
	}{
		{
			"",
			nil,
		},
		{
			" \t ",
			nil,
		},
		{
			"\n",
			[]token{
				{TokenEOL, ""},
			},
		},
		{
			"<a>",
			[]token{
				{TokenURI, "a"},
			},
		},
		{
			" <http://xyz/æøå.123> ",
			[]token{
				{TokenURI, "http://xyz/æøå.123"},
			},
		},
		{
			`""`,
			[]token{
				{TokenLiteral, ""},
			},
		},
		{
			`"a"`, []token{
				{TokenLiteral, "a"},
			},
		},
		{
			`"a"`,
			[]token{{TokenLiteral, "a"}},
		},
		{
			`"100"^^<int>`,
			[]token{
				{TokenLiteral, "100"},
				{TokenTypeMarker, ""},
				{TokenURI, "int"},
			},
		},
		{
			`"hei"@nb-No `,
			[]token{
				{TokenLiteral, "hei"},
				{TokenLangTag, "nb-No"},
			},
		},
		{
			`"\\" "\""`,
			[]token{
				{TokenLiteral, "\\"},
				{TokenLiteral, "\""},
			},
		},
		{
			`"\t\r\n\f\b\\\u00b7\u00B7\U000000b7\U000000B7"`,
			[]token{
				{TokenLiteral, "\t\r\n\f\b\\····"},
			},
		},
		{
			`"abc\tæøå"`,
			[]token{
				{TokenLiteral, "abc\tæøå"},
			},
		},
		{
			`"line #1\nline #2"`,
			[]token{
				{TokenLiteral, "line #1\nline #2"},
			},
		},
		{
			"<a>^^<f>", []token{
				{TokenURI, "a"},
				{TokenTypeMarker, ""},
				{TokenURI, "f"},
			},
		},
		{
			`"a" 	"b"`,
			[]token{
				{TokenLiteral, "a"},
				{TokenLiteral, "b"},
			},
		},
		{
			"<a><b> <c> .\n",
			[]token{
				{TokenURI, "a"},
				{TokenURI, "b"},
				{TokenURI, "c"},
				{TokenDot, ""},
				{TokenEOL, ""},
			},
		},
		{
			"<a> # a comment <b>\n<c>",
			[]token{
				{TokenURI, "a"},
				{TokenEOL, ""},
				{TokenURI, "c"},
			},
		},
		{
			`<http://example/æøå> <http://example/禅> "\"\\\r\n Здра́вствуйте	☺" .`,
			[]token{
				{TokenURI, "http://example/æøå"},
				{TokenURI, "http://example/禅"},
				{TokenLiteral, "\"\\\r\n Здра́вствуйте\t☺"},
				{TokenDot, ""},
			},
		},
		{
			"<s> a ",
			[]token{
				{TokenURI, "s"},
				{TokenURI, "http://www.w3.org/1999/02/22-rdf-syntax-ns#type"},
			},
		},
		{
			"_:a <b> _:c .",
			[]token{
				{TokenBNode, "a"},
				{TokenURI, "b"},
				{TokenBNode, "c"},
				{TokenDot, ""},
			},
		},
		{
			"?s ?pred <o> .",
			[]token{
				{TokenVariable, "s"},
				{TokenVariable, "pred"},
				{TokenURI, "o"},
				{TokenDot, ""},
			},
		},
		{
			"?s ?p [ <p> <o> ; <p2> <o2> ] .",
			[]token{
				{TokenVariable, "s"},
				{TokenVariable, "p"},
				{TokenBracketStart, ""},
				{TokenURI, "p"},
				{TokenURI, "o"},
				{TokenSemicolon, ""},
				{TokenURI, "p2"},
				{TokenURI, "o2"},
				{TokenBracketEnd, ""},
				{TokenDot, ""},
			},
		},
	}

	for _, test := range tests {
		scanner := NewScanner(test.input)
		if tokens := collectTokens(scanner); !equalTokens(tokens, test.want) {
			t.Errorf("scanning %q got %v; want %v", test.input, tokens, test.want)
		}
	}
}

func TestScanErrors(t *testing.T) {
	tests := []struct {
		input    string
		errWant  string
		textWant string
	}{
		{"<a", "unterminated URI", "<a"},
		{`"xy\z"`, "illegal escape sequence", `\z`},
		{`"\u00F ij"`, "illegal escape sequence", `\u00F`},
		{`"\u123"`, "illegal escape sequence", `\u123`},
		{`"\u123ø.a"`, "illegal escape sequence", "\\u123\xc3"},
		{`"hei`, "unterminated Literal", `"hei`},
		{`"hei\n`, "unterminated Literal", "\"hei\n"},
		{`abc <a>`, "unexpected token", "abc"},
		{`^a b`, "unexpected token", "^a"},
		{`@ <a>`, "invalid language tag", ""},
		{"abc", "unexpected token", "abc"},
		{"_a", "unexpected token", "_a"},
	}

	for _, test := range tests {
		scanner := NewScanner(test.input)
		tok := scanner.Scan()
		if tok.Type != TokenIllegal {
			t.Errorf("scanning %q got %v; want %v", test.input, tok.Type, TokenIllegal)
		}
		if scanner.Error != test.errWant || tok.Text != test.textWant {
			t.Errorf("scanning %q got %q, %q; want %q, %q",
				test.input, scanner.Error, tok.Text, test.errWant, test.textWant)
		}
	}
}

/*
func BenchmarkScanner(b *testing.B) {
	var tok token
	small := datasetAsByte(b, "small")
	for n := 0; n < b.N; n++ {
		s := newScanner(bytes.NewBuffer(small))
		for {
			tok = s.Scan()
			if tok.Type == TokenEOF {
				break
			}
		}
	}
}*/
