# funcy

[![GoDoc](https://godoc.org/github.com/smartystreets/gunit?status.svg)](http://godoc.org/github.com/smartystreets/gunit)

What is this?

> A library providing functional-style operations like Map/Reduce/Filter (to name a few) implemented in Go w/ generics.

But Rob Pike already did this (and way before generics were introduced) and then told us we should just write for loops instead.

> Sorry, I guess [I just couldn't resist](https://twitter.com/codewisdom/status/1056162850220240896).

Are you aware that this approach lacks lazy evaluation, generates a ton of garbage, and won't scale for increasingly large inputs?

> Yup.

How would one install it? (Asking for a friend...)

> `go get github.com/mdwhatcott/funcy`
> 
> Enjoy!
