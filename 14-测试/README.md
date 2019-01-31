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





