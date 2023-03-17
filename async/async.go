package async

import (
	"sync"
)

type Async[T any] struct {
	locker sync.Locker
	val    T
}

func Eventually[T any](f func() T) *Async[T] {
	a := &Async[T]{
		locker: &sync.Mutex{},
		val:    *new(T),
	}
	a.locker.Lock()
	go func() {
		a.val = f()
		a.locker.Unlock()
	}()
	return a
}

func Await[T any](a *Async[T]) T {
	a.locker.Lock()
	defer a.locker.Unlock()
	return a.val
}
