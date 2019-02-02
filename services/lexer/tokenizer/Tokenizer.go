package tokenizer

import (
	"unicode"
	"unicode/utf8"

	"github.com/InvisibleTech/norvig_lisp/services/lexer/token"
)

// Tokenizer represents current state token parsing.
type Tokenizer struct {
	SourceName string
	Input      string
	Tokens     chan token.Token
	State      TokenizerFunc

	Start int
	Pos   int
	Width int
}

// CreateTokenizer is a factory method.
func CreateTokenizer(name, input string) *Tokenizer {
	l := &Tokenizer{
		SourceName: name,
		Input:      input,
		State:      Begin,
		Tokens:     make(chan token.Token, 3),
	}

	return l
}

func (tokenizer *Tokenizer) InputToEnd() string {
	return tokenizer.Input[tokenizer.Pos:]
}

func (tokenizer *Tokenizer) SkipWhitespace() {
	for {
		ch, _ := tokenizer.Peek()
		if !unicode.IsSpace(ch) {
			break
		}

		tokenizer.Next()
		if ch == token.EOF {
			tokenizer.Emit(token.TokenEOF)
		}
	}
}

func (tokenizer *Tokenizer) Emit(tokenType token.TokenType) {
	tokenizer.Tokens <- token.Token{
		Type:  tokenType,
		Value: tokenizer.Input[tokenizer.Start:tokenizer.Pos]}
	tokenizer.Start = tokenizer.Pos
}

func StrPeek(input string, offset int) (rune, int) {
	if offset >= utf8.RuneCountInString(input) {
		return token.EOF, 0
	}

	return utf8.DecodeRuneInString(input[offset:])
}

func (tokenizer *Tokenizer) TwinPeeks() (rune, rune) {
	rune, width := StrPeek(tokenizer.Input, 0)
	if rune == token.EOF {
		return rune, token.EOF
	}

	reRune, _ := StrPeek(tokenizer.Input, width)

	return rune, reRune
}

func (tokenizer *Tokenizer) Peek() (rune, int) {
	return StrPeek(tokenizer.Input, 0)
}

func (tokenizer *Tokenizer) Next() {
	if tokenizer.Pos >= utf8.RuneCountInString(tokenizer.Input) {
		tokenizer.Width = 0
	}

	_, width := utf8.DecodeRuneInString(tokenizer.Input)

	tokenizer.Width = width
	tokenizer.Pos += tokenizer.Width
}
