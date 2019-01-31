### 6. Struct结构体

struct结构体可以组合多个不同的类型在一起，并将其绑定在一个自定义的结构体类型中, 便于统一的管理。同时可以绑定专结构体的函数，用于后期调用。下面是定义结构体和创建结构体的方式：

```go
    type Message struct {
        sender   string
        body     string
        receiver string
        private  bool
    }
	m1 := Message{
		sender:   "mike",
		receiver: "alice",
		body:     "hi alice, how are you",

	}
	m2 := Message{"alice", "Fine, Thanks", "mike"}

```

上面的代码中，使用了两种方式来创建结构体，第一种明确的属性名称，这样我们就不需要严格的按照结构体的顺序来补充内容，第二种则必须按照顺序进行填写。

#### 6.1 嵌套结构体

对于上面的Message结构体，我们可以直接增加新的属性，或者使用嵌套的方式组合不同的结构体。组合的优势在于新的结构体的属性和方法都会直接的被加入进来。

```go
type MessageDetail struct {
	IP   string
	Date time.Time
}

// Message struct
type Message struct {
	sender   string
	body     string
	receiver string
	private  bool
	MessageDetail
}
```

同样对于新建的结构体对象，我们需要传递更多的值到结构体中，其中MessageDetail并没有属性名称(当然可以指定一个名称)，我们传递的方式如下：

```go
	m1 := Message{
		sender:   "mike",
		receiver: "alice",
		body:     "hi alice, how are you",
		MessageDetail: MessageDetail{
			IP:   "127.0.0.1",
			Date: time.Now(),
		},
	}
	m2 := Message{"alice", "Fine, Thanks", "mike", false, MessageDetail{"127.0.0.1", time.Now()}}
```

对于其中的time.Date本身也是一个结构体对象。嵌入的结构体的内部属性或者方法可以直接在对象上被访问到：

```go
fmt.Printf("m1： ip = %s , date = %s\n", m1.IP, m1.Date)
// m1： ip = 127.0.0.1 , date = 2019-01-07 03:46:07.243995163 +0800 CST m=+0.000202821
```

#### 6.2 匿名结构体

有时候结构体在包中定义后仅仅使用一次，这时候我们可以使用匿名结构体的方式进行处理。使用匿名结构体，使得程序更加的简洁，无需在整个的包范围内创建一些不必要的结构。

```go
	m1 := struct {
		Name, Address string
	}{
		Name:    "Mike",
		Address: "Beijing",
	}
	fmt.Printf("%+v", m1)
```

#### 附录

##### 1. Go语言是否是面向对象的语言

传统的面向对象编程语言是通过类和继承的方式实现，在go中并没有class类的概念，而是通过type的方式实现，初始化类变成了创建类型的值。Go通过使用type和绑定函数的方式来实现面向对象类似的编程方式。同时使用interface提供一种类的抽象，类似Java中接口，并且利用嵌套的方式实现类似于子类的编程方式。

对于方法绑定go的实现也更加灵活简单，可以绑定方法在任何的类型上，甚至是一个简单的数据类型（整型，字符串类型等）。

参考链接:

https://www.ardanlabs.com/blog/2013/07/object-oriented-programming-in-go.html

