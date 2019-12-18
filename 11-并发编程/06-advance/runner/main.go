package main

import (
	"log"
	"time"

	"github.com/zhangmingkai4315/concurrent-example/runner/runner"
)

func main() {

	timeout := time.Duration(5) * time.Second
	r := runner.NewRunner(timeout)
	r.Add(runner.TaskGen(), runner.TaskGen(), runner.TaskGen())
	err := r.Start()
	if err != nil {
		log.Printf("get error from runner : %v", err)
		return
	}
	log.Printf("runner stop success")

}
