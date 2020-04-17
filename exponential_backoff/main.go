package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type ExponentialBackoff struct {
	Operation             func() (string, error)
	BackoffPeriodDuration time.Duration
	MaximumRetries        int
	NoOfRetriesLeft       int
	Jitter                int
}

// helper function to check if we have hit the max retries
func (e *ExponentialBackoff) ShouldRetry() bool {
	return e.NoOfRetriesLeft > 0
}

// method that sleeps for the given duration of time in seconds
func (e *ExponentialBackoff) WaitUntilNextTry() {
	fmt.Println("wait for:", e.BackoffPeriodDuration.Seconds(), "seconds before retrying")
	time.Sleep(e.BackoffPeriodDuration)
}

// method that handles exception and retry
func (e *ExponentialBackoff) OperationFailed() {
	e.NoOfRetriesLeft -= 1

	// check if we should retry again
	if !e.ShouldRetry() {
		fmt.Println(
			"Failed: Max no. of attempts:",
			e.MaximumRetries,
			"| Total duration:",
			e.BackoffPeriodDuration.Seconds(),
			"seconds",
		)
		os.Exit(1) // exit the program since max. no of attempts have been made
	}

	e.WaitUntilNextTry() // sleep for the random duration

	// randomise the delay for next sleep operation
	seconds := fmt.Sprintf("%vs",
		(e.BackoffPeriodDuration.Seconds()*2.0)+float64(e.Jitter),
	)
	duration, _ := time.ParseDuration(seconds)
	e.BackoffPeriodDuration = duration

	e.Perform() // retry
}

// method to invoke the operation which is bound to fail
func (e *ExponentialBackoff) Perform() {
	fmt.Println("Trying to perform the operation ...")
	str, err := e.Operation()

	if str != "" {
		fmt.Println(str)
	}

	if err != nil {
		fmt.Println(err)
		e.OperationFailed()
	}
}

// ExponentialBackoff constructor
func NewExponentialBackoff(o func() (string, error), d string, m, j int) *ExponentialBackoff {
	duration, _ := time.ParseDuration(d)

	return &ExponentialBackoff{
		Operation:             o,
		BackoffPeriodDuration: duration,
		MaximumRetries:        m,
		NoOfRetriesLeft:       m,
		Jitter:                j,
	}
}

// Example struct to simulate a real world operation
type Hello struct{}

func (h *Hello) FailToSayHi() (string, error) {
	return "", errors.New("Failed to say hello to the world!")
}

func (h *Hello) SayHi() (string, error) {
	return "Hello, world!", nil
}

func main() {
	h := Hello{} // struct with methods
	fmt.Println("Execute with BackOff Strategy:")

	backOff := NewExponentialBackoff(
		h.FailToSayHi, // method which could fail
		"1s",          // duration of backoff periods
		5,             // criteria for giving up is max number of attempts
		rand.Intn(10), // jitter or random delay
	)

	backOff.Perform() // try executing the method which could fail with backoff strategy
}
