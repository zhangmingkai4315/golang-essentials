### 1. Golang函数

函数在几乎所有的编程语言中都占用很重要的地位，利用函数可以将代码进行封装，用于提供特定的逻辑处理和通过复用来消除重复代码。Go语言中的函数和其他语言中的函数创建方式基本相似，包含函数名称，函数参数以及函数的返回值组成。

#### 1. 函数的创建

以下是几种创建函数的方式，这些函数具有不同的参数以及不同的返回值。其中Go语言中加入了很多特性来支持更灵活的编写代码，比如多值返回，缺省返回以及后面我们可以看到的defer和函数异常处理等。下面分别对于函数的一些使用方式进行举例：

- 无参数
- 单一参数输入
- 多参数输入
- 单返回值
- 多返回值
- 缺省返回值

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

函数在使用的时候可以通过以下的方式处理

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

另外函数本身也是一种类型，因此我们可以通过常规类型声明的方式来声明或者定义函数比如下面的例子：

```go
	f := func(name string) string {
		return fmt.Sprintf("My name is %s", name)
	}
	f("mike")
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

#### 4. 接口Interface

接口是go语言中用来定义一些具有相同行为的一种类型，比如下面的例子中我们定义了Human是一个接口对象，包含了一个speak()函数，因此我们可以认为所有具有speak函数（行为）的类型都可以称之为Human, 这也就表示一个对象可能同时符合多个接口的约束条件。

```go
type Human interface {
	speak() string
}
```

我们定义了两个对象，分别是一个结构体对象和一个自定义的类型对象， 都实现了speak方法，也就是都覅金额Human接口，

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

这样我们可以对于不同的类型利用接口的方式进行统一处理，比如设置传递参数为接口，任何符合接口规范的都可以传递到函数中进行处理, 代码如下， 同时利用类型断言和类型转换，我们可以很方便的进行细粒度的类型划分和处理。

```go
func Say(h Human) {
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
>
> 函数的定义**func Println(a ...interface{}) (n int, err error)**， 借助于... 可以接收任意数量的参数输入。

参考：https://www.ardanlabs.com/blog/2015/09/composition-with-go.html



#### 5. 匿名函数

匿名函数允许我们实现函数定义和立即执行，适用于仅仅调用一次的情况，不需要额外的执行函数定义操作，比如一些函数中的defer函数操作一般可以使用匿名函数实现，如下面的例子所示：

```go
func defered() {
	defer func() {
		fmt.Println("this is defered function")
	}()
	fmt.Println("normal function call")
}

func main() {
	func() {
		fmt.Println("this is a anonymous function")
	}()
	defered()
}

// this is a anonymous function
// normal function call
// this is defered function
```



#### 6 传递和返回函数

函数作为go语言中的一类公民( [Wiki](https://en.wikipedia.org/wiki/First-class_citizen) ,具有所有其他对象类型可以执行的操作，比如作为参数传递，或者作为返回值返回，被修改或者被重新的分配变量等）, 我们可以通过下面的实例来看一下如何将函数作为参数进行传递，以及如何返回一个函数:

```go
func Timeit(f func()) func() {
	return func() {
		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed: %v ---n", time.Since(t))
		}(time.Now())
		f()
	}
}
```

通过上面的例子我们看到，函数的传递和返回都很简单，标准库中有很多类似的用法，特别是http库中或者第三方web框架中大量的使用函数的返回和传递，实现流程的灵活控制和绑定执行。

##### 6.1 回调函数

回调函数也是利用函数作为参数实现的一种逻辑处理方式，函数接收一个函数作为参数，并执行该函数的同时，加入一些自身的逻辑处理，比如下面的例子中 使用even函数来计算所有偶数的和，因为之前已经编写了计算整数的和的代码，我们可以直接拿来使用(该实例仅用于介绍回调函数的定义)

```go
func sumInt(arr ...int) int {
	var total = 0
	for _, i := range arr {
		total += i
	}
	return total
}

func even(sum func(...int) int, arr ...int) int {
	newArray := []int{}
	for _, i := range arr {
		if i%2 == 0 {
			newArray = append(newArray, i)
		}
	}
	return sumInt(newArray...)
}

```

##### 6.2 闭包

闭包是一种利用返回函数实现的高级用法，通过闭包我们可以传递一些共享环境和对象，比如下面的累加器，闭包由于返回函数中包含了本身作用域外的对象， 这些对象不会被自动回收, 因此多次执行使用的是相同的作用域对象。

```go
func incfactor(base int) func() int {
	i := base
	return func() int {
		i = i + 1
		return i
	}
}

func main() {
	inc := incfactor(10)
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	// 11
	// 12
	// 13
}

```

另外一个示例是生成斐波那契数列，通过保存两个对象的值实现每次打印一个斐波那契数的方式：

```go
func fibFactoty() func() int {
	x, y := 0, 1
	return func() (r int) {
		r = x
		x, y = y, x+y
		return
	}
}

func main() {
	fib := fibFactoty()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fib())
	}
}
// 0 1 1 2 3 5 8 13 21 34
```



#### 7. 递归函数

递归函数是通过不断的调用自身实现的一种编程方式，由于会调用自身，因此必须设置恰当的退出机制，否则会导致无限循环，直至程序崩溃退出。下面的程序中利用递归的方式计算阶乘， 特别注意如果没有函数中开始的退出判断，函数会不断的计算下去。

```go
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}
func main() {
    fmt.Println(fact(7))
}
```





