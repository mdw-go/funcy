# funcy

[![GoDoc](https://godoc.org/github.com/mdwhatcott/funcy?status.svg)](http://godoc.org/github.com/mdwhatcott/funcy)

What is this?

> A library providing functional-style operations like Map/Reduce/Filter (to name a few) implemented in Go w/ generics.

But Rob Pike [says we should just write for loops instead](https://github.com/robpike/filter)...

> Sorry, I guess [I just couldn't resist](https://twitter.com/codewisdom/status/1056162850220240896). I extend apologies to all the Go purists who are annoyed by this.

Are you aware that this approach lacks lazy evaluation, generates a ton of garbage, and won't scale for increasingly large inputs?

> Yup. I extend apologies to all the functional purists who are annoyed by this.

How would one install it? (Asking for a friend...)

> `go get github.com/mdwhatcott/funcy`
> 
> Enjoy!
