package token

type TokenType string

const (
	ILLEGAL    = "ILLEGAL"
	EOF        = "EOF"
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT" // 1343456
	LPAREN     = "("
	RPAREN     = ")"
)

type Token struct {
	Type    TokenType
	Literal string
}
