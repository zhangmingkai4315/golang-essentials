### 14. 测试

Go语言内置了测试模块[testing](https://godoc.org/testing), 使用该库可以用来编写测试代码，完成程序的测试工作。执行测试的方式可以直接使用```go test ```即可，但是为了能够自动找到对应的文件以及执行编写测试，需要遵循三个基本的测试代码规范：

1. 所有的测试文件需要以**_test.go**结尾
2. 测试文件需要与被测试的文件放于同一个包下面
3. 所有的测试函数都必须按照**func TestXxx(*testing.T)**的编写方式进行编写。Test后面必须以大写字母开头，否则测试被忽略。

#### 14.1 测试模块testing.T

规范3中的函数接收的参数为```testing.T```，T为测试模块中的一个结构体对象，包含一些常用的测试函数用来输出错误信息，管理错误状态以及日志等,常用的方法如下面所示：
```go
func (c *T) Error(args ...interface{}) 
//打印信息并执行Fail
func (c *T) Errorf(format string, args ...interface{})
func (c *T) Fail()  
//标记为失败并继续执行
func (c *T) Log(args ...interface{}) 
//执行测试时，当程序出错或者使用-test.v参数的时候打印信息
func (c *T) Logf(format string, args ...interface{})
func (t *T) Parallel() 
//标记测试为并行执行模式
func (t *T) Run(name string, f func(t *T)) bool 
//运行在单独的goroutine中，并返回执行状态
func (c *T) Skip(args ...interface{}) 
//跳过当前测试的执行
func (c *T) Skipf(format string, args ...interface{})
...

```

接下来我们将利用testing库以及对应的T模块来完成测试代码的编写。

#### 14.2 编写测试

首先我们编写一个自己的utils库(utils.go)，这个库包含一些常用的函数，比如Sum函数可接收任意数量的整形数并返回所有的数的和。

```go
package utils

func Sum(arr ...int) int {
	var result int
	for _, value := range arr {
		result += value
	}
	return result
}
```
针对上面的代码我们编写的测试文件为utils_test.go,并将该文件与utils.go文件放在同一个文件夹（包）下。

```go
package utils

import "testing"

func TestSum(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expect := 21
	result := Sum(input...)
	if result != expect {
		t.Errorf("Expect %d ,but got %d", expect, result)
	}
}
```

我们可以通过go test来执行测试代码,可以尝试修改测试代码中的expect的值改为100，查看对应的输出是否不同，尝试加入-v参数查看详细的输出信息。

```
$ go test ./utils
ok      github.com/zhangmingkai4315/golang-essentials/14-测试/01-test-example/util
```

#### 14.3 表格测试

上面的测试中，我们在测试函数中仅仅写了一个测试例子是，这往往是不够的，有时候我们需要对于任何可能出错的情况的例子都要进行测试，涉及很多个类型的测试，因此我们可以考虑将所有的测试例子放在一起，编写代码，使得代码更紧凑，代码量更少，如下所示就是利用表格测试的方式来管理测试示例：
```go
func TestSum(t *testing.T) {
	type testdata struct {
		input  []int
		expect int
	}
	tests := []testdata{
		testdata{
			input:  []int{1, 2, 3, 4, 5, 6},
			expect: 21,
		},
		testdata{
			input:  []int{},
			expect: 0,
		},
		testdata{
			input:  []int{-1, -2, 2, 1},
			expect: 0,
		},
		testdata{
			input:  []int{1},
			expect: 1,
		},
	}
	for _, data := range tests {
		result := Sum(data.input...)
		if result != data.expect {
			t.Errorf("Expect %d ,but got %d", data.expect, result)
		}
	}
}
```
当我们执行go test的时候，将通过loop循环来执行所有的测试示例。

#### 14.4 编写示例测试

Example示例测试是go语言中测试的一种特殊形式，如果我们执行go test，所有的测试包括示例测试也会被同时执行，同时示例测试也被用于生成程序包的文档。以下是示例测试

```go 
// sum.go

// Sum will receive unlimit number of int
// and return sum of all numbers
func Sum(arr ...int) (result int) {
	for _, i := range arr {
		result += i
	}
	return
}

// sum_test.go

func ExampleSum() {
	fmt.Println(Sum(1, 2, 3, 4, 5))
	// Output:
	// 15
}

```

当我们对于该包执行测试的时候，测试程序会执行通过，加入我们修改Output中的值为16，则重新执行测试的时候会报错，说明go程序在检查测试的时候依赖于注释中的Output信息。

```sh
--- FAIL: ExampleSum (0.00s)
got:
15
want:
16
FAIL
FAIL    github.com/zhangmingkai4315/golang-essentials/14-测试/03-example        1.693s
```

同时为了建立示例程序与原始程序的关联，go语言有以下的规则：
```go
func Example() { ... } 提供包相关示例
func ExampleF() { ... } 提供函数相关示例
func ExampleT() { ... } 提供类型相关示例
func ExampleT_M() { ... }  提供类型函数相关示例
```
如果需要建立多个示例的话可以在上述规则的后面追加一些以小写字母开头的字符，比如

```go
func ExampleF_first() { ... }
func ExampleF_second() { ... }
func ExampleF_third() { ... }
```
#### 14.5 语法检查

针对go语言代码，有一些比较常见的go语言规范化工具，比如gofmt,go vet,golint,其中gofmt用于格式化go源代码，golint可以提供语法错误检查，并给出一些建议，go vet检查代码并报告有问题的代码结构，比如Printf调用函数中不包含任何format字符串，尽管这不会导致程序出错，但是仍旧是存在问题的一种使用方式。

安装golint的方式如下：
```
$ go get -u golang.org/x/lint/golint
$ golint -h
Usage of golint:
        golint [flags] # runs on package in current directory
        golint [flags] [packages]
        golint [flags] [directories] # where a '/...' suffix includes all sub-directories
        golint [flags] [files] # all must belong to a single package
```
如果我们对于03-example包执行golint则会提示如下的信息：
```
$ golint ./03-example/
03-example\sum.go:1:1: package comment should be of the form "Package example ..."

```
这是由于我们在写Package example时候的注释信息不太规范，应该尽量遵守golint提示的相关警告信息，使得代码尽量保持最佳规范实践。

gofmt是一个随go程序一起安装的二进制程序，用于格式化代码，比如使用tab来完成缩进，加入一些空格使得程序更加整洁等等如果使用vscode安装go语言插件，当我们保存的时候一般会自动的执行gofmt来格式化程序代码。

#### 14.6 压力测试

压力测试是测试功能中的一种，用于测试某一些函数或者操作执行的情况，通过同时运行大量的重复测试来计算出平均函数或者操作执行的时间，下面我们将写一个简单的封装函数来输出字符串的md5值，函数如下，主要是封装了crypto/md5库的函数来完成操作。

```go
import (
	"crypto/md5"
	"fmt"
)

func GetStringMd5(input string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}

```

针对压力测试的方式和普通的测试不太一样，压力测试的函数以Benchmark开头，同时传递的参数必须是**testing.B**的对象，压力函数执行过程中需要使用一个循环体结构，循环的次数由系统决定：

```go
func BenchmarkGetStringMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetStringMd5("hello world")
	}
}

```

函数编写完成后执行测试的方式也和上面的不同，如果需要执行压力测试，需要指定测试的参数**-bench**,执行过程如下：

```sh
go test ./04-benchmarking/ -bench GetStringMd5
goos: windows
goarch: amd64
pkg: github.com/zhangmingkai4315/golang-essentials/14-测试/04-benchmarking
BenchmarkGetStringMd5-12         2000000               618 ns/op
PASS
ok      github.com/zhangmingkai4315/golang-essentials/14-测试/04-benchmarking   6.449s
```