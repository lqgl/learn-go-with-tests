package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s', but want '%s'", got, want)
		}

	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySleeperPrinter := &CountdownOperationsSpy{}
		Countdown(spySleeperPrinter, spySleeperPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeperPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeperPrinter.Calls)
		}
	})
}
