### 6. Array、Slice和Map类型

#### 6.1. Array数组类型

golang中使用数组的机会比较少，一般都会使用后面介绍的Slice代替，Array具有以下的一些特点需要注意：

- 分配一个数组给另外一个将会拷贝所有的内容
- 传递数组给函数将会重新进行数组的复制
- 数组的大小也是类型的一部分，所以在传递参数的时候，要确保数组大小也一致。

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

数组在创建的时候，可以不用写具体的大小而使用...代替，比如下面的初始化：

```
array := [...]float64{7.0, 8.5, 9.1}
```

#### 6.2 slice切片类型 

slice类型是在编写go程序中经常被运用到的类型，可以看作是长度动态变化的数组，底层尽管使用的是array,但是操作是利用引用的方式，提供了很多灵活的操作方式比如下面的：

```go
	x := []int{1, 2, 3, 4, 5}
	fmt.Printf("len(x)=%d\n", len(x))
	// len(x)=5
	
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println(x[:len(x)-2])
	// [2 3 4 5]
	// [2 3]
	// [1 2 3]
```

上面代码中使用了切片操作，对于Slice(切片)本身重新进行截取操作，同时我们可以利用append实现数组的拼接和删除。

```go
	y := []int{100, 200, 300}

	x = append(x, y...)
	fmt.Println(x) //[1 2 3 4 5 100 200 300]

	x = append(x[:2], x[4:]...)
	fmt.Println(x) //[1 2 5 100 200 300]

```

使用go语言内置的make操作可以对于slice进行初始化，并设置长度和容量。使用这种方式可以优化slice的操作，默认情况下，初始化的slice长度和容量相同，但是一旦进行追加操作，容量无法满足的时候会创建一个新的slice并将原先的复制到新的容器中，如果持续的修改或者添加，将导致多次的内容复制。加入我们知道Slice的长度估计为1000左右我们可以直接使用make创建一个容量为1000的容器。

默认的容量扩展将按照两倍的方式进行处理，原来的复制到新的容器中。

```
	z := make([]int, 10, 11)
	fmt.Printf("len of z is %d ,cap of z is %d\n", len(z), cap(z))
	// len of z is 10 ,cap of z is 11
	z = append(z, x...)
	fmt.Printf("len of z is %d ,cap of z is %d\n", len(z), cap(z))
	// len of z is 16 ,cap of z is 22

```



同时不管是数组还是Slice我们都可以使用多维的方式存储数据比如下面的代码，操作二维Slice如同一个excel表格一样。

```
	persons := []string{"mike", "alice", "bob"}
	city := []string{"beijing", "sanjun", "tokyo"}
	info := [][]string{persons, city}

	fmt.Println(info)
	// [[mike alice bob] [beijing sanjun tokyo]]

```



#### 6.3 map类型

map类型是go语言中内建的一种数据结构，内部使用键值的方式进行存储和查询。	map的了类型可以是整型，浮点数，字符类型，结构体以及数组（只要类型支持等式比较操作，切片操作不可以，由于不支持等式操作计算）。

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

### 6.4 练习

暂无

