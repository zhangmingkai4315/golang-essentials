### 3 程序控制流

控制流用来控制程序的执行，在程序中默认情况下将顺序执行一条条的语句，通过一些程序本身提供的控制流管理，我们可以实现代码执行的规划，比如循环执行，中断执行，跳转执行等等，使得程序运行更加灵活。详细的介绍可以参考Wiki的介绍：[什么是控制流?](https://zh.wikipedia.org/wiki/%E6%8E%A7%E5%88%B6%E6%B5%81%E7%A8%8B)


#### 3.1 for循环语句

在Go语言中循环执行的关键词是使用for语句,不同与其他的编程的语言，Go语言中的for语句通过不同的使用方式承担了其他编程语言比如Java, JavaScript,Python等需要多个关键词才能实现的功能。

首先一个基本的for循环有三个组成部分，且每一个部分都通过分号进行分割：

- init初始化语句，只在第一次循环前执行
- condition条件表达式，在每一次循环前计算条件是否满足
- post循环后处理表达式，在每一次循环后执行

通过三个组成部分，进行不断的循环执行，直到第二个条件计算不再满足或者返回false，才终止程序执行。比如下面的代码所示：
```go

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

```
> 不同与其他的编程语言，这里的for循环三个条件并不需要使用括号括起来

for循环的使用中如果我们省略掉开始的init和后面的post执行表达式，则程序就和其他编程语言中的while关键词一样，不断的循环，直到满足条件为止, 下面的两个函数执行效果是一样的，当我们省略掉init和post，其链接表达式的分号也可以直接省略，程序变得更加简练。

```golang

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

```

假如我们省略三个表达式，则只剩下for关键词，则程序执行进入一个无限循环的状态，因为不再需要任何的条件进行判断，这时候如果内部的函数体不主动退出执行的话， 则程序将不会被终止,比如下面的语句,将不断的循环，直到其中的v值降低到0 后调用break退出程序：

```golang
    v := 100
	for {
        if v ==0 {
            break
        }
        fmt.Println("loop...")
        v--
	}
```

我们可以嵌套多个for语句来实现多重循环，如下面的示例所示, 假如外层循环次数为n，内层循环数为m，则程序多重循环执行的次数为n*m:

```go
for i := 0; i < 10; i++ {
    for j := 0; j < 10; j++ {
        fmt.Printf("i=%d,j=%d\n", i, j)
    }
}
```

利用for循环我们还可以迭代出数组，Slice以及map对象中的内容，对于数组和Slice类型，每次迭代的都是[index,value]的数值对，其中index为存储的数值位置，value为存储的值，而map对象则每次迭代的是[key,value]的数值对，其中key为存储的键，value为对应的值。如果我们循环中不需要其中的每一个数值，比如对于数组循环不需要index序列号，可以借助于下面的方式实现，对于map类型同样适用。另外for循环还可以从channel中获取数值，这个会在后面的channel章节来进行详细的介绍。

```golang
type user struct {
	name  string
	email string
}
users := []user{
    {"Mike", "mike@example.com"},
    {"Alice", "alice@example.com"},
}
for _, user := range users {
    fmt.Printf("name = %s ; email = %s \n", user.name, user.email)
}
// name = Mike ; email = mike@example.com
// name = Alice ; email = alice@example.com

persons := map[int]user{
    100: user{name: "mike", email: "mike@example.com"},
    101: user{name: "alice", email: "alice@example.com"},
}

for k, v := range persons {
    fmt.Printf("No%d :name = %s ; email = %s \n", k, v.name, v.email)
}
// No100 :name = mike ; email = mike@example.com
// No101 :name = alice ; email = alice@example.com

```

#### 3.2 break和continue

如果熟悉其他编程语言的话比如java或者python，对于break和continue两个关键词一定不会感觉陌生，同样在go语言中这两个关键词的作用与其他语言一致，分别用于退出当前的循环和继续当前的循环。比如下面的例子中，我们创建一个for无限循环，当i的数值为偶数的时候打印，否则继续下一轮循环，直到i的数值大于100退出程序的执行。

```golang
func main(){
	i := 0
	for {
		i++
		if i>100{
			break
		}
		if i %2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
```

注意continue只能用于loop循环中，而break只能用于loop, switch或者select语句中(后面的两个类型我们将稍后介绍), 否则程序将报错退出。

#### 3.3 if判断语句

在Go语言中if判断语句的形式和for的一些使用方式比较相似，不同的是if语句不会循环执行后面的函数体，如果条件满足，则仅仅会执行一次，同时我们在if语句中可以加入初始化过程（可以是赋值也可以是函数体计算）来完成变量的初始化，下面的示例介绍了if的一些典型使用方式。

```golang
	if true {
		fmt.Println("Always print")
    }
    // Always print

	if false {
		fmt.Println("Never print")
    }
    
	if num := 9; num < 0 {
		fmt.Println("Number is negative")
	} else if num > 0 {
		fmt.Println("Number is positive")
	} else {
		fmt.Println("Number is zero")
	}
    // Number is positive
    
    if err := foo(); err != nil {
        panic(err)
    }
```


#### 3.4 Switch选择语句

上面的if else 语句，可以实现简单的条件判断，但是如果我们有多个条件的时候，不断的增加else的数量会导致程序既不美观，也难以阅读，因此不同的编程语言都有类似于Switch的关键次来优化程序的设计，实现上面的目标。Go语言中的switch也不例外，基本的使用方式如下所示, switch语句后面跟条件判断语句，通过计算该值来与后面的case语句进行**顺序匹配**，任何匹配成功后会执行case后的函数体，并退出（除非使用fallthrough强制继续执行）。

```golang

func main() {
	now := time.Now().Unix()
	mins := now % 2
	switch mins {
	case 0:
		fmt.Println("even")
	case 1:
		fmt.Println("odd")
	}
}

```

Go语言中switch还有以下的使用要点需要注意：
- 数值匹配的时候，类型一定要一致
- 可以使用default来设置一个缺省执行的操作,用来满足没有任何匹配成功的情况
- 可以在case中使用表达式或者函数来计算一个匹配值（前面已经满足的时候，该表达式或函数不会执行）
- case中可以使用多个值，来实现多值匹配
- switch关键词可以不包含任何的表达式，此时将选择第一个case中表达式为true的进行执行
- 多个case表达式可以相同，使用fallthrough关键词来允许多个匹配执行

```go
package main

import "fmt"

func main() {
	score := 7
	switch score {
	case 0, 1, 3:
		fmt.Println("Terrible")
	case 4, 5:
		fmt.Println("Mediocre")
	case 6, 7:
		fmt.Println("Not bad")
	case 8, 9:
		fmt.Println("Almost perfect")
	case 10:
		fmt.Println("hmm did you cheat?")
	default:
		fmt.Println(score, " off the chart")
	}
}
```

以下的例子将使用无表达式的switch语句，这时候程序将选择第一个满足true条件的进行执行， 另外多个case的值也可以相同，使用fallthrough来同时执行多个语句：

```go
switch {
case (2 == 2):
    fmt.Println("2==2 is true")
    fallthrough
case (3 == 3):
    fmt.Println("3==3 is true too")
default:
    fmt.Println("default case")
}
// 2==2 is true
// 3==3 is true too
```
同时使用switch还可以借助于类型断言表达式来进行类型的判断和选择，比如下面的示例,函数```typePrintFunc ```中的参数i可以是任何的类型（interface{}满足任何类型），使用```i.(type)```我们可以获得该类型的信息，通过switch来进行判断执行来执行对应的语句。

```golang

typePrintFunc := func(i interface{}) {
    switch i.(type) {
    case bool:
        fmt.Println("args is a bool type")
    case string:
        fmt.Println("args is a string")
    default:
        fmt.Println("args type unknown")
    }
}
typePrintFunc("mike")
typePrintFunc(12)
// args is a string
// args type unknown
    
```
