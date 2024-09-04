# github.com/mdwhatcott/funcy `[DEPRECATED]`

NOTE: the `funcy` package described first below has been deprecated, in effect having been replaced by the `funcy/ranger` package (describe farther below).

---

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

---

# github.com/mdwhatcott/funcy/ranger

[![GoDoc](https://godoc.org/github.com/mdwhatcott/funcy/ranger?status.svg)](http://godoc.org/github.com/mdwhatcott/funcy/ranger)

Despite what I thought was a very cleverly written disclaimer (above), Go went and released version 1.23 with iterators, which means we can have our cake and eat it too, a phrase which here means that lazy evaluation is now possible and so we no longer need upset functional purists!

I'm not sure how Rob feels about all this, but I sincerely hope you do enjoy the `funcy/ranger` package, with [many accompanying examples](https://github.com/mdwhatcott/funcy/tree/main/ranger/examples).