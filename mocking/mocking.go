package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "go!"
const countdownStart = 3

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type Sleeper interface {
	Sleep()
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprint(out, i)
		sleeper.Sleep()
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

// At this point, I still need a refresher on pointers, using & and *, https://www.golang-book.com/books/intro/8

// also, interfaces - https://www.golang-book.com/books/intro/9
// Like a struct an interface is created using the type keyword, followed by a name and the keyword interface. But instead of defining fields, we define a “method set”. A method set is a list of methods that a type must have in order to “implement” the interface.
