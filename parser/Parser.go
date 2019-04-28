package parser

type Integer struct {
	value int64
}

type Symbol struct {
	value string
}

type AtomExpression struct {
	value interface{}
}

type ListNode struct {
	element interface{}
}

type ListExpression struct {
	head *ListNode
}

type QuutedListExpressoon struct {
	head *ListNode
}

func NewInteger(intValue int64) Integer {
	return Integer{
		value: intValue,
	}
}

func NewSymbol(symbolValue string) Symbol {
	return Symbol{
		value: symbolValue,
	}
}

func NewIntegerAtom(value int64) AtomExpression {
	return AtomExpression{
		value: NewInteger(value),
	}
}

func NewSymbolAtom(value string) AtomExpression {
	return AtomExpression{
		value: NewSymbol(value),
	}
}
