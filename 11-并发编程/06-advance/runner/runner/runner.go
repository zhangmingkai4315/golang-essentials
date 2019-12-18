package runner

import (
	"os"
	"os/signal"
	"time"
)

// Runner defination will stop the tasks when receive
// interrupt from os or timeout
type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	tasks     []func(int)
}

// NewRunner create a new Runner with timeout
func NewRunner(timeout time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(timeout),
	}
}

// Add some tasks to the runner
func (runner *Runner) Add(tasks ...func(int)) {
	runner.tasks = append(runner.tasks, tasks...)
}

// Start a runner
func (runner *Runner) Start() error {
	signal.Notify(runner.interrupt, os.Interrupt)
	go func() {
		runner.complete <- runner.run()
	}()
	select {
	case err := <-runner.complete:
		return err
	case <-runner.timeout:
		return ErrTimeout
	}
}

func (runner *Runner) run() error {
	for id, task := range runner.tasks {
		if runner.gotInterupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (runner *Runner) gotInterupt() bool {
	select {
	case <-runner.interrupt:
		signal.Stop(runner.interrupt)
		return true

	default:
		return false
	}
}
