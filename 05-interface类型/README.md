### 5 interface接口类型

接口是go语言中用来定义一些具有相同行为（函数）的一种类型，是一种抽象的类型表示方式。对于只有一个函数的接口类型，则接口的名称一般是函数名加-er的方式，比如Reader, Writer，Formatter等等，对于具有多个函数的接口可自定义接口类型的名称。

对于实际的接口对象值，在内存中是通过两个指针完成的，第一个指针指向存储有接口类型信息表的地址，接口类型信息表包含已存储的类型信息以及接口所包含的一组方法，第二个指针指向实际存储值的地址。


#### 5.1 接口Interface
下面的例子中我们定义了human是一个接口对象，包含了一个speak()函数，因此我们可以认为所有具有speak函数（行为）的类型都可以称之为human, 这也就表示一个对象可能同时符合多个接口的约束条件。


```go
type human interface {
	speak() string
}
```

我们定义了两个对象，分别是一个结构体对象和一个自定义的类型对象， 都实现了speak方法，也就是都覅金额human接口，

```go
type Programmer struct {
	name string
	age  int
}

func (p Programmer) speak() string {
	return fmt.Sprintf("i am a programmer")
}

func (p Programmer) Doing() string {
	return "Coding..."
}

type Doctor string

func (b Doctor) speak() string {
	return "i am a doctor"
}
```

由于不是总可以获取值的地址，因此对于值对象，如果存在方法集合调用，则只能调用值的方法，而指针则可以及调用值的又可以调用指针的

```
|values  |     methods 	   |
|T  	 | 		(t T) 	   |
|*T 	 | (t T) 或者(t *T) |
```


#### 5.2 类型断言

我们可以对于不同的类型利用接口的方式进行统一处理，比如设置传递参数为接口类型，任何符合接口规范的都可以传递到函数中进行处理, 同时利用类型断言和类型转换，我们可以很方便的进行细粒度的类型划分和处理。

```go
func Say(h human) {
	switch h.(type) {
	case Programmer:
		fmt.Printf("Programmer say: %s and i am %s \n", h.speak(), h.(Programmer).Doing())
	case Doctor:
		fmt.Printf("Doctor say: %s\n", h.speak())
	default:
		fmt.Printf("Some one say: %s\n", h.speak())
	}
}
```

另外一个类型断言的实例如下面所示：

```go

package main
import (
    "fmt"
)
func ShowMessage(message interface{}) {
    switch _message := message.(type) {
    case string:
        fmt.Printf("string message: %s\n", _message)
    case int:
        fmt.Printf("int message: %d\n", _message)
    default:
        fmt.Println("Unknown type")
    }
}
func ShowStringMessage(message interface{}) {
    str, ok := message.(string)
    if ok == true {
        fmt.Printf("this is a string message :%s \n", str)
    } else {
        fmt.Println("value is not a string")
    }
}
func main() {
    ShowMessage("hello")
    ShowMessage(12)
    ShowMessage(1.00)
    ShowStringMessage("this is mike")
}
// string message: hello
// int message: 12
// Unknown type
// this is a string message :this is mike

```
ShowStringMessage函数中使用类型断言的方式，对于一个interface类型进行强制转换，如果转换成功则获得转换后的类型，以及ok=true,否则则转换失败返回ok=false,**如果不进行返回值判断，则可能导致运行时的错误**。

#### 5.3 实现sort接口

有时候我们需要实现一些标准库或者第三方库中的接口来满足函数调用的需求，比如为了实现数据的排序操作，标准库sort类中的排序函数Sort可以满足需求，其函数定义如下：

```
func Sort(data Interface)
    Sort sorts data. It makes one call to data.Len to determine n, and
    O(n*log(n)) calls to data.Less and data.Swap. The sort is not guaranteed to
    be stable.

```

该函数可以接收一个interface类型，但是整个接口需要我们满足Less，Swap和Len三个函数行为才能使用，因此为了实现自定义结构的排序操作，我们首先需要满足上述的接口，然后才可以调用sort实现排序操作。

```go
type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
    return len(s)
}
func (s Sequence) Less(i, j int) bool {
    return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
    sort.Sort(s)
    str := "["
    for i, elem := range s {
        if i > 0 {
            str += " "
        }
        str += fmt.Sprint(elem)
    }
    return str + "]"
}
```

> 任何类型都满足interface{}， 因此才使得fmt.Println()这样的函数可以接收任何的类型输入
> 函数的定义**func Println(a ...interface{}) (n int, err error)**， 借助于... 可以接收任意数量的参数输入。

参考：https://www.ardanlabs.com/blog/2015/09/composition-with-go.html


#### 5.4 HTTP接口

在Golang的标准库http模块中，定义了一个接口类型Handler，任何满足Handler接口的都可以来处理HTTP请求，这也在web编程中经常会遇到的一个接口实现，该接口的定义如下：

```golang
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```

其中的ResponseWriter也是一个接口类型，Request则是客户端传递的请求的结构体对象。因此我们可以通过满足Handler接口来实现一个自己的Http处理器, 下面的代码我们利用接口的方式来是实现一个自增计数器的HTTP接口，由于ResponseWriter满足io.Writer接口,我们可以直接使用fmt.Fprintf来将数据传递给ResponseWriter，实现http的响应的返回。


```golang
import (
	"fmt"
	"net/http"
)

type Counter struct {
	n int
}

func (c *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c.n++
	fmt.Fprintf(w, "Counter=%d\n", c.n)
}

func main() {
	c := new(Counter)
	http.Handle("/counter", c)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

```
通过上述代码我们可以启动一个web服务器，并监听在127.0.0.1:8080端口上,如果想要查看web的响应可以通过访问http://127.0.0.1:8080/counter来检查是否正确。

