### 7. Struct结构体

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

#### 7.1 嵌套结构体

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

#### 7.2 匿名结构体

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

#### 7.3 使用new关键词
使用new关键词可以用来创建一个结构体，但是返回的是对应的结构体指针，所有的内部类型将按照类型进行零值处理，比如对于整形则返回0，对于string类型则返回空字符串等等。

```golang

func main() {
	type person struct {
		name   string
		age    int
		isMale bool
		phones []string
	}

	p := new(person)
	fmt.Printf("the type of p = %T\n", p)
	// the type of p = *main.person
	fmt.Printf("the value of p = %+v", p)
	// the value of p = &{name: age:0 isMale:false phones:[]}
	p.phones = append(p.phones, "12345678")
	fmt.Printf("the value of p = %+v", p)
	// the value of p = &{name: age:0 isMale:false phones:[12345678]}
}
```
使用这种方式可以用来进行结构体的初始化，然后对于结构体再进行自定义的修改，这在编写代码的时候是一种很常见的编写方式，比如下面的例子, 使用new来创建一个File结构体指针，再将一些内容传递到结构体中实现对于结构体的修改。

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```
除了使用上面的new方式创建，我们可以直接使用结构体本身的初始化方式创建一个对象并返回：
```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```

在6.2中我们介绍了使用make来创建slice类型，其实除了slice和数组类型以外，make还可以初始化map类型和channel类型，返回一个初始化的类型值，而不是其类型指针。使用make和new的区别在与是否进行初始化还是执行取零操作,如下面的操作所示:

```go
var p *[]int = new([]int)      
var v  []int = make([]int, 100) 

```
第一个初始化过程中，使用new来创建一个int的切片类型，这时候仅仅返回一个指针，该指针此时满足``` *p == nil ``` 后续如果使用还需要进行初始化操作，因此很少使用这种方式，第二种方式则比较常见，使用make初始化操作并创建一个占用100个int空间的数据结构，返回对应的引用，而不是其指针。

> 使用make仅仅可以用在切片，数组和map类型以及channel类型上，返回的是类型而不是类型指针，如果需要指针的话使用取地址操作符即可。


#### 附录

##### 1. Go语言是否是面向对象的语言

传统的面向对象编程语言是通过类和继承的方式实现，在go中并没有class类的概念，而是通过type的方式实现，初始化类变成了创建类型的值。Go通过使用type和绑定函数的方式来实现面向对象类似的编程方式。同时使用interface提供一种类的抽象，类似Java中接口，并且利用嵌套的方式实现类似于子类的编程方式。

对于方法绑定go的实现也更加灵活简单，可以绑定方法在任何的类型上，甚至是一个简单的数据类型（整型，字符串类型等）。

参考链接:

https://www.ardanlabs.com/blog/2013/07/object-oriented-programming-in-go.html

