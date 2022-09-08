package main

import (
	"testing"
	"time"
)

func TestCompileChannels(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	{
		out := CompileChannels(
			sig(2*time.Hour),
			sig(5*time.Minute),
			sig(1*time.Millisecond),
			sig(1*time.Hour),
			sig(1*time.Minute),
		)

		select {
		case <-out:
		case <-time.After(time.Second):
			t.Fail()
		}
	}

	{
		out := CompileChannels(
			sig(500*time.Millisecond),
			sig(2*time.Hour),
			sig(5*time.Minute),
			sig(1*time.Hour),
			sig(1*time.Minute),
		)

		select {
		case <-out:
		case <-time.After(time.Second):
			t.Fail()
		}
	}

	{
		out := CompileChannels(
			sig(2*time.Hour),
			sig(5*time.Minute),
			sig(1*time.Hour),
			sig(1*time.Minute),
		)

		select {
		case <-out:
			t.Fail()
		case <-time.After(500 * time.Millisecond):
		}
	}
}