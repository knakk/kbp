package scanner

import (
	"bufio"
	"bytes"
	"io"
	"unicode/utf8"
)

type tokenType int

const (
	// Special tokens
	TokenIllegal tokenType = iota
	TokenEOF
	TokenEOL

	// N-Triples tokens
	TokenURI
	TokenBNode
	TokenLiteral
	TokenLangTag
	TokenTypeMarker
	TokenDot

	// Triple pattern
	TokenVariable
)

const eof = rune(-1)

type token struct {
	Type tokenType
	Text string
}

func (t tokenType) String() string {
	switch t {
	case TokenIllegal:
		return "Illegal"
	case TokenEOF:
		return "EOF"
	case TokenEOL:
		return "EOL"
	case TokenBNode:
		return "Blank Node"
	case TokenURI:
		return "URI"
	case TokenLiteral:
		return "Literal"
	case TokenLangTag:
		return "Language tag"
	case TokenTypeMarker:
		return "DataType marker"
	case TokenDot:
		return "Dot"
	case TokenVariable:
		return "Variable"
	default:
		panic("token String() TODO")
	}
}

// Scanner is an N-Triples Scanner.
// Comment tokens are not emitted.
type Scanner struct {
	r     *bufio.Reader
	input string // input being scanned

	pos      int  // position in input
	start    int  // start of current token
	unescape bool // true when token needs unescaping

	// Keep track of position in stream for error reporting:
	Row int // line number
	Col int // position in line (in runes, not bytes)

	// Error holds the last encountered error explanation.
	// It is invalidated on next call to Scan()
	Error string
}

// NewStreamingScanner returns a new streaming Scanner over the given reader.
func NewStreamingScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r), Row: 1}
}

// NewScanner returns a new Scanner over the given input string.
func NewScanner(input string) *Scanner {
	return &Scanner{input: input, Row: 1}
}

// token returns the next token in the stream.
func (s *Scanner) Scan() token {
	s.start = s.pos
	s.Error = ""
	addStart, addEnd := 0, 0
	r := s.next()

	for isWhitespace(r) {
		s.ignore()
		r = s.next()
	}

	var tok tokenType
runeSwitch:
	switch r {
	case '\n':
		s.Row++
		s.ignore()
		s.Col = 0
		tok = TokenEOL
	case '<':
		if !s.scanTo('>') {
			s.Error = "unterminated URI"
			break
		}
		addStart = 1
		addEnd = -1
		tok = TokenURI
	case 'a':
		if !isWhitespace(s.peek()) {
			s.Error = "unexpected token"
			s.scanUntilNextToken()
			break runeSwitch
		}
		return token{TokenURI, "http://www.w3.org/1999/02/22-rdf-syntax-ns#type"}
	case '.':
		s.ignore()
		tok = TokenDot
	case '#':
		if !s.scanTo('\n') {
			tok = TokenEOF
			break runeSwitch
		}
		s.ignore()
		tok = TokenEOL
	case '"':
		s.ignore()
		for {
			end := s.scanTo('"')
			if !end {
				s.start-- // we want starting quote in error message
				s.Error = "unterminated Literal"
				break runeSwitch
			}
			if s.pos > 1 && (s.input[s.pos-2] != '\\' || s.input[s.pos-3] == '\\') {
				break
			}
		}
		if s.pos > s.start {
			addEnd = -1
		}
		tok = TokenLiteral
	case '@':
		s.ignore()
		p := s.pos
		s.scanUntilNextToken()
		if p == s.pos {
			s.Error = "invalid language tag"
			break runeSwitch
		}
		tok = TokenLangTag
	case '^':
		if s.peek() != '^' {
			s.Error = "unexpected token"
			s.scanUntilNextToken()
			break runeSwitch
		}
		s.pos++ // consume ^^
		s.ignore()
		tok = TokenTypeMarker
	case '_':
		if s.peek() != ':' {
			s.Error = "unexpected token"
			s.scanUntilNextToken()
			break runeSwitch
		}
		s.scanUntilNextToken()
		addStart = 2
		tok = TokenBNode
	case '?':
		s.scanUntilNextToken()
		addStart = 1
		tok = TokenVariable
	case eof:
		tok = TokenEOF
	case utf8.RuneError:
		s.Error = "illegal UTF-8 encoding"
	default:
		s.scanUntilNextToken()
		s.Error = "unexpected token"
	}

	if s.Error != "" {
		tok = TokenIllegal
	}
	if s.unescape {
		s.unescape = false
		return s.unescaped(tok, s.start+addStart, s.pos+addEnd)
	}
	return token{tok, s.input[s.start+addStart : s.pos+addEnd]}
}

func (s *Scanner) ignore() {
	s.start = s.pos
}

func (s *Scanner) next() rune {
	if s.pos == len(s.input) {
		if s.r == nil {
			// This is not a streaming scanner, and we reached end of input string.
			return eof
		}
		line, err := s.r.ReadString('\n')
		if err != nil && len(line) == 0 {
			return eof
		}
		s.input = line
		s.start = 0
		s.pos = 0
		s.Col++
	}

	r, w := utf8.DecodeRuneInString(s.input[s.pos:])
	s.pos += w
	s.Col++

	return r
}

func (s *Scanner) peek() rune {
	r, _ := utf8.DecodeRuneInString(s.input[s.pos:])
	return r
}

func (s *Scanner) scanTo(stop rune) bool {
	for r := s.next(); r != stop; r = s.next() {
		switch r {
		case eof:
			return false
		case '\n':
			return false
		//case utf8.RuneError:
		//	s.Error = "illegal UTF-8 encoding"
		//	return false
		case '\\':
			s.unescape = true
		}
	}
	return true
}

func (s *Scanner) scanUntilNextToken() {
	for {
		r := s.peek()
		switch r {
		case '<', '"', '.', ';', ',', '\n', ' ', '_', eof, utf8.RuneError:
			return
		default:
			s.next()
		}
	}
}

func (s *Scanner) unescaped(typ tokenType, from, to int) token {
	i := from
	buf := bytes.NewBuffer(make([]byte, 0, to-i))
	for i < to {
		r, w := utf8.DecodeRuneInString(s.input[i:])
		if w == 0 {
			break
		}
		i += w
		if r != '\\' {
			buf.WriteRune(r)
			continue
		}
		var c byte
		switch s.input[i] {
		case 't':
			c = '\t'
		case 'b':
			c = '\b'
		case 'n':
			c = '\n'
		case 'r':
			c = '\r'
		case 'f':
			c = '\f'
		case '"':
			c = '"'
		case '\'':
			c = '\''
		case '\\':
			c = '\\'
		case 'u', 'U':
			d := uint64(0)
			start := i
			digits := 4
			if s.input[i] == 'U' {
				digits = 8
			}
			for i < start+digits {
				i++
				if i == len(s.input) {
					s.Error = "illegal escape sequence"
					return token{TokenIllegal, s.input[start-1 : i]}
				}
				x := uint64(s.input[i])
				if x >= 'a' {
					x -= 'a' - 'A'
				}
				d1 := x - '0'
				if d1 > 9 {
					d1 = 10 + d1 - ('A' - '0')
				}
				if 0 > d1 || d1 > 15 {
					j := i
					for !utf8.FullRuneInString(s.input[j:i]) {
						i++
					}
					s.Error = "illegal escape sequence"
					return token{TokenIllegal, s.input[start-1 : i-1]}
				}
				d = (16 * d) + d1
			}
			buf.WriteRune(rune(d))
			i++
			continue
		default:
			s.Error = "illegal escape sequence"
			return token{TokenIllegal, s.input[i-1 : i+1]}
		}
		buf.WriteByte(c)
		i++
	}

	return token{typ, buf.String()}
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}
