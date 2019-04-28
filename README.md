# norvig_lisp
A Go lang implementation of Norvig's LISP subset originally done in Python.
## Why
There's lots of `LISP` parsers and this one is yet another one.  For me this is a chance to
learn `Go` in context and build on what I learned in [coding along with this blog](https://adampresley.github.io/2015/04/12/writing-a-lexer-and-parser-in-go-part-1.html).
It is also a warmup before I attempt the actual `ANSI Common LISP` spec and anything I can glean from Christian Queinnec's *Lisp In Small Pieces*.

This is a side project, for pleasure and nostalgia.  When I think of `LISP` I think of *The Eternal Golden Braid* and that first day I sat on a bench on a hill reading about Achilles and the Tortoise.  If this project lands anywhere I would like it to be, eventually, a compiler for LISP that supports code that creates code and applies JIT compilation to the newly added code.  I would like a fast, standard LISP.  However, mostly I want to have fun with it and share some narration of what happened along the way.

## Approach
I know how to use branches and pull requests.  However, for this I am going to just push to master.  I want to include some notes and observations along the way.  Maybe they can help other programmers implementing their favorite language.  So by picking a commit to browse you can see the code and any notes I added.

## The Bloody Trail of Commits
### Initial
* No Unit tests, yet.  I really don't know `Go` that well and I decided to go with the rough sketch approach vs. the finely crafted Shaker Nightstand approach that TDD produces.  Tests will come and they are useful, but I have found by doing TDD you really need to understand tools and domains to be efficient.  Since this is a side project in my spare time... I'm doing the sketch approach.
* I started with ideas learned in another blog on parsers written in `Go`, but there are some serious differences between something like `LISP` and something like `Property File` syntax.   The most glaring difference being that `LISP` as I am implementing it is pretty empty headed regarding keywords, operators and other crap.  It is aware of, for the purposes of this project, `Atoms` which are `Symbols` and `Numbers` organized into `Lists`.  If the `List` isn't quoted then we will use `Prefix Evaluation` to derive the value of the `List`. I follow the spirit, but maybe not the letter of [Peter Norvig's LISPY post](http://norvig.com/lispy.html).
* Right now this commit doesn't work.
### Changing The Tokenizer: A Hard Lesson in Channel Abuse
* Added a module level Repl.go to make this build and be a module until I get a better handle on Go mod stuff.
* Pretty much dumped all the old parser stuff.  I was digging into what I copied into my code and found out the original parser was an incomplete
implementation of one that used Goroutines.  I wasn't sure why there was a channel of length > 1, by experimenting and reading the documentation I learned that code would block without a channel that was bigger than needed.  However, the original code did not appear to use Goroutines as I understood them.  Instead I went with a simpler model based on Thorsten Ball's book.
* Differences from Thorsten Ball's example: go.mod, using rune vs. character, target syntax is mini-LISP.  Some naming and other conventions.
* I now have a simple Unit Test.
## Small Addition: Tokenize and Gathering Info for Parsing and Default Environment
* Lining up with what I think I need: Tokenize to give back a slice of all tokens.
* Look up what I think I need to create the Default Environment (built in functions) and the
semantics of these.  Some of them, like subtraction, have a unary and n-ary behavior.
    * https://stackoverflow.com/questions/21326109/why-are-lists-used-infrequently-in-go
    * https://gobyexample.com/closures
    * https://gobyexample.com/collection-functions
    * https://en.wikibooks.org/wiki/Common_Lisp/First_steps/Beginner_tutorial#Subtraction

## Started Trying to Figure Out the Syntax PArsing
* Created a set of structs that might be how I do this.  Learning that Go
type system rules and limitations.  Since I am basing this on the original
article that used Python, there is a bit of a difference.
* Added lexical support for double quoted strings.  Want them, now have
them.  I want to use a single quote for quoting lists and symbols I think.

## Maybe Tomorrow... maybe someday
* Maybe rename files to snake case?  Need to figure out the conventions here.

## Sources
* http://norvig.com/lispy.html
* https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3
* https://adampresley.github.io/2015/04/12/writing-a-lexer-and-parser-in-go-part-1.html
* Writing An Interpreter In Go - Thorsten Ball




salespersonUserId: "ik-bBujsQmy39JZniaKrpw"
