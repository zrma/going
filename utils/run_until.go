package utils

import (
	"time"
)

// Done chan wait ending of test function
type Done chan struct{}

// T interface abstract testing.T
type T interface {
	Fatal(args ...interface{})
}

// RunUntil function testing with running duration limit
func RunUntil(t T, f func(Done), duration int) {
	timeout := time.After(time.Duration(duration) * time.Second)
	waitCh := make(Done)
	go f(waitCh)

	select {
	case <-timeout:
		t.Fatal("timeout")
	case <-waitCh:
	}
}
