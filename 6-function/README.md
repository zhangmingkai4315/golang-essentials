### 1. Golang函数

函数在几乎所有的编程语言中都占用很重要的地位，利用函数可以将代码进行封装，用于提供特定的逻辑处理和通过复用来消除重复代码。Go语言中的函数和其他语言中的函数创建方式基本相似，包含函数名称，函数参数以及函数的返回值组成。

#### 1. 函数的创建

以下是几种创建函数的方式，这些函数具有不同的参数以及不同的返回值。其中Go语言中加入了很多特性来支持更灵活的编写代码，比如多值返回，缺省返回以及后面我们可以看到的defer和函数异常处理等。

```go
// Foo is a zero arguments function
func Foo() {
	fmt.Println("Foo: function")
}

// FooWithArgs function in go pass by value
func FooWithArgs(s string) {
	fmt.Printf("FooWithArgs: %s\n", s)
}

// FooWithMultiArgs with multi arguments
func FooWithMultiArgs(s string, prefix bool) {
	if prefix == true {
		fmt.Printf("FooWithArgs: %s\n", s)
	} else {
		fmt.Printf("%s\n", s)
	}

}

// FooWithArgsAndReturn with string return
func FooWithArgsAndReturn(s string) string {
	return fmt.Sprintf("FooWithArgsAndReturn: %s\n", s)
}

// FooWithArgsAndMultiReturn will return multi return
func FooWithArgsAndMultiReturn(s string) (string, error) {
	if s == "error" {
		return "", errors.New("FooWithArgsAndMultiReturn: Error")
	}
	return fmt.Sprintf("FooWithArgsAndMultiReturn: %s\n", s), nil
}

// FooWithArgsAndMultiDefaultReturn will return multi return with some default
func FooWithArgsAndMultiDefaultReturn(s string) (message string, err error) {
	if s == "error" {
		err = errors.New("FooWithArgsAndMultiDefaultReturn: Error")
		return
	}
	message = fmt.Sprintf("FooWithArgsAndMultiDefaultReturn: %s\n", s)
	return
}
```

函数在使用的时候可以通过以下的方式处理：

```go
func main() {
	Foo()
	// Foo: function

	FooWithArgs("Hello")
	// FooWithArgs: Hello
	FooWithMultiArgs("Hello", false)
	// Hello

	fmt.Print(FooWithArgsAndReturn("Hello world"))
	// FooWithArgsAndReturn: Hello world

	_, err := FooWithArgsAndMultiReturn("error")
	if err != nil {
		fmt.Println(err)
	}
	// FooWithArgsAndMultiReturn: Error

	_, err = FooWithArgsAndMultiDefaultReturn("error")
	if err != nil {
		fmt.Println(err)
	}
	// FooWithArgsAndMultiReturn: Error
}
```



#### 2. 可变参数

我们在之前上面的示例中定义的函数函数的参数都是固定的，也就是我们必须传递指定数量的参数到函数中，否则程序在编译的时候就会报错。但是对于我们常用的fmt.Println函数，却可以传递可变数量的参数，这是如何实现的呢？我们可以通过函数的签名来看一下：

```go
go doc fmt.Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.

```

函数签名中使用了...方式来接收参数并传递到变量a中，我们可以仿照这种方式来创建一个可以接收多个整数的函数来计算一个整数的和。下面的示例中针对数组我们可以使用反向拆分的方式将数组中的元素拆分成单独的元素传递给函数。

```go
// SumWithVariadicArgs with variadic parameter
func SumWithVariadicArgs(a ...int) int {
	sum := 0
	for _, i := range a {
		sum = sum + i
	}
	return sum
}

fmt.Printf("sum(10,20)=%d\n", SumWithVariadicArgs(10, 20))
fmt.Printf("sum(10,20,30)=%d\n", SumWithVariadicArgs(10, 20, 30))
arr := []int{10, 20, 30, 23, 23}
fmt.Printf("sum(10,20,30,23,23) = %d\n", SumWithVariadicArgs(arr...))
// sum(10,20)=30
// sum(10,20,30)=60
// sum(10,20,30,23,23) = 106
```



#### 3. defer函数

defer关键词用于推迟一个函数的执行，直到当前函数退出执行或者遇到异常退出，通过defer我们可以确保函数在执行结束或者出现异常后去能够去执行某一些操作，比如删除临时文件，关闭文件，关闭数据库连接等操作。

```go
func deferedFunc() {
	fmt.Println("this function is defered!")
}

func funcWithDefer() {
	defer deferedFunc()
	fmt.Println("this is function with defer")
}
funcWithDefer()
// this is function with defer
// this function is defered!
```

一个稍微复杂的例子如下面的所示, 通过defer将完成程序的收尾工作（关闭文件），即便是整个的执行期间遇到异常，比如写文件的时候出现异常函数退出，在退出之前也会保证函数能够被安全的关闭，这也是defer函数经常被使用的一种场景。

```go
func main() {
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}
func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}
func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
    fmt.Println("closing")
    f.Close()
}
```

对于一个函数中存在多个defer的时候，执行顺序将按照入栈的方式依次执行(FILO先入后出的方式)，示例代码如下：

```go
func FuncWithMultiDefer(){	
    fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d ", i)
	}

	fmt.Println("done")
}
	// counting
	// done
	// 9 8 7 6 5 4 3 2 1 0
```

