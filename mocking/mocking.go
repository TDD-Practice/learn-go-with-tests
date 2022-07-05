package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep(time.Duration)
}

type CountodownSleeper struct {
}

func (s *CountodownSleeper) Sleep(t time.Duration) {
	time.Sleep(t)
}

func main() {
	sleeper := CountodownSleeper{}
	Countdown(os.Stdout, &sleeper, 3)
}

func Countdown(w io.Writer, s Sleeper, n int) {
	for i := n; i >= 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep(1 * time.Second)
	}
}
