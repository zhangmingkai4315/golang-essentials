#### 12.1 go doc命令

go语言内置的doc子命令可以通过命令行来查询文档，我们可以使用go help命令来查看如何使用go doc

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
 
#### 12.2 godoc命令

go语言安装后会自动的同时安装godoc命令，这个命令有一个内置的web服务，可以在本地启动一个标准库的文档服务器，用于查询当前标准库的相关内容。使用下面的命令启动该web服务：

```sh
godoc.exe -http=127.0.0.1:8080
```
可以通过web服务中的搜索框或者Packages来查询，同时godoc本身支持类似于go doc查询库包的命令

```go
$ godoc fmt Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline
    is appended. It returns the number of bytes written and any write error
    encountered.

$ go doc fmt Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.

```

#### 12.3 使用godoc索引

假如自己写了一个简单的go语言的库，想要别人能够通过godoc检索到，该如何操作？

其实很简单，这里我们以当前文件夹中的mylib库为例子，我们推送到git上后，会获得该库的一个访问路径，以下是实例的连接：
```

```



[Golang官方文档](https://golang.org/doc/)
[godoc.org索引](https://godoc.org/)
