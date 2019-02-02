package token

type TokenType int

const (
	// TokenError for those times when you fail.
	TokenError TokenType = iota
	// TokenEOF demarks end of input
	TokenEOF
	// TokenLeftParen is the start of a list of TokenSymbol and TokenNm=umber
	TokenLeftParen
	// TokenRightParen closes the list
	TokenRightParen
	//TokenSymbol is any list of chars, symbols and is usually meant to
	// be a variable name or if the first element of an eval list a
	// string representing a built in or compiled operation.
	TokenSymbol
	// TokenNumber is distinct from a symbol since its symbol is also
	// its value. Too many to enumerate so we treat them not like
	// variable names of symbols to which we assign the value of itself.
	TokenNumber
)
