package lexer

import (
	"bufio"
	"fmt"
	"io"
)

type Lexer interface {
	Token() (*Token, error)
}

type Token struct {
	Type TokenType
	Raw  string
}

type TokenType int

const (
	IdentifierToken TokenType = iota
	IntegerToken
	StarToken
)

type lexer struct {
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) Lexer {
	return &lexer{
		reader: bufio.NewReader(reader),
	}
}

func (l *lexer) Token() (*Token, error) {
	r, err := l.reader.ReadRune()
	if err != nil {
		return err
	}
	if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {
		l.reader.UnreadRune()
		return l.identifier()
	} else if isWhitespace(r) {
		l.reader.UnreadRune()
		return l.whitespace()
	} else if '0' <= r && r <= '9' {
		l.reader.UnreadRune()
		return l.number()
	}
	return nil, fmt.Errorf("could not lex")
}

func (l *lexer) identifier() (*Token, error) {
	return nil, fmt.Errorf("unimplemented identifier")
}

func (l *lexer) whitespace() (*Token, error) {
	return nil, fmt.Errorf("unimplemented whitespace")
}

func (l *lexer) number() (*Token, error) {
	result := Token{
		Type: NumberToken,
		Raw:  "",
	}

	for {
		r, err := l.reader.ReadRune()
		if err != nil {
			return nil, err
		}
		if ('0' <= r && r <= '9') || r == '.' || r == '_' {
			result.Raw += r
		} else if isWhitespace(r) {
			l.reader.UnreadRune()
			break
		} else {
			result.Raw += r
			return nil, fmt.Errorf("expected number, found %q", result.Raw)
		}
	}
	return &result, nil
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\r' || r == '\t'
}
