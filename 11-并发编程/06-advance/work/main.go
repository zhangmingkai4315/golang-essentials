package main

import (
	"log"
	"sync"
	"time"

	"github.com/zhangmingkai4315/concurrent-example/work/work"
)

var names = []string{
	"www.baidu.com",
	"www.sina.com.cn",
	"www.google.com",
	"www.mircosoft.com",
	"www.abc.com",
}

type namePrinter struct {
	name string
}

func (np *namePrinter) Task() {
	log.Println(np.name)
	time.Sleep(time.Millisecond * 10)
}

func main() {
	p := work.NewWorkPool(10)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))
	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			go func() {
				p.Submit(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	p.Shutdown()
}
