# norvig_lisp
A Go lang implementation of Norvig's LISP subset originally done in Python.
## Why
There's lots of `LISP` parsers and this one is yet another one.  For me this is a chance to
learn `Go` in context. It is also a warmup before I attempt the actual `ANSI Common LISP` spec and anything I can glean from Christian Queinnec's *Lisp In Small Pieces*.

This is a side project, for pleasure and nostalgia.  When I think of `LISP` I think of *The Eternal Golden Braid* and that first day I sat on a bench on a hill reading about Achilles and the Tortoise.  If this project lands anywhere I would like it to be, eventually, a compiler for LISP that supports code that creates code and applies JIT compilation to the newly added code.  I would like a fast, standard LISP.  However, mostly I want to have fun with it and share some narration of what happened along the way.

## Approach
After some overly complicated attempts I decided to do the simplest thing possible.  Code the the Go following Norvig's original Python design to the closest degree possible.

The outcome may be somewhat questionable Go, but I wanted to figure out things that type systems for Python, Java C, and Scala can do (inheritance, union types, etc.) that are not allowed by design in Go.  For this reason the Atom struct is kind of shady in that it has a symbol and an integer value.  I decided to make my main "union type" Exp to be an Atom or a list (slice) of Exp pointers.  So in the end when I code the eval processing, it will use the Exp.head to be a value (an Atom that has no list of arguments) or Exp.head will be how we look up in the Env (environment) the definition of the function to use.

## Pending Work
* REPL
* Unit tests
* Env (binding environment)
* Eval (evaluator for the AST composed of Exps)

