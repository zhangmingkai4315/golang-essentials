package work

import "sync"

// Worker interface for job excution
type Worker interface {
	Task()
}

// Pool for workers
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// NewWorkPool create a work pool
func NewWorkPool(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

// Shutdown close the pool
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

// Submit a new work to do
func (p *Pool) Submit(w Worker) {
	p.work <- w
}
