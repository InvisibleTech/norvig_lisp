# norvig_lisp
A Go lang implementation of Norvig's LISP subset originally done in Python.
## Why
There's lots of `LISP` parsers and this one is yet another one.  For me this is a chance to 
learn `Go` in context and build on what I learned in [coding along with this blog](https://adampresley.github.io/2015/04/12/writing-a-lexer-and-parser-in-go-part-1.html).
It is also a warmup before I attempt the actual `ANSI Common LISP` spec and anything I can glean from Christian Queinnec's *Lisp In Small Pieces*.  

This is a side project, for pleasure and nostalgia.  When I think of `LISP` I think of *The Eternal Golden Braid* and that first day I sat on a bench on a hill reading about Achilles and the Tortoise.  If this project lands anywhere I would like it to be, eventually, a compiler for LISP that supports code that creates code and applies JIT compilation to the newly added code.  I would like a fast, standard LISP.  However, mostly I want to have fun with it and share some narration of what happened along the way.

## Approach
I know how to use branches and pull requests.  However, for this I am going to just push to master.  I want to include some notes and observations along the way.  Maybe they can help other programmers implementing their favorite language.  So by picking a commit to browse you can see the code and any notes I added.

## Initial Code Commit 
What do we have?

* No Unit tests, yet.  I really don't know `Go` that well and I decided to go with the rough sketch approach vs. the finely crafted Shaker Nightstand approach that TDD produces.  Tests will come and they are useful, but I have found by doing TDD you really need to understand tools and domains to be efficient.  Since this is a side project in my spare time... I'm doing the sketch approach.
* I started with ideas learned in another blog on parsers written in `Go`, but there are some serious differences between something like `LISP` and something like `Property File` syntax.   The most glaring difference being that `LISP` as I am implementing it is pretty empty headed regarding keywords, operators and other crap.  It is aware of, for the purposes of this project, `Atoms` which are `Symbols` and `Numbers` organized into `Lists`.  If the `List` isn't quoted then we will use `Prefix Evaluation` to derive the value of the `List`. I follow the spirit, but maybe not the letter of [Peter Norvig's LISPY post](http://norvig.com/lispy.html).
* Right now this commit doesn't work.
## Next
* Tests for tokenizing functions.
* Driver to run the tokenizer from start to finish.
* Add FunctionTable, OSLT, to allow interpreter to pre-define some default functions.  This means that even `+` and `*` will be treated as symbolic names and refer to functions that apply to lists of other `Atoms` and `Lists`.  The homoiconic aspect of `LISP` is really quite wonderful as it allows us to work on implementing a [Turing Complete](https://lispcast.com/the-most-important-idea-in-computer-science/) language implementation using only a few concepts.    

## Sources
* http://norvig.com/lispy.html
* https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3
* https://adampresley.github.io/2015/04/12/writing-a-lexer-and-parser-in-go-part-1.html
