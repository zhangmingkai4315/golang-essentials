package runner

import (
	"errors"
	"log"
	"time"
)

// Err define some errors variables
var (
	ErrTimeout   = errors.New("timeout error")
	ErrInterrupt = errors.New("interrupt error")
	ErrUnknown   = errors.New("unknown")
)

// TaskGen create a simple task
func TaskGen() func(int) {
	return func(id int) {
		log.Printf("start %d task", id)
		time.Sleep(time.Duration(id) * time.Second)
		log.Printf("stop %d task", id)
	}
}
