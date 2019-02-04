
### 2. Go语言基础

#### 2.1 编写第一个Go语言程序

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


#### 2.2 使用golang标准包

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

#### 2.3 参数和类型

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

参数中的interface{}代表任何类型（不包含任何函数的接口类型），也就是可以允许传递任意类型对象到函数中执行。对于不需要的返回值（直接丢弃）可以使用 _ 来代替，go语言中要求任何变量声明后却不使用都是不允许的，这在程序编译期间就会报错退出。

对于上面的程序我们可以看到main()函数后面紧跟了一个{大括号，这是Go语言的一种使用规范，假如将{移到下一行中，则程序会报错无法执行。主要是Go语言和C语言一样，会使用；来终止一行代码，但是这分号并不会出现在代码中，而是自动管理和添加，插入分号的规则很简单：
**如果每一行中最后的字符是一个标识符比如int, float,或者一个数字以及string类型再或者是break, continue, fallthrough return ++ -- ) } 这些符号，则追加一个分号来结束此行。**。根据此规则，如果我们main()后面没有任何的字符，则直接添加;结束了代码，编译肯定失效了。

 
 #### 2.4 Go语言中的变量定义
 
 之前的函数使用中我们定义了n和n1分别代表我们自己的变量名称，但是go语言中对于变量名以及类型名声明是有要求的比如如下的格式：
 
 ```
identifier = letter { letter | unicode_digit } 
 ```

 
 Go语言中定义名称一般不会使用下划线来定义，使用类似于MixedCaps或者mixedCaps的形式，另外下面已经定义的关键词不能被作为名称使用：
 
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


#### 2.5 类型


golang包含的[基础类型](https://golang.org/ref/spec#Types)包含：
- 数值类型  
- 数组类型
- 切片类型
- map类型
- interface类型
- 函数类型
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

