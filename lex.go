package tinyml

import (
	"io"

	"github.com/tdewolff/parse"
	"github.com/tdewolff/parse/v2/buffer"
)

type TokenType uint32

const (
	ErrorToken TokenType = iota
	UnknownToken
	NewLineToken
	StartTagToken
	EndTagToken
	SecurityTagToken
	TextToken
)

func (tt TokenType) String() string {
	switch tt {
	case ErrorToken:
		return "Error"
	case UnknownToken:
		return "Unknown"
	case StartTagToken:
		return "StartTag"
	case EndTagToken:
		return "EndTag"
	case SecurityTagToken:
		return "SecurityTag"
	case NewLineToken:
		return "NewLine"
	case TextToken:
		return "Text"
	}

	return ""
}

type Lexer struct {
	r     *buffer.Lexer
	inTag bool
	text  []byte
}

// NewLexer returns a new Lexer for a given io.Reader.
func NewLexer(r io.Reader) *Lexer {
	return &Lexer{
		r:     buffer.NewLexer(r),
		inTag: false,
	}
}

// Err returns the error encountered during lexing, this is often io.EOF but also other errors can be returned.
func (l *Lexer) Err() error {
	return l.r.Err()
}

// Restore restores the NULL byte at the end of the buffer.
func (l *Lexer) Restore() {
	l.r.Restore()
}

// Offset returns the current position in the input stream.
func (l *Lexer) Offset() int {
	return l.r.Offset()
}

// Next returns the next Token. It returns ErrorToken when an error was encountered. Using Err() one can retrieve the error message.
func (l *Lexer) Next() (TokenType, []byte) {
	c := l.r.Peek(0)
	// fmt.Println("peek", string(c), c, l.r.Pos())

	switch c {
	case '\n':
		l.r.Move(1)
		for l.cunsumeNewLine() {
		}
		return NewLineToken, l.r.Shift()
	case '[':
		return l.shiftTag()
	default:
		if l.cunsumeText() {
			return TextToken, l.r.Shift()
		}

		if c == 0 && l.r.Err() != nil {
			return ErrorToken, nil
		}
	}

	return UnknownToken, nil
}

func (l *Lexer) cunsumeNewLine() bool {
	c := l.r.Peek(0)
	if c == '\n' {
		l.r.Move(1)
		return true
	}
	return false
}

func (l *Lexer) cunsumeText() bool {
	c := l.r.Peek(0)
	if c == 0 || c == '\n' || c == '[' {
		return false
	}

	for {
		c := l.r.Peek(0)
		if c == 0 && l.r.Err() != nil {
			break
		} else if c == '\n' || c == '[' {
			break
		}

		l.r.Move(1)
	}

	return true
}

func (l *Lexer) shiftTag() (TokenType, []byte) {
	for {
		c := l.r.Peek(0)
		// loop to read end with tag
		if c == ' ' || c == '/' && l.r.Peek(1) == ']' || c == '\t' || c == '\n' || c == '\r' || c == '\f' || c == 0 && l.r.Err() != nil {
			break
		}
		l.r.Move(1)

		if c == ']' {
			break
		}
	}

	l.text = parse.ToLower(l.r.Lexeme()[1:])
	if l.inTag {
		l.inTag = false
		return EndTagToken, l.r.Shift()
	}

	l.inTag = true
	return StartTagToken, l.r.Shift()
}
