package utils

import (
	"time"
)

// Wait chan wait ending of test function
type Wait chan struct{}

// T interface abstract testing.T
type T interface {
	Fatal(args ...interface{})
}

// RunUntil function testing with running duration limit
func RunUntil(t T, f func(Wait), duration int) {
	timeout := time.After(time.Duration(duration) * time.Second)
	waitCh := make(Wait)
	go f(waitCh)

	select {
	case <-timeout:
		t.Fatal("timeout")
	case <-waitCh:
	}
}
