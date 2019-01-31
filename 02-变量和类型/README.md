
### 4. Go语言基础

#### 4.1 编写第一个Go语言程序

github上面创建一个基础项目比如：
https://github.com/zhangmingkai4315/golang-essentials.git

```shell
go get github.com/zhangmingkai4315/golang-essentials
cd $GOPATH/src/github.com/zhangmingkai4315/golang-essentials
```
hello world的go 语言版本， 包含了一个main()函数，这也是程序的入口函数，执行从此处开始后，直到函数结束后退出。

```golang

package main
import (
    "fmt"
)
func main() {
    fmt.Println("hello world")
}

```

尝试执行go run 或者go build或者go fmt并检查输出内容。


#### 4.2 使用golang标准包

https://golang.org/pkg/ 包含所有已经内置的标准包内容，除了标准包以外还有一些第三方的包提供下载使用。可以直接在线查询某一些包的使用方法：比如https://godoc.org/fmt
或者使用内置的命令检查整个的标准包或者某一些函数的介绍

```
go doc fmt
go doc fmt.Println
```

如果无法翻墙使用godoc也可以直接使用下面的命令启动一个本地的web服务
```
 godoc -http=:8080
```
godoc.org网站包含除了标准库之外的第三方库


使用import将导入包到运行环境，使用包名称加函数名的方式调用包中的函数比如```fmt.Println()```

#### 4.3 参数和类型

我们可以通过go doc fmt.Println查看函数的定义和用法介绍，其中函数定义如下：
```func Println(a ...interface{}) (n int, err error)``` 其中函数的参数与很多语言不同之处在于参数名称在前，类型在后，使用三个点号代表可以接受任意数量的参数个数，返回值这里包含两个返回值，一个整型n代表了实际打印的内容长度，另一个err代表了错误类型error是否包含错误信息，这些都是标准库中定义的类型，当然自己定义函数的时候也可以定义自己创建类型。


```golang

func main(){
    fmt.Println("hello",12,false)
    n, err := fmt.Println("hello world")
    fmt.Println(n, err)

    n1, _ := fmt.Println(12)
    fmt.Println(n1)


}
// hello 12 false
// hello world
// 12 <nil>
// 12
// 3
```

参数中的interface代表任何类型，也就是可以允许传递任意类型对象到函数中执行。

对于不需要的返回值可以使用_来代替，go语言中要求任何变量声明后却不使用都是不允许的，这在程序编译期间就会报错退出。
 
 #### 4.4 Go语言中的变量定义
 
 之前的函数使用中我们定义了n和n1分别代表我们自己的变量名称，但是go语言中对于变量名以及类型名声明是有要求的比如如下的格式：
 
 ```
identifier = letter { letter | unicode_digit } 
 ```
 
 另外下面已经定义的关键词不能被作为名称使用：
 
 ```

break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
 ```
 
 一些golang操作符定义如下：
 
 ```
 
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
 ```
 
通过这些操作符号和关键词以及我们自定义的名称，组成了所有的go程序的代码。

类型声明的方式,如下是两种声明的方式，一种是直接的使用:=符号来定义变量，这种方式定义的变量只在当前的作用域下面有效，而var则可以在整个的包范围内有效（尝试在main函数外定义变量，并在main函数中使用）

```golang

package main
import (
    "fmt"
)
func main() {
    name := "Mike"
    var city string
    city = "Beijing"
    fmt.Printf("%s in %s", name, city)
}
// Mike in Beijing

```

变量声明中的缺省值设置比如设置一个var a int 缺省设置a=0，而对于bool类型或者string则分别为false和"", 浮点型类型的缺省值为0.0
其他的类型比如指针，函数以及切片,map和channel的缺省都为nil
https://golang.org/ref/spec#The_zero_value


#### 4.5 类型


golang包含的基础类型包含：
- 数值类型  
- 函数类型
- 数组类型
- 切片类型
- interface类型
- map类型
- channel类型
- 结构体类型
- 指针类型
- 自定义类型

这些类型的声明和使用如下面的代码所示：

```golang

package main
import (
    "fmt"
)
type Person struct {
    name string
}
type MyInterface interface {
    Read(message string) string
}
func (t Person) Read(message string) string {
    return fmt.Sprintln(t.name + " is reading " + message)
}
func main() {
    t1 := 64
    fmt.Printf("%d is %T\n", t1, t1)
    t2 := "mike"
    fmt.Printf("%s is %T\n", t2, t2)
    t3 := false
    fmt.Printf("%t is %T\n", t3, t3)
    t4 := &t1
    fmt.Printf("%v is %T\n", t4, t4)
    t5 := []string{"i", "really", "like", "golang"}
    fmt.Printf("%v is %T\n", t5, t5)
    t6 := [4]string{"i", "really", "like", "golang"}
    fmt.Printf("%v is %T\n", t6, t6)
    t7 := map[int]string{
        1: "mike",
        2: "alice",
        3: "anjoue",
    }
    fmt.Printf("%v is %T\n", t7, t7)
    t8 := struct {
        name string
        age int
    }{"mike", 12}
    fmt.Printf("%v is %T\n", t8, t8)
    t9 := func(message string) {
        fmt.Println(message)
    }
    fmt.Printf("t9 is %T\n", t9)
    type Mytype int8
    var mytype Mytype
    fmt.Printf("mytype is %T\n", mytype)
    p := Person{"Mike"}
    fmt.Println(p.Read("novel"))
}

```

https://golang.org/ref/spec#Types


上面的实例中使用了fmt包，该包中包含了一些格式化的函数，具体的文档可以参考https://golang.org/pkg/fmt/，另外使用%v可以打印一些结构体的内容，如果%+v可以包含相关域的名称,以下就是使用使用%v的时候针对不同类型等效的表达方式：

```

bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p

struct:             {field0 field1 ...}
array, slice:       [elem0 elem1 ...]
maps:               map[key1:value1 key2:value2 ...]
pointer to above:   &{}, &[], &map[]

```

类型转换，对于golang静态语言，无法实现不同类型的直接转换，但是通过类型转换，我们可以将不同类型的值之间进行类型的动态修改，

https://tour.golang.org/basics/13
https://golang.org/doc/effective_go.html#conversions

对于interface类型，比如传递的参数为interface，有时候我们需要根据不同的类型执行不同函数操作，这时候可以利用类型断言的方式进行判断，判断的方式如下

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
ShowStringMessage函数中使用类型断言的方式，对于一个interface类型进行强制转换，如果转换成功则获得转换后的类型，以及ok=true,否则则转换失败返回ok=false,如果不进行返回值判断，则可能导致运行时的错误。

#### 4.8 练习

略