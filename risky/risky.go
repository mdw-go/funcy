// Deprecated: use github.com/mdwhatcott/funcy/ranger/* instead.
// Package risky houses convenience functions that introduce runtime risks (i.e. cavalier use of reflection)
// and therefore are separated from the root funcy package.
package risky

import "reflect"

// Field provides a func() for convenient use with funcy.Map, funcy.Filter, funcy.Sort*, etc.
func Field[S, V any](name string) func(s S) V {
	return func(s S) V { return reflect.ValueOf(s).FieldByName(name).Interface().(V) }
}
