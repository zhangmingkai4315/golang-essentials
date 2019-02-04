### 8. 指针类型

#### 8.1  地址和指针

Go语言中的关键词“&”可以允许我们获取程序中变量的地址，取得地址的类型我们称之为指针类型，也就是保存的是实际的存储变量的位置，通过变量的位置我们找到之前定义的变量，并进行查询或修改等操作。下面是获取地址的简单示例：

```go
	var i = 10
	fmt.Printf("the address of i=%d is %v\n", i, &i)
	// the address of i=10 is 0xc000062058

	var j = &i
	fmt.Printf("the address of j=%d is %v\n", *j, j)
	// the address of j=10 is 0xc000062058
```

在上述例子中还存在一个关键词使用，那就是“*“这个关键词用于对于指针类型的取值，比如j保存的是i的地址，当我们使用 "\*j"操作的时候，实际上等效的是取得的i的值 因此通过上面的例子我么可以看到如何获取地址和获取地址中存储的值。

指针类型根据存储的变量类型不同而不同，也就是*int代表的整形指针和\*float代表的浮点型指针，尽管都是指针但是属于不同的类型，在作为参数传递的时候无法直接传递。

```go
	fmt.Printf("the type of i is %T, the type of j is %T\n", i, j)
	// the type of i is int, the type of j is *int
```

当我们获得一个变量的指针的时候，也就获得了所有针对该变量的操作权限，我们可以对于变量进行修改和查询比如下面的代码，尽管是针对变量j的操作，但是实际上却同时修改了i的值，也就是两者实际代表的是相同的内容。由于j为指针类型，因此在赋值的时候需要使用星号作取值操作。

```go
	fmt.Printf("the value of i is %d\n", *&i)
	// the value of i is 10

	*j = 20
	fmt.Printf("the value of i is %d\n", *&i)
	// the value of i is 20
```



#### 8.2. 指针类型的使用

对于指针类型的用法比较多的是在函数中进行参数的传递，**在go语言中所有的传递都是按值传递的**，因此如果常规类型作为参数传递，需要复制一份数据到函数中进行处理，为了减少数据的复制操作，同时满足对于原数据的修改，我们可以通过传递指针的值，这样就可以在函数中直接进行修改和查询操作。如下面的代码所示：

```go
func addValue(i int) {
	i = i + 1
}
func addValueByPointer(i *int) {
	*i = *i + 1
}
func main() {
	a := 10
	addValue(a)
	fmt.Printf("a = %d\n", a)
	// a = 10
	addValueByPointer(&a)
	fmt.Printf("a = %d\n", a)
	// a = 11
}
```



第一个函数接收的是整型类型，当我们传递a到函数addValue的时候，实际上程序会复制一份新的a对象，传递到函数中，这样尽管我们对于其进行了加一操作，但是修改的是复制的对象，原来的对象未发生变化。第二个函数则是通过接收整型指针的方式，获取到了原来存储a的地址，这样就可以找到a并对其进行修改，顺利的改变了a的内容。



#### 8.3 结构体函数指针

结构体中我们可以定义接收函数，既可以是结构体类型，也可以是结构体指针类型，如何选择可以根据实际需求，但是一般选择结构体指针的情况比较多，一方面可以减少数据复制的问题，另一方面也方便修改结构体数据。

```go
type city struct {
	name      string
	provience string
	country   string
}

func (c city) ShowCountry() string {
	return c.country
}

func (c *city) Info() string {
	return fmt.Sprintf("%s is in %s", c.name, c.ShowCountry())
}

func (c city) ChangeCountryName(country string) {
	c.country = country
}

func (c *city) ChangeCountryNameByPointer(country string) {
	c.country = country
}

```



针对类型T和其指针类型*T的接收函数在调用的时候需要确认以下：

- 对于指针类型接收函数（*T）可以接收T或者\*T的对象的调用
- 对于类型接收函数（T）如果T本身是可以取地址的，则可以使用T和其指针类型*T的变量调用
- 对于类型接收函数（T）如果T本身是不可以取地址(比如 nil)的，则只可以使用T的变量调用



尽管上述规则可能理解起来比较繁琐，但是大部分情况下即便是混合使用也不会出现问题，但是当需要对其进行修改（使用指针类型接收），或者与接口一起工作的时候需要注意。

特别是当结构体与接口结合使用的使用的时候要确认是结构体本身类型还是结构体指针类型满足接口的需求，比如下面的information接口，只有具有Info()函数的才能满足接口，也就是上面的结构体中，只有city指针类型满足要求，当我们以接口的形式接收参数或者执行其他的操作时，只能传递city指针类型，否则程序报错。

```go
type information interface {
	Info() string
}

func show(i information) {
	fmt.Println(i.Info())
}

// show(c)
//  cannot use c (type city) as type information in argument to show:
// city does not implement information (Info method has pointer receiver)

show(&c)
// beijing is in China
```