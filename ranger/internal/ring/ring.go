package ring

// Ring implements a very simple, append-only ring data structure.
// See github.com/mdwhatcott/generic-ring for a more complete implementation.
type Ring[T any] struct {
	next, prev *Ring[T]
	Value      T
}

func New[T any](n int) *Ring[T] {
	if n <= 0 {
		return nil
	}
	r := new(Ring[T])
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring[T]{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}

func (this *Ring[T]) Next() *Ring[T] {
	if this.next == nil {
		this.next = this
		this.prev = this
		return this
	}
	return this.next
}
