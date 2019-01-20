### 8 . 编写应用程序

#### 8.1 使用Go语言编写JSON序列化应用

JSON 也就是**J**ava**S**cript **O**bject **N**otation的缩写， 被广泛的用于当今的互联网应用之间存储和交换数据，尽管源自于JavaScript语言，但是属于一个跨平台支持的数据格式，类似于传统的XML，但是更加简洁，使用方便。Go语言标准库中提供了字符串和JSON对象的序列化操作的接口，可以直接调用完成序列化和反序列化。如果还不了解JSON，可以参考[W3C介绍](http://www.w3school.com.cn/json/json_syntax.asp). 

首先json的序列化和反序列化的接口所在位置位于"encoding/json"中，这个"encoding"还包含了比如xml，csv,gob等多种格式的序列化编码库, 如果感兴趣可以查阅[encoding官方文档](https://godoc.org/encoding)

```go
import (
	"encoding/json"
)
type student struct {
	Name        string
	Age         int
	privateInfo string
}
```

我们定义一个结构体用于存储我们的数据，JSON序列化就是将我们存储的结构体数据转换为基本的[]byte类型的形式，如果没有指定转换方式（后面会介绍），则按照默认的名称进行转换。如下面所示：

```go
func main() {
	s1 := student{
		Name:        "mike",
		Age:         26,
		privateInfo: "i like music",
	}
	s2 := student{
		Name:        "alice",
		Age:         30,
		privateInfo: "i hate mike",
	}

	school := []student{s1, s2}

	schoolInfo, err := json.Marshal(school)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("school infomation is %v", string(schoolInfo))
	}
	// school infomation is [{"Name":"mike","Age":26},{"Name":"alice","Age":30}]
}
```



在做序列化的时候，要注意的是结构体中的属性名称如果为小写的话，则过程中不会被转换，输出的字符串中将不会包含此部分的信息，这也是为什么没有在输出的信息中，找到privateInfo的信息。

JSON反序列化指的是将json字符串转换为指定类型的过程，这个过程中和序列化一样简单：

```go
	school := []student{}
	rawdata := `[{"Name":"mike","Age":26},{"Name":"alice","Age":30}]`

	err := json.Unmarshal([]byte(rawdata), &school)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("school infomation is %v\n", school)
	}
	// school infomation is [{mike 26 } {alice 30 }]
```



可以看到我们的数据被成功的反序列化到结构体中存储，但是有时候我们想修改部分字段的名称，比如json字符串中属性名称改为小写，该如何处理，毕竟按照之前的介绍小写属性名不会被转换，这时候我们就需要借助于标签的方式进行显式的声明转换，如下面的代码所示：

```go
type studentWithLabel struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	privateInfo string `json:"private"` // useless
}

rawdata2 := `[{"name":"mike","age":26,"private":"secret information"},{"name":"alice","age":30}]`

	studentLabelArray := []studentWithLabel{}
	err = json.Unmarshal([]byte(rawdata2), &studentLabelArray)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("school infomation is %v\n", studentLabelArray)
	}

	// school infomation is [{mike 26 } {alice 30 }]
```

上述的代码中使用json标签进行转换的声明，实现了数据的成功反序列化，同时可以看到尽管我们标记了private为转换标签，也在字符串传递了对应的数据，这种转换也是不能实现的，这就是go语言中对于私有数据的一种处理方式，使其对外完全的不可见。

#### 8.2 官方源码程序

经常阅读官方的源码有助于深入的学习Go语言并理解程序的执行逻辑，比如下面我们就从最简单的fmt.Println这个函数看一下，官方是如何实现的。这个函数其实简单的封装了另一个函数Fprintln的调用，参数直接传递下去。

```go
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}
```

继续追踪Fprintln可以看到函数的签名如下,该函数接收的第一个参数是一个io.Writer的接口类型，该类型支持任何具有Write方法是实现的对象类型，因此os.Stdout必定是一个具有Write方法的对象。

```go
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
```

```go
// package io
type Writer interface {
  Write(p []byte) (n int, err error)
}

```

通过查询os的包可以看到，其中的定义如下所示, NewFile返回的是一个os.File类型，该类型包含了一个Write方法，用于写入数据到文本中。感兴趣的可以继续看一下os.NewFile的实现，这里可能就会涉及到一些跨操作系统的底层实现，而且不同操作系统在执行的时候操作也各不相同。

```
var (
	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)

func (f *File) Write(b []byte) (n int, err error)
```

通过上面的例子可以看到，仅仅为了实现打印数据到终端，实际的程序确执行了如此多的操作，通过理解这些背后的处理逻辑，可以帮助更加深刻的理解Go语言程序。

#### 8.3 排序操作

排序是经常用到的一种程序编写方式，如果学过算法的话，对于排序肯定不会陌生，Go程序语言在底层已经帮助实现了基本的排序算法，因此只需要直接调用即可，程序库会根据数据自动选择最佳的排序算法进行处理。

```go
import (
	"fmt"
	"sort"
)
func main() {
    arr := []int{10, 2, 3, 42, 4, 3, 2, 4, 21}
    sort.Ints(arr)
    fmt.Println(arr)
    // [2 2 3 3 4 4 10 21 42]

    arrString := []string{"mike", "alice", "bob"}
    sort.Strings(arrString)
    fmt.Println(arrString)
    // [alice bob mike]
}
```



Go语言程序在1.8之前，如果需要对于自定义排序方式需要实现一个sort的接口，这个接口要求必须实现三个方法，Len, Swap以及Less，具体的方式如下：

```go
type student struct {
	name  string
	age   int
	score int
}
type ByScore []student

func (b ByScore) Len() int { return len(b) }
func (b ByScore) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ByScore) Less(i, j int) bool {
	return b[i].score < b[j].score
}
func main() {
	school := []student{
		{"Mike", 12, 96},
		{"Alice", 13, 100},
		{"Bob", 12, 60},
	}
	fmt.Printf("Unsorted student array is %v\n", school)
	// Unsorted student array is [{Mike 12 96} {Alice 13 100} {Bob 12 60}]
	sort.Sort(ByScore(school))

	fmt.Printf("Sorted student array is %v\n", school)
	// Sorted student array is [{Bob 12 60} {Mike 12 96} {Alice 13 100}]
}
```

在1.8之后的版本中，简化了编写代码的方式，可以使用另外一种方式实现，代码量更少，只需要传递上面的Less函数到sort.Slice函数中即可实现代码的排序操作：

```go
school2 := []student{
    {"Mike", 12, 96},
    {"Alice", 13, 100},
    {"Bob", 12, 60},
}

sort.Slice(school2, func(i, j int) bool {
    return school2[i].score > school2[j].score
})
fmt.Printf("Sorted student2 array by score is %v\n", school)
// Sorted student2 array by score is [{Bob 12 60} {Mike 12 96} {Alice 13 100}]
```

如果感兴趣的话， 可以查看以下官方的[Sort源码](https://golang.org/src/sort/sort.go)，里面有各种排序的具体实现方式，比如快速排序，堆排序，插入排序等等。

#### 8.4 使用Go语言编写密码加密和验证

不管是web系统还是客户端系统，只要涉及到用户登入的，大部分都会遇到存储用户密码的情况，如何存储是一个比较重要的问题，因为之前有报道某一些企业安全意识比较差，存储用户密码时候以明文存储，导致一些数据库被泄露后用户信息完全的被公开。

为了安全的存储用户的密码信息，推荐的一个库是bcrypt库，该库提供密码的加密和解密(验证)操作，同时可以有效的防止彩虹表的破解（通过Hash值反查询密码）每次执行都会输出不同的信息。但是由于并非官方的标准库中内容，因此为了使用这个库，我们需要将库代码下载到本地使用：

```
 go get golang.org/x/crypto/bcrypt
```

使用的方式如下面所示：

```go
package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// go get golang.org/x/crypto/bcrypt
func main() {
	s := "password"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
	// $2a$04$48e/0SGKpjsJ/nD6G4S83ey65T1.kCXdplusBEX/iIkBzlegGi.b6
	// 每次即便是输入相同，输出也不同

	confirmPassword := "password"

	err = bcrypt.CompareHashAndPassword(bs, []byte(confirmPassword))
	if err != nil {
		panic(err)
	}
	fmt.Println("password is correct")

}

```