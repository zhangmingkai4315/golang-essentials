### 13. 错误处理

在java, paython以及javascript编程语言中，如果需要处理错误，捕获异常都有相似的语法结构，使用类似于try.catch语法来实现异常的捕获和处理。在Go语言中没有这样的语法定义，官方的文档中曾经介绍过这样的设计初衷，开发人员觉得这样的设计（try...catch）不仅使得代码过于负载，而且使得程序员也会倾向于标记过多的错误，使得错误类型太多，最终程序变得不够简练。

Go语言中本身使用了多值返回使得错误报告变得不仅简单，而且这种错误返回感觉已经变成了语言的一种规范，另外go语言中使用一种recover机制来捕获任何出现问题的代码。后面将给出error的具体例子。

#### 13.1 error类型

error类型是程序语言的内置类型，无需引入其他的包来使用，其定义如下面所示, error类型本身是一种接口类型，也就是任何实现了Error()方法的类型不管是结构体，还是整形，字符类型都属于错误类型，
```go
type error interface {
    Error() string
}
```

我们可以通过一个标准库errors来看一下错误类型的使用，errors库非常简单，主要用来创建一个错误类型，后期写代码会经常使用到它，其源代码仅仅十几行，如下所示：

```golang

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

errors库通过一个内置的结构体errorString保存错误的描述信息，该结构体通过实施Error()方法来适配error的接口类型，这样我们只需要使用New()方法就可以返回一个满足error接口类型的对象，也就是一个error类型。

我们创建error的方式除了上面的errors.New()外，fmt包中也存在Errorf函数用来产生error类型的对象，不过如果查看fmt的源码就可以看到该函数的定义如下， 还是借用了errors.New()函数来创建错误对象。
```go
func Errorf(format string, a ...interface{}) error {
	return errors.New(Sprintf(format, a...))
}
```


#### 13.2 错误检查

Go语言中绝大多数的函数都会返回相关错误信息，如果没有错误则返回的错误信息为nil,即便是我们之前经常使用的fmt.Println函数也有错误信息的返回，只是我们很少用来检查错误，但是大部分情况下，比如建立连接，打开文件，写入信息，最好是检查错误信息的状态，从而减少程序出错的可能性。

```go
f, err := os.Create("temp.log")
if err != nil {
    fmt.Println(err)
    return
}
defer f.Close()

r := strings.NewReader("hello world")
_, err = io.Copy(f, r)
if err != nil {
    fmt.Println(err)
}
```
上述代码中使用os模块来打开文件并写入一些信息到文件中，对于错误检查的方式基本都是直接与nil进行比较即可，如果该处存在错误，则编写错误的处理函数即可。

#### 13.3 错误处理

Go语言中对于错误检查，一旦我们发现程序执行出现错误，一般可以通过以下集中方式来进行处理：
- 使用fmt模块打印错误
- 使用log模块打印错误
- 使用log模块打印并退出
    - Panic模式
    - Fatal模式

使用fmt直接打印错误信息，之前的例子已经见过了，此处不再介绍，我们看下log模块的处理方式，log模块顾名思义就是记录日志，标准库提供的log模块具有多个配置选项，我们可以将日志打印到终端或者输出到指定文件，如下面所示, 我们将程序输出到创建的日志文件中，输出信息默认会追加告警出现的时间。这里我们只是简单的打印程序，打印结束后程序继续执行并最终退出main函数。

```go
func badFunc() (string, error) {
	return "", errors.New("error occure")
}

func main() {
	f, err := os.Create("temp.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		f.Close()
		fmt.Println("close file success")
	}()

	log.SetOutput(f)
	defer fmt.Println("not working")
	_, err = badFunc()
	if err != nil {
		log.Println(err)
	}
}

```

对于log的异常退出有两种方式，分别是panic和fatal模式,如下所示为Fatal模式，这种模式下，程序输出完成错误信息后，会调用os.Exit()来退出，并且所有之前设置的defer函数将被跳过不会继续执行。

```go
	_, err = badFunc()
	if err != nil {
		log.Fatalln(err)

	}
```
另外一种方式为panic，这种方式类似于输出后直接调用内置的panic函数,不会影响已经定义的defer函数的执行，所以比较推荐这种方式执行错误的处理。

```go
	_, err = badFunc()
	if err != nil {
		log.Panicln(err)
	}
```
#### 13.4 defer和recover

关于defer函数之前已经简单的介绍过，这里举一个具体的实例来加深对于defer的执行逻辑的理解,该程序的输出结果为2，是否能够理解为何输出结果不是1？

```go
func c() (i int) {
	defer func() {
		i++
	}()
	return 1
}
func main() {
	x := c()
	fmt.Println(x)
}

```
首先要理解defer函数是于return之后执行，本身程序返回的是i的值，返回后再次执行defer，i的值由1增加到2，函数返回值也就增加到值2.

recover函数是go语言中内建的函数，仅仅用于defer函数中，如果没有任何错误和异常的发生，recover函数简单的返回nil，因此我们可以通过判断来捕获并处理任何函数执行中的错误。如下面的代码所示：

```go
func recoveredFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("function panic is somewhere")
			// function panic is somewhere
		}
	}()

	panic("kernal panic")
}

```
该函数由于panic原因退出程序执行，但是我们在defer函数中加入的recover函数成功的捕获到了该异常，如果尝试打印r的话可以看到r就是实际panic的相关信息。另外如果我们捕获并处理了panic，则程序会继续执行该函数之外的后面的代码，恢复执行而不是直接退出。


#### 13.5 自定义错误

有时候我们需要封装自己的错误信息，比如加入一些额外的信息来存储用于将来的错误判断和处理，而不是简单的string类型，我们可以通过定义自己的结构体对象，并实现error接口即可，如下面所示：

```go
type MyError struct {
	funcName    string
	funcLine    int
	description string
}

func (myerror *MyError) Error() string {
	return fmt.Sprintf("Error: funcname=%s line=%d", myerror.funcName, myerror.funcLine)
}

func myFunc() error {
	return &MyError{funcName: "myFunc", funcLine: 1}
}

```