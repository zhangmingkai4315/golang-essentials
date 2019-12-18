package pool

import "errors"

// Define Custom error
var (
	ErrPoolClosed         = errors.New("Pool has been closed")
	ErrPoolSizeNotCorrect = errors.New("Pool size must be set great than zero")
)
