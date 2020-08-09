package utils

import (
	"time"
)

// Holder wait ending of test function
type Holder interface {
	Done()
	ch() chan struct{}
}

type waitImpl struct {
	c chan struct{}
}

func (w *waitImpl) Done() {
	close(w.c)
}

func (w *waitImpl) ch() chan struct{} {
	return w.c
}

// T interface abstract testing.T
type T interface {
	Fatal(args ...interface{})
}

// RunUntil function testing with running duration limit
func RunUntil(t T, f func(Holder), duration int) {
	timeout := time.After(time.Duration(duration) * time.Second)
	waitCh := &waitImpl{c: make(chan struct{})}
	go f(waitCh)

	select {
	case <-timeout:
		t.Fatal("timeout")
	case <-waitCh.ch():
	}
}
