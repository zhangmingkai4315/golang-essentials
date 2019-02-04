### 14. 程序文档

go命令提供两种查询程序文档的方式，第一种是直接用```go doc```来查询，第二种是通过与go程序一起自动安装的godoc程序来查询，如果仅仅是查询标准库的话，两者基本上输出的内容比较相似，但是godoc提供一些额外的扩展功能，比如内置的web服务，提供网页版的文档检索，同时支持来自于github.com的网站中程序包的文档的查询，后面我们会依次介绍联众文档管理方式。

#### 14.1 go doc命令

通过命令``go doc`可以来查询相关的程序包文档，我们可以使用go help命令来查看如何使用go doc

```
$ go help doc

        go doc
                Show documentation for current package.
        go doc Foo
                Show documentation for Foo in the current package.
                (Foo starts with a capital letter so it cannot match
                a package path.)
        go doc encoding/json
                Show documentation for the encoding/json package.
        go doc json
                Shorthand for encoding/json.
        go doc json.Number (or go doc json.number)
                Show documentation and method summary for json.Number.
        go doc json.Number.Int64 (or go doc json.number.int64)
                Show documentation for json.Number's Int64 method.
        go doc cmd/doc
                Show package docs for the doc command.
        go doc -cmd cmd/doc
                Show package docs and exported symbols within the doc command.
        go doc template.new
                Show documentation for html/template's New function.
                (html/template is lexically before text/template)
        go doc text/template.new # One argument
                Show documentation for text/template's New function.
        go doc text/template new # Two arguments
                Show documentation for text/template's New function.

```
根据上面的描述信息，假如我们想要查看fmt.Println的文档可以使用下面的方式来查询。
```
$ go doc fmt.Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.

```
另外默认情况下直接使用go doc将列举当前位置的库的文档信息，后面我们可以看到如何去编写自己写的库的文档信息。
 
#### 14.2 godoc命令

go语言安装后会自动的同时安装godoc命令，这个命令有一个内置的web服务，可以在本地启动一个标准库的文档服务器，用于查询当前标准库的相关内容。使用下面的命令启动该web服务：

```sh
godoc.exe -http=137.0.0.1:8080
```
可以通过web服务中的搜索框或者Packages来查询，同时godoc本身支持命令行的查询，但是与go doc有所不同，包和函数中间是不包含任何点号连接的，以空格分割。

```go
$ godoc fmt Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline
    is appended. It returns the number of bytes written and any write error
    encountered.

```



#### 14.3 使用godoc索引

假如自己写了一个简单的go语言的库，想要别人能够通过godoc检索到，该如何操作？

其实很简单，这里我们以当前文件夹中的mylib库为例子，我们推送到git上后，会获得该库的一个访问路径，以下是实例的连接：
```
https://github.com/zhangmingkai4315/golang-essentials/tree/master/13-documentation/mylib
```

我们将该链接放置到godoc中的搜索框后点击Go按钮即可，系统会自动的重定向到一个新的URL，而且所有的文档信息会自动的显示出来。这时候任何人如果通过搜索框都可以查询到该库的相关信息。

同时我们可以使用godoc命令行来查询远程Web链接下的文档信息：

```sh
 godoc github.com/zhangmingkai4315/golang-essentials/13-documentation/mylib

```

#### 14.4 编写文档

如何编写文档，可以通过学习一些标准库的程序编写规范，这里我们以标准库中的errors库作为实例, 代码文档一般包含以下几个部分，当然这些文档都不是必须存在的，但是为了更好的提升代码的质量，特别是协同编程的情况下，提高代码可阅读的能力，建议尽量包含以下信息：

- 头部版权信息
- 包通用描述信息
- 针对具体函数或者类型的定义信息（第一个单词必须以类型名称或者函数名开始）

如下面的errors实例所示：

```
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errors implements functions to manipulate errors.
package errors

// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

```
每一个包都尽量包含包描述信息，对于一些包的描述信息如果比较长，包含一些示例的话，可以创建一个单独的文件用于保存这些信息，比如fmt包中存在一个独立的文件doc.go文件来保存这些信息，具体实现可访问链接[doc.go](https://golang.org/src/fmt/doc.go)， doc.go文件除了注释信息以外，只包含一个单独的```package fmt```的声明


对于外部可访问的结构体或者函数对象，编写文档注释的时候要以名称作为首个单词， 比如下面的示例所示：

```go
// Compile parses a regular expression and returns, if successful,
// a Regexp that can be used to match against text.
func Compile(str string) (*Regexp, error) {
 ...
}

```

对于一组对象的声明go程序允许按组来进行文档注释的编写，如下面所示我们创建了多个错误对象，这些对象由于都可以归于一类，因此我们在定义文档的时候仅需要定义一次即可
```go
// Error codes returned by failures to parse an expression.
var (
ErrInternal = errors.New("regexp: internal error")
ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
...
)
```

如果我们查询标准库的话，会看到很多库都带有example示例，这些示例其实也是通过解析代码库中的相关代码提取出来的，而且这些示例提取后，可以直接形成可执行的go代码。比如golang/example示例中的stringutil包，为了提供示例代码我们可以编写example_test.go文件：
```go
package stringutil_test

import (
    "fmt"

    "github.com/golang/example/stringutil"
)

func ExampleReverse() {
    fmt.Println(stringutil.Reverse("hello"))
    // Output: olleh
}
```

文件的package名称为包名称加test，同时示例函数必须以Examle开头，对任何函数或者类型的示例后面跟函数名称和结构体名称。
示例代码中的```Output: olleh```会自动被转换输出

![](https://blog.golang.org/examples/reverse.png)

同时当我们使用go test的时候，example代码也会被自动执行测试
```sh
$ go test -v
=== RUN TestReverse
--- PASS: TestReverse (0.00s)
=== RUN: ExampleReverse
--- PASS: ExampleReverse (0.00s)
PASS
ok      github.com/golang/example/stringutil    0.009s
```

这里的example的测试比较实际的输出和Output中定义的输出，**如果不包含Output也就不会执行相关测试**

多个输出的时候，按行完成output的编写：

```golang

    fmt.Println(people)
    sort.Sort(ByAge(people))
    fmt.Println(people)

    // Output:
    // [Bob: 31 John: 42 Michael: 17 Jenny: 26]
    // [Michael: 17 Jenny: 26 Bob: 31 John: 42]
```
