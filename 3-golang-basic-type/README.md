### 3 go语言基本类型
#### 3.1 number数值类型

使用数值类型用于存储数字，比如整数或者小数，go语言中用于存储数字的类型比较多，比如uint8, uint16, uint32, int, float32 , float64以及一些复数类型。下面的代码中我们通过自动识别的方式存储数值到int和float中。

```go
	x := 10
	y := 10.2
	fmt.Printf("x = 10 is %T\n", x)
	fmt.Printf("y= 10.2 is %T\n", y)
	// x = 10 is int
	// y= 10.2 is float64
	
```

自动分配的int到底占有多少字节呢？ 这是根据不同的操作系统决定的，比如操作系统为32位的int和uint都只有32位也就是4个字节，操作系统为64位的则是64位=8个字节。

>  go语言中的byte类型其实是uint8类型; rune类型为int32类型,用于存储一个UTF8字符。

另外需要搞清楚二进制，八进制以及十六进制的区别以及如何转换。

#### 3.2 string字符类型

string类型代表了一系列的byte组成的字符串（可能为空），string类型为不可变类型，一旦创建无法被修改。但是可以修改变量引用的位置

```go
func main() {
	welcome := "hello world"
	welcomeByte := []byte(welcome)
	fmt.Println(welcomeByte)
	fmt.Printf("welcomeByte is %T and size is %d", welcomeByte, len(welcomeByte))
	// [104 101 108 108 111 32 119 111 114 108 100]
	// welcomeByte is []uint8

	welcomeCN := "hello 世界"
	welcomeByteCN := []byte(welcomeCN)
	fmt.Println(welcomeByteCN)
	fmt.Printf("welcomeByteCN is %T and size is %d", welcomeByteCN, len(welcomeByteCN))
	// welcomeByte is []uint8 and size is 11[104 101 108 108 111 32 228 184 150231 149 140]
	// welcomeByteCN is []uint8 and size is 12

	welcomeRuneCN := []rune(welcomeCN)
	fmt.Println(welcomeRuneCN)
	fmt.Printf("welcomeRuneCN is %T and size is %d", welcomeRuneCN, len(welcomeRuneCN))
	// [104 101 108 108 111 32 19990 30028]
	// welcomeRuneCN is []int32 and size is 8
}

```

关于string, bytes和rune的更多信息可以参考官方的blog:

https://blog.golang.org/strings



#### 3.3 const常量类型

常量代表了该对象一旦被初始化后，无法在进行重新赋值. 在定义常量的时候我们同样可以指定类型，

```go
	const name = "mike"
	const age = 26
	const (
        a1 = 1
        a2 = 2
        a3 = 3
	)
```



在go中我们定义常量有多种方式，同时利用iota关键词可以更容易的定义一组常量比如下面的代码

```go
const (
	b1 = iota        // 0
	b2               // 1
	b3 				 // 2
)

type Season uint8
const (
	Spring = Season(iota)          // Season(0)
	Summer                         // Season(1)
	Autumn						   // Season(2)
	Winner						   // Season(3)
)
const (
	_  = iota
	kb = 1 << (10 * iota)
	mb
	gb
	tb
)
```