### 11. 并发编程

#### 11.1 Concurrency is not parallelism

并发不同于并行，并行需要依赖于特殊的硬件结构去实现，如果仅仅只有一个处理器，无法实现并行运行，但是仍旧可以实现多个任务的并发执行，同时一些并行的程序可以在多核的处理器上获得更好的运行效率。并发执行的概念可以追溯到1978年的时候提出的CSP模型，一些传统的编程语言比如Erlang, Limbo等基于该模型实现并发的事务处理。

> Erlang的高并发基于进程级别的，每个事务独立占有一个轻量级的进程，每个进程有自己的垃圾回收器，而go通过goroutine则是使用微线程的方式实现高并发，全局共享一个垃圾回收，因此回收资源的时候，可能导致所有任务的中断和延迟，但是Erlang本身的动态语言特性也导致程序运行效率的缓慢。

**goroutines**: 是一个独立执行的函数，通过使用go方法来生成，每一个goroutines具有自己的调用栈，实现的代价较小，一个程序可以同时生成几千和上万个goroutines一起运行。**goroutine本身不是线程**，一个程序可以只有一个线程，但是具有上千个goroutines,大量的goroutines可以被多个线程动态的复用去实现并发执行。一些参考的资料如下：


- [Concurrency-is-not-parallelism](https://blog.golang.org/concurrency-is-not-parallelism)
- [Concurrency slide](https://talks.golang.org/2012/concurrency.slide)
- [Effective Go](https://golang.org/doc/effective_go.html#concurrency)


#### 11.2 goroutine

如果接触过其他的编程语言，可能听说过利用线程，进程以及协程等来完成程序的调度执行，而goroutine是go语言中实现任务调度的一种方式，goroutine简单理解就是一个可以并发执行的函数体。Go语言通过go关键词来执行一个函数，这个函数将被独立的运行在一个并发的goroutine中，与其他进程（goroutine）并发执行。

下面的是一个goroutine执行的实例，但是如果我们直接执行的话，会发现程序立即退出，指定的goroutine执行函数还来不及执行，因为go语言中的goroutine并不是阻塞的，当产生一个新的goroutine不会影响原有的goroutine执行，main函数一旦结束，程序自动的退出，所有已经产生的goroutine也就会被直接销毁退出执行。

```golang
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}

}
func main() {
    go boring("hello")
}
```

Goroutine和线程的关系是，在操作系统层面的线程与goroutine是多对多复用关系的，多个goroutine可以被多个线程所调度执行，如果当前线程绑定的一个goroutine被阻塞，比如处于I/O等待状态，其他的goroutine会被当前的线程调度起来执行。go语言层面隐藏了很多线程创建和管理的复杂逻辑，使得利用goroutine编写并发程序更加的简单方便。

#### 11.2 WaitGroup

为了实现上述例子中，子goroutine的正常执行和退出，我们可以使用等待的方式来防止程序退出，比如在主程序最后加入等待5秒或者10秒退出，但是这种方式灵活性太差，一旦子goroutine执行时间超过等待时间，仍旧会被清理掉，无法顺利执行。go语言中可以通过很多种方式可以并发的顺利执行，这里先介绍一种标准库中经常被用到的方式WaitGroup，等待组的概念。具体的接口和函数位于sync库中。

WaitGroup中包含的函数比较简介，只有三个函数可以使用：
- Add: 每次启用一个新的goroutine增加一个，或者同时增加多个计数。
- Done：每个goroutine退出前执行
- Wait：等待直到所有的WaitGroup都已经执行结束

具体的使用方式如下面所示： 

```golang
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("No %d goroutine begin\n", i)
			fmt.Printf("No %d goroutine is done\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```

通过利用WaitGroup可以保证goroutine的顺利执行完毕，相对于利用time休眠程序可以减少不必要的等待时间。另外wg.Done()函数建议放入defer函数中，这样goroutine程序如果发生异常不会导致程序一直处于等待状态无法顺利退出。


#### 11.2 并发环境下的共享变量

在并发编程的情况，访问共享变量如果处理不当将会造成数据的读写出现问题，有可能出现多个线程或者goroutine同时访问同一个对象的情况，着就会导致程序输出结果和预期不一致，如下面的一个代码所示。如果我们试着执行该函数，会发现程序的输出结果与期望的结果相差比较大，在我的机器上运行的输出如下面所示：
```go
var wg sync.WaitGroup

func main() {
	var shared = 0
	wg.Add(2)
	go func() {

		for i := 0; i < 10000; i++ {
			shared++
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			shared++
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("the value of shared = %d , expected=20000", shared)
}

# the value of shared = 11649 , expected=20000
```


整个的程序执行流程与下面的图示比较相似，当两个线程或者goroutine尝试同时访问一个对象并对对象进行+1操作的时候，会导致读写不一致，新写入的值可能无法立刻被下一个线程感知，本身读写都需要一定的时间，这就造成同时读写的数据竞争问题。
![](http://syshex.files.wordpress.com/2011/10/thread-concurrency.png)

Go语言在1.5版本前修改了默认的最大CPU核心使用数目，原来默认的数目是1，当前版本的数目是CPU核心的数量，可以通过查看GOMAXPROCS查看该信息。


> Do not communicate by sharing memory; instead, share memory by communicating.


#### 11.4 Mutex锁机制

使用go语言可以通过sync库中的mutex来实现数据的读写锁机制，防止发生数据竞争的问题。如下的例子是对于上面的例子的重写，通过利用sync.Mutex来锁定对于数据的读写，读写锁定同时可以设定锁的范围（利用Lock和Unlock范围设定范围区域）。由于锁定后只有解锁之后才能被其他的goroutine执行，因此又称为全局锁。

```go

type sharedWithMutex struct {
	shared int
	mu     sync.Mutex
}

func main() {
	wg.Add(2)
	s := sharedWithMutex{
		shared: 0,
	}
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 0; i < 10000; i++ {
			s.mu.Lock()
			s.shared++
			s.mu.Unlock()
		}
	}()
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 0; i < 10000; i++ {
			s.mu.Lock()
			s.shared++
			s.mu.Unlock()
		}
	}()
	wg.Wait()
	fmt.Printf("the value of shared = %d , expected=20000", s.shared)
	// the value of shared = 20000 , expected=20000
}

```

sync库中还有一个可以提供程序锁定的接口： RWMutex相对于Mutex提供了读锁和写锁的分离，使用读锁可以允许多个goroutine同时读取数据而不被锁定，一旦有goroutine尝试修改则锁定不能被执行，直到读取完成，通过分离实现更加灵活的锁定机制，提高程序的执行效率。

#### 11.5 Atomic原子操作

为了消除数据竞争的问题，go语言还提供了另一种简单的解决方案，使用atomic库来直接原子的修改数据，修改操作是一个原子操作（无法被任何线程，goroutine等并发机制拆分）以下是上述实例的atomic版本
```go
func main() {
	var shared int64
	wg.Add(2)
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 0; i < 10000; i++ {
			atomic.AddInt64(&shared, 1)
		}
	}()
	go func() {
		defer func() {
			wg.Done()
		}()
		for i := 0; i < 10000; i++ {
			atomic.AddInt64(&shared, 1)
		}
	}()
	wg.Wait()
	fmt.Printf("the value of shared = %d , expected=20000", shared)
	// the value of shared = 20000 , expected=20000

}
```

