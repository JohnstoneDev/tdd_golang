package main

import (
	"io"
	"os"
	"fmt"
	"time"
)

// constants
const finalWord = "Go!"
const startCount = 3
const write = "write"
const sleep = "sleep"

// types
type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// Configurable sleeper that can sleep for n-seconds
type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

// type that records the time slept
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// function that prints a countdown, sleeps
// for some time between printing
func Countdown(out io.Writer, sleeper Sleeper) {
	// Countdown then print, when loop ends print final word
	for i := startCount; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}