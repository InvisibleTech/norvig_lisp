package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Symbol string

type Tokens []string

type Atom struct {
	symbol  Symbol
	integer int64
}

type Exp struct {
	head *Atom
	list []*Exp
}

// Create a new Exp (expression) which makes all things
// an LISP expression.  So an Atom is in effect treated
// like (<atom>)  and so evaluting it will simple be
// the value.
func NewExp(atom *Atom) *Exp {
	return &Exp{
		head: atom,
		list: make([]*Exp, 0, 10),
	}
}

// TODO: Add logic to parse the string to
// a numberic value if it is not just a symbol.
func NewAtom(value string) *Atom {
	return &Atom{
		symbol:  Symbol(value),
		integer: 0,
	}
}

// TODO: unit tests for all the things

func tokenize(chars string) Tokens {
	splitFn := func(c rune) bool {
		return c == ' '
	}
	replacer := strings.NewReplacer("(", " ( ", ")", " ) ")

	return strings.FieldsFunc(replacer.Replace(chars), splitFn)
}

func readFromTokens(tokens Tokens) (*Exp, Tokens, error) {
	if len(tokens) < 1 {
		return nil, tokens, errors.New("unexpected EOF")
	}

	token, restOfTokens := tokens[0], tokens[1:]

	fmt.Printf("token = %#v\n", token)

	if token == "(" {
		fmt.Println("In list")
		head := NewExp(nil)

		for restOfTokens[0] != ")" {
			if exp, returnedTokens, err := readFromTokens(restOfTokens); err == nil {
				head.list = append(head.list, exp)
				restOfTokens = returnedTokens
			}
		}
		restOfTokens = restOfTokens[1:]
		fmt.Println("Exit list")
		return head, restOfTokens, nil

	} else if token == ")" {
		return nil, restOfTokens, errors.New("Unexpected ')'")
	} else {
		return NewExp(NewAtom(token)), restOfTokens, nil
	}
}

func parse(program string) (*Exp, error) {
	exp, _, error := readFromTokens(tokenize(program))
	return exp, error
}

// TODO: Move main out to be a REPL with error reporting.
func main() {
	fmt.Println("Hello, playground")
	code := "(begin (define r 10) (* pi (* r r)))"
	tokens := tokenize(code)
	fmt.Printf("tokens %v \n", tokens)

	parsed, _ := parse(code)
	fmt.Printf("passed %s\n", spew.Sdump(*parsed))
}
