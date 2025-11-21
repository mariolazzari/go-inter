package pool

import "sync"

type Pool[T any] struct {
	pool sync.Pool
}

func New[T any](newFn func() T) *Pool[T] {
	var p Pool[T]
	if newFn != nil {
		p.pool.New = func() any {
			return newFn()
		}
	}

	return &p
}

func (p *Pool[T]) Put(v T) {
	p.pool.Put(v)
}

func (p *Pool[T]) Get() (T, bool) {
	v := p.pool.Get()
	if v == nil {
		var zero T
		return zero, false
	}

	return v.(T), true
}
