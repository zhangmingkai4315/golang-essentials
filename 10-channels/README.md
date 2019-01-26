### 10. Channel

#### 10.1 channel基础

创建一个channel的最简单的方式是通过关键词make来实现，对于用来传递具体的类型的channel,在使用make创建channel的时候需要同时指定，另外channel包含两种类型：

- 阻塞channel： 必须有其他的goroutine去消费传递给该channel的数据，否则任何写入数据的操作都会阻塞当前goroutine执行。
- 缓冲channel： 允许最多N个数据的写入，而不会阻塞当前goroutine的执行。

```
ci := make(chan int)            // 整型非缓冲channel
cj := make(chan int, 0)         // 同上
cs := make(chan *os.File, 100)  // 最大100个数据（文件指针类型）存储的缓冲channel
```

下面是使用channel的两种方式，非缓冲channel和缓冲channel的工作原理：

```go
func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
	fmt.Println("exit success")
}


```

上面的程序中启动独立的goroutine来消费主goroutine中写入阻塞channel的数据，使得程序能够正常的执行下去，否则如果新创建的独立goroutine的话，程序会处于deadlock状态。

```go


func main() {
	c := make(chan int, 1)
	c <- 1
	fmt.Println("not block this message")
	c <- 2
	fmt.Println("block again, never show this message")
}

```

第二个实例是利用了缓冲channel的存储特性，程序不会在写入1的时候阻塞，但是一旦超过了最大的存储容量，则处于阻塞状态，因此程序会阻塞在写入2的位置，无法写入，处于deadlock状态。



#### 10.2 单向channel

上面创建的channel都是双向channel，也就是既可以写入也可以读取操作，而单向channel则是仅仅只能读取或者只能写入的channel， 主要用于函数参数或者返回值中使用。比如下面的例子：

```go
func producer() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			time.Sleep(time.Second * 1)
		}
		close(c)
	}()
	return c
}

func consumer(c <-chan int) {
	for x := range c {
		fmt.Println(x)
	}
}

func main() {
	p := producer()
	consumer(p)
}


```



上述实例中我们通过producer产生一个单向的channel，能够读取数据的channel， 这样可以防止在此之外的程序不会写入数据到channel中。同时我们使用函数consumer来接收一个只读的**单向channel类型**用于消费生产的数据。利用**range来迭代channel中的数据**，需要注意的是，如果channel最后没有close关闭的话，程序会处于琐死状态，consumer会认为这个channel将会传递数据过来因此处于等待状态。因此务必在确定不写入数据后关闭该channel.

#### 10.3 select语句

select语句用于在多个channel中选择数据，一旦任何channel数据可读，则执行对应的语句，**如果在选择的时刻多个channel可读，则随机选择其中一个执行**，我们对于上面的例子进行改造，使其能够实现超时关闭的特性。这里我们引入了一个新的函数time.Ticker()，该函数可以用来计时，返回一个channel类型，一旦时间到期则传递数据到该channel中，我们可以利用该机制来进行超时管理.

```go
func producer() <-chan int {
	c := make(chan int)
	go func() {
		for {
			c <- rand.Int() % 1000
			time.Sleep(time.Millisecond * 100)
		}
	}()
	return c
}

func main() {
	p := producer()
	t := time.Tick(time.Second * 5)
	for {
		select {
		case v := <-p:
			fmt.Printf("%d\t", v)
		case <-t:
			fmt.Println("time out")
			return
		}
	}
}
```

注意select本身是单次执行，为了不断的从多个channel中获取数据，我们一般将其放入一个for循环中，保证不断轮询不同的channel.  当五秒超时时间到达的时候，则使用return退出程序执行.

另外我们还可以借助于额外的channel来设定退出，比如下面的producer我们传递了一个独立的channel来接收退出信号。一旦生产数据时满足某一些条件则传递退出信号（或者直接关闭channel）,当select的时候可以根据channel的状态来判断是否退出程序，这里我们关闭了退出channel，则select的时候```case q, ok := <-quit:```会收到信号提示channel已经关闭。

```go
func producer(quit chan<- bool) <-chan int {
	c := make(chan int)
	go func() {
		for {
			v := rand.Int() % 10
			if v == 9 {
				close(quit)
			} else {
				c <- v
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	quit := make(chan bool)
	p := producer(quit)
	for {
		select {
		case v := <-p:
			fmt.Printf("%d\t", v)
		case q, ok := <-quit:
			fmt.Printf("receive data=%v and status=%v", q, ok)
            //receive data=false and status=false
			return
		}
	}
}
```

#### 10.4 Fan in和Fan out模型

Fan in 模型通过接收多个channel的值，将其合并入一个channel，主要用来汇聚数据使用，比如我们有多个producer， 我们可以使用fan in 模型将其汇聚成一个producer. 我们修改原来的代码如下：

```go
func producer() <-chan int {
	c := make(chan int)
	go func() {
		for {
			c <- rand.Int() % 1000
			time.Sleep(time.Millisecond * 1000)
		}

	}()
	return c
}

func merge(producers ...<-chan int) <-chan int {
	all := make(chan int)
	for _, p := range producers {
		go func(p <-chan int) {
			for {
				all <- <-p
			}
		}(p)
	}
	return all
}

func main() {
	pl := []<-chan int{}
	for i := 0; i < 10; i++ {
		p := producer()
		pl = append(pl, p)
	}
	all := merge(pl...)
	t := time.Tick(time.Second * 5)
	for {
		select {
		case v := <-all:
			fmt.Printf("%d\t", v)
		case <-t:
			fmt.Println("timeout")
			return
		}
	}
}

```

上述代码中我们使用了一个merge函数来接收任意数量的channel， 根据数量来产生对应的goroutine来单独的接收数据并写入唯一的一个汇聚channel中，这里我们直接使用for循环读取channel数据直接写入另一个channel（无需select）. merge函数返回汇聚的channel提供给其他的代码消费该数据集合。

Fan out模型，最常见的使用场景是，接收大量的任务，这些任务需要启动多个work一起完成，进行负载分发的情况，如下面所示我们有一个producer，用于生产任务，而创建10个worker来执行任务，每个worker接收到任务根据传递数据进行耗时的操作（休眠），代码如下所示:

```go
var wg sync.WaitGroup

func producer(counter *int) <-chan int {
	c := make(chan int)
	go func() {
		t := time.Tick(time.Second * 10)
		for {
			c <- rand.Int() % 5
			*counter++
			select {
			case <-t:
				close(c)
				return
			default:
				time.Sleep(time.Millisecond)
			}
		}
	}()
	return c
}

func worker(id int, job <-chan int, counter *int32) {
	defer wg.Done()
	for x := range job {
		fmt.Printf("ID=%d, receive job %d\n", id, x)
		time.Sleep(time.Duration(x) * time.Second)
		fmt.Printf("ID=%d, Done job %d\n", id, x)
		atomic.AddInt32(counter, 1)
	}
}
func main() {
	var producerJobCounter int
	producer := producer(&producerJobCounter)
	var doneCounter int32
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, producer, &doneCounter)
	}
	wg.Wait()
	fmt.Printf("Total create job %d, done %d jobs\n", producerJobCounter, doneCounter)
	// Total create job 57, done 57 jobs
}

```

上面的代码中我们使用worker来接收数据并执行任务，利用两个计数器producerJobCounter和doneCounter来追踪任务的执行情况，其中doneCounter由于多个channel可以同时访问并修改，为了防止并发执行环境下的数据竞争问题，我们通过atomic库来执行数据的增加（原子操作），这里我们启用10个worker来同时消费数据，尝试修改最大的worker数量，可以看到任务的执行数量也会发生变化，实际运行情况下，worker可以设置固定的或者根据任务情况动态修改。

 #### 10.5 context上下文管理

context包主要用与上下文的管理，超时，取消任务等操作，go语言在1.7版本中将原来放在扩展库中的context包转移到了标准库中。在goweb编程中，context被用来管理用户的查询请求链，每一个请求都被放在一个独立的goroutinezhong , 同时程序还会启动多个额外的goroutine来查询数据库，执行rpc请求，这些goroutine集合可以通过context来统一管理，比如取消用户查询的同时关闭所有涉及这个请求的所有goroutine.,从而尽快的回收资源降低消耗。

下面是使用context来执行任务取消的一个例子:

```go
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Canceled by someone, i am leaving")
				return
			default:
				fmt.Println("working....")
				time.Sleep(time.Second * 1)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()

}

```

这个例子中我们使用context来创建一个具有取消功能的上下文管理器，任何时候调用cancel都可以将该上下文关闭，这样我们在实际执行的任务中可以通过监听上下文执行Done()的状态来判断是否被取消掉了。从而执行一些任务的结束和退出。

其中```context.Background()``` 是上下文管理器的根，无法被取消，所有扩展功能的均从根部创建，除了具有取消功能的还有一些其他的扩展上下文管理器：

- [WithCancel](https://golang.org/pkg/context/#example_WithCancel)

- [WithDeadline](https://golang.org/pkg/context/#example_WithDeadline)

- [WithTimeout](https://golang.org/pkg/context/#example_WithTimeout)

- [WithValue](https://golang.org/pkg/context/#example_WithValue)

上述的管理器分别用来执行取消，超时，传递数据等操作，具体的例子可以参考对应的官方手册，此处不再一一列出。

#### 10.6 附录

- [Go语言经典语句](https://go-proverbs.github.io/)
- [Golang Pipelines](https://blog.golang.org/pipelines)
- [Go Context ](https://blog.golang.org/context)
- [Peter's blog about context](https://peter.bourgon.org/blog/2016/07/11/context.html)