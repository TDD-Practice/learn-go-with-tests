package main

import (
	"bytes"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep(t time.Duration) {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("Should print the countdown from 'n' to 0", func(t *testing.T) {
		buffer := bytes.Buffer{}
		spySlipper := SpySleeper{}

		Countdown(&buffer, &spySlipper, 3)

		got := buffer.String()
		want := "3\n2\n1\n0\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
