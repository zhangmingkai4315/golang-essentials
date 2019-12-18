package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zhangmingkai4315/concurrent-example/pool/pool"
)

var idForDBConnection int32

type dbConnection struct {
	ID int32
}

func (db *dbConnection) Close() error {
	log.Printf("close the db connection ID=%d", db.ID)
	return nil
}

func createDBConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idForDBConnection, 1)
	log.Printf("create a new connection id=%d", id)
	return &dbConnection{ID: id}, nil
}

func performQuery(q int, p *pool.Pool) {
	connection, err := p.Acquire()
	if err != nil {
		log.Printf("get error %s", err.Error())
		return
	}
	defer p.Release(connection)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("Query ID=%d PoolConnection ID=%d \n", q, connection.(*dbConnection).ID)
}

func main() {

	// example for runner
	// timeout := time.Duration(5) * time.Second
	// r := runner.NewRunner(timeout)
	// r.Add(runner.TaskGen(), runner.TaskGen(), runner.TaskGen())
	// err := r.Start()
	// if err != nil {
	// 	log.Printf("get error from runner : %v", err)
	// 	return
	// }
	// log.Printf("runner stop success")

	// example for pool

	var wg sync.WaitGroup
	maxGoroutine := 25

	wg.Add(maxGoroutine)
	p, err := pool.NewPool(createDBConnection, 10)
	if err != nil {
		log.Printf("create pool error :%s", err)
		return
	}
	log.Println("create pool success")
	for q := 0; q < maxGoroutine; q++ {
		go func(q int) {
			defer wg.Done()
			performQuery(q, p)
		}(q)
	}
	wg.Wait()
	p.Close()

}
