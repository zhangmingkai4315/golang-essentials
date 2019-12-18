### 6. Array、Slice和Map类型

#### 6.1. Array数组类型

golang中使用数组的机会比较少，一般都会使用后面介绍的Slice代替，Array具有以下的一些特点需要注意：

- 分配一个数组给另外一个将会拷贝所有的内容
- 传递数组给函数将会重新进行数组的复制
- 数组的大小也是类型的一部分，所以在传递参数的时候，要确保数组大小也一致，因此[10]int和[20]int不是一个类型。

```go
func showArray(arr [5]int) {
	fmt.Println(arr)
}

func main() {
	x := [10]int{}
	x[5] = 3
	fmt.Println(x)
	// [0 0 0 0 0 3 0 0 0 0]

	// showArray(x)
	//  cannotuse x (type [10]int) as type [5]int
}

```

数组在创建的时候，可以不用写具体的大小而使用...代替，同时我们可以指定赋值的位置，忽略其他的位置（自动使用初始化的值填充）比如下面的初始化方式：

```
y := [...]int{10, 20, 30, 40}
z := [5]int{1: 20, 2: 20}

fmt.Println(y, z)
// [10 20 30 40] [0 20 20 0 0]
```

针对数组的赋值操作，如果两个数组满足类型和长度一致，当直接进行赋值的时候，会进行内部的值的拷贝，新的数值和原来数组不属于同一个数组（很多语言程序赋值其实是引用传递）。比如下面的例子,我们对于上面的z数组进行复制操作，zCopy是一个具有相同长度和类型的数组，复制完成后修改其中一个，不会对另外的产生影响：

```go
zCopy := [5]int{}
zCopy = z
zCopy[0] = 1
fmt.Println(z, zCopy)
// [0 20 20 0 0] [1 20 20 0 0]
```

同样对于函数来说，如果传递的参数是数组的话，则整个数组会被完全的复制，按值传递，不管数组是一个空数组，还是一个具有百万级别元素的数组，完成复制操作有时候是需要很高的代价的，为了减少复制操作，可以使用引用的方式获取数组的地址比如&z,通过传递指针来完成数据的访问(小心数据被修改的问题)

#### 6.2 slice切片类型 

slice类型是在编写go程序中经常被运用到的类型，可以看作是长度动态变化的数组，底层尽管使用的是array，因此也是连续的内存块操作。slice本身存储占用24个字节(64位机器)，其中8字节的地址位置指向实际存储的数组位置，另外的两个8字节分别存储长度和容量值，因此操作切片更多的是利用引用的方式进行处理，减少了不必要的数据的复制：

```go
	x := []int{1, 2, 3, 4, 5}
	fmt.Printf("len(x)=%d\n", len(x))
	// len(x)=5 cap(x)=5
	
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println(x[:len(x)-2])
	// [2 3 4 5]
	// [2 3]
	// [1 2 3]
	subx := x[1:3]
	fmt.Printf("len(subx)=%d cap(subx)=%d\n", len(subx), cap(subx))
	// len(subx)=2 cap(subx)=4
```
截取片段的操作，会导致长度和容量的变化，比如上面subx的长度为3-1=2， subx的容量变为5-1=4， 5是原始数组的容量，1是起始位置。

上面代码中使用了切片操作，对于Slice(切片)本身重新进行截取操作，同时我们可以利用append实现数组的拼接和删除。

```go
	y := []int{100, 200, 300}

	x = append(x, y...)
	fmt.Println(x) //[1 2 3 4 5 100 200 300]

	x = append(x[:2], x[4:]...)
	fmt.Println(x) //[1 2 5 100 200 300]

```

使用go语言内置的make操作可以对于slice进行初始化，并设置长度和容量。使用这种方式可以优化slice的操作，默认情况下，初始化的slice长度和容量相同，但是一旦进行追加操作，容量无法满足的时候会创建一个新的slice并将原先的复制到新的容器中，如果持续的修改或者添加，将导致多次的内容复制。假如我们知道Slice的长度估计为1000左右我们可以直接使用make创建一个容量为1000的容器。

默认的容量扩展将按照两倍的方式进行处理，原来的复制到新的容器中。如果超过1000则按照25%的速度增加

```go
	z := make([]int, 10, 11)
	fmt.Printf("len of z is %d ,cap of z is %d\n", len(z), cap(z))
	// len of z is 10 ,cap of z is 11
	z = append(z, x...)
	fmt.Printf("len of z is %d ,cap of z is %d\n", len(z), cap(z))
	// len of z is 16 ,cap of z is 22
```
同时不管是数组还是Slice我们都可以使用多维的方式存储数据比如下面的代码，操作二维Slice如同一个excel表格一样。

```golang
persons := []string{"mike", "alice", "bob"}
city := []string{"beijing", "sanjun", "tokyo"}
info := [][]string{persons, city}

fmt.Println(info)
// [[mike alice bob] [beijing sanjun tokyo]]

```

数组的函数参数传递依赖于复制，是通过全部拷贝的方式，代价较大，Slice的传递则仅仅传递24个字节的数据（64位）



#### 6.3 map类型

map类型是go语言中内建的一种数据结构，内部使用键值的方式进行存储和查询。通过Hash表的方式实现数据查询和存储。hash函数生成的低位bit用于决定存储在哪个桶里面，而高位用于决定在桶里面的位置。同时为了减少空间占用，键和值都被写入桶中（所有键在前面，后面跟所有的值）
map键类型可以是整型，浮点数，字符类型，结构体以及数组（只要类型支持等式比较操作，切片操作不可以，由于不支持等式操作计算）。

```go
	person := map[string]uint16{
		"mike":  25,
		"alice": 20,
		"bob":   24,
	}
	fmt.Printf("mike age is %d\n", person["mike"])
	// mike age is 25
	if v, isExist := person["tom"]; isExist {
		fmt.Printf("tom age is %d\n", v)
	} else {
		fmt.Println("tom is not in list")
	}
	// tom is not in list
	
```

使用range操作可以迭代所有的map中的类型，只不过返回的是k,v键值对。使用delete可以删除map中的数据，但是不存在的数据当删除的时候不会有任何副作用。

```go
	for k, v := range person {
		fmt.Printf("%s age is %d\n", k, v)
	}
	// mike age is 25
	// alice age is 20
	// bob age is 24

	delete(person, "mike")
	fmt.Printf("%+v", person) //  map[alice:20 bob:24]
	delete(person, "nickey")
	fmt.Println(len(person)) // 2
```

#### 6.4 Append的实现

在slice中如果我们需要连接两个slice数据结构，我们可以借助于内部已经实现的append函数来完成操作，或者我们可以尝试自己取实现一个简单的针对int切片类型的Append函数，如下面的代码所示, 实现代码需要注意的地方在于，如果我们尝试直接copy数据到旧的切片中，可能导致容量不够，因此我们需要首先进行判断，是否容量超出，如果超出的话则尝试对其进行扩容，当然扩容的过程涉及到原有数据的复制操作。至于为什么多分配一倍的新切片对象空间，是为了将来进一步执行Append的时候不需要每次都重复进行数据复制的操作/

```golang

// Append will append the newSlice to the old slice and
// return the whole slice
func Append(oldSlice []int, newSlice []int) []int {
	if len(newSlice) == 0 {
		return oldSlice
	}
	length := len(oldSlice)
	if length+len(newSlice) > cap(oldSlice) {
		temSlice := make([]int, (length + cap(newSlice)*2))
		copy(temSlice, oldSlice)
		oldSlice = temSlice
	}
	oldSlice = oldSlice[0 : length+len(newSlice)]
	copy(oldSlice[length:], newSlice)
	return oldSlice
}
```

使用该函数如下面所示：
```go
oldSlice := []int{1, 2, 3, 4}
fmt.Printf("len(oldslice)=%d, cap(oldslice)=%d\n", len(oldSlice), cap(oldSlice))
//len(oldslice)=4, cap(oldslice)=4

newSlice := []int{5, 6, 7, 8, 9, 10}
oldSlice = Append(oldSlice, newSlice)
fmt.Println(oldSlice)
//[1 2 3 4 5 6 7 8 9 10]

fmt.Printf("len(oldslice)=%d, cap(oldslice)=%d\n", len(oldSlice), cap(oldSlice))
// len(oldslice)=10, cap(oldslice)=16
```

