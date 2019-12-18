package pool

import (
	"io"
	"log"
	"sync"
)

// Pool defination
type Pool struct {
	sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// NewPool create a new pool
func NewPool(f func() (io.Closer, error), size uint) (*Pool, error) {
	if size == 0 {
		return nil, ErrPoolSizeNotCorrect
	}
	return &Pool{
		factory:   f,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire a resource from pool
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		if ok == false {
			return nil, ErrPoolClosed
		}
		log.Println("acquire the old resource from pool")
		return r, nil
	default:
		log.Println("create new resource from factory")
		return p.factory()
	}
}

// Release the resource to  pool
func (p *Pool) Release(r io.Closer) {
	p.Lock()
	defer p.Unlock()

	if p.closed == true {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Printf("[Pool] : return the resoure")
	default:
		// the resourecs is full now
		r.Close()
	}
}

// Close the pool and close all resource
func (p *Pool) Close() {
	p.Lock()
	defer p.Unlock()

	if p.closed {
		return
	}

	p.closed = true
	close(p.resources)
	for r := range p.resources {
		r.Close()
	}
}
