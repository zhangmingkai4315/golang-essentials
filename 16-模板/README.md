### 16 模板

标准包```template```提供了基于数据驱动的模板编程的方式， 通过模板可以实现数据的动态渲染和批量生成。Go实现了两个标准包分别是**html/template**和**text/template**，两个包提供的接口基本相同，功能也基本相同，不同之处在于**html/template**可以帮助处理一些具有html标签的转义处理，防止恶意的代码注入操作。我们先通过**text/template**来介绍一些常用的接口函数。

#### 16.1 模板渲染

下面我们通过一个简单的实例来介绍模板包的使用，在这个例子中我们的数据集合是一些学生student的切片类型，保存了很多person结构体数据。
templateStr是一个模板字符串，通过一些预定义的占位符来代表将来数据传递过来后需要接收的位置，比如其中的{{.Name}}代表了传递数据中Name字段，{{ .Age}}代表传递的数据中Age字段，如果{{ . }}则代表数据本身。```template.New("test").Parse(templateStr)```创建一个模板对象，并给予该对象test的名称（后面部分例子中会用到该名称），最后通过template.Execute来执行数据的传递，将每一个person对象传递到模板中生成最后的数据。

```go
import (
	"fmt"
	"os"
	"text/template"
)
type person struct {
	Name string
	Age  int
}
var templateStr = `
Hi, {{.Name}}. I know you are {{ .Age}} years old.
`

func main() {
	students := []person{
		{"Tom", 26},
		{"Bob", 24},
		{"Alice", 25},
		{"Mike", 29},
	}
	tmpl, err := template.New("test").Parse(templateStr)
	if err != nil {
		panic(err)
	}
	for _, s := range students {
		err = tmpl.Execute(os.Stdout, s)
		if err != nil {
			panic(err)
		}
	}

```
模板执行的结果我们直接写入标准输出，我们也可以传递到任何io.Writer的对象中。输出的结果如下面所示：


```go
Hi, Tom. I know you are 26 years old.
Hi, Bob. I know you are 24 years old.
Hi, Alice. I know you are 25 years old.
Hi, Mike. I know you are 29 years old.
```


#### 16.2 切片或数组类型渲染

template包本身还支持很多语法操作来提供灵活的模板调用，比如当传递的数据不是结构体而是切片或者数组结构，又或者是map结构的时候都可以借助于template预设的关键词动态的渲染该数据类型. 下面的实例中我们依然传递的是students结构体数组，不同之处在于此处我们没有使用for循环去取值后传递到模板，而是直接将整个数组对象传递进去。 

模板```templateRangeStr```中使用了range来迭代{{ . }}, 上面我们讲到了{{ . }}在模板中代表传递的值本身，并使用{{ end }}来结束循环，中间的可以看做是每次的循环体，其中循环体中的{{ . }}则代表每一个迭代出来的值，也就是每一个person结构体


```go
var templateRangeStr = `
{{ range . }}
Hi, {{.Name}}. I know you are {{ .Age}} years old. (range version)
{{ end }}
`

rangeTmpl, err := template.New("rangetest").Parse(templateRangeStr)
if err != nil {
    panic(err)
}
err = rangeTmpl.Execute(os.Stdout, students)
if err != nil {
    panic(err)
}
```

重新执行实例后的结果如下面所示：

```go
Hi, Tom. I know you are 26 years old. (range version)
Hi, Bob. I know you are 24 years old. (range version)
Hi, Alice. I know you are 25 years old. (range version)
Hi, Mike. I know you are 29 years old. (range version)
```


#### 16.3 Map类型渲染

Map类型与数组类型不同在于，Map类型是按照键值的方式存储的，当我们循环迭代的时候需要同时将其中的键key和值value提取出来，如下面实例所示：

```go
var templateKVStr = `
{{ range $key, $value := . }}
Hi ID={{ $key }}, {{ $value.Name }}. I know you are {{ $value.Age }} years old. (map version)
{{ end }}
`
studentsWithID := map[int]person{
    1: {"Tom", 26},
    2: {"Bob", 24},
    3: {"Alice", 25},
    4: {"Mike", 29},
}
kvtmpl, err := template.New("kvtest").Parse(templateKVStr)
if err != nil {
    panic(err)
}
err = kvtmpl.Execute(os.Stdout, studentsWithID)
if err != nil {
    panic(err)
}
```
上面的例子中我们使用range迭代一个Map类型，获得的值分别的存储在$key和$value中，这里带有$的前缀代表了模板中创建的变量内容，可以之后的模板中直接使用，比如{{ $value.Name }} 对变量的Name属性取值。上面的例子执行结果如下：

```go
Hi ID=1, Tom. I know you are 26 years old. (map version)
Hi ID=2, Bob. I know you are 24 years old. (map version)
Hi ID=3, Alice. I know you are 25 years old. (map version)
Hi ID=4, Mike. I know you are 29 years old. (map version)
```

#### 16.4 结构体函数调用

在模板中我们可以直接调用函数或结构体的方法，比如下面我们给person结构体增加一个方法Info用于打印相关信息，为了在模板中使用该方法，可以直接使用类似于取字段值的方式来调用，如下面的实例所示：

```go
type person struct {
	Name string
	Age  int
}

func (p person) Info() string {
	return fmt.Sprintf("Name=%s,Age=%d", p.Name, p.Age)
}

var templateKVWithMethodStr = `
{{ range $key, $value := . }}
Hi ID={{ $key }}, {{ $value.Info }}
{{ end }}
`

kvMethodTmpl, err := template.New("kvMethodTest").Parse(templateKVWithMethodStr)
if err != nil {
    panic(err)
}
err = kvMethodTmpl.Execute(os.Stdout, studentsWithID)
if err != nil {
    panic(err)
}

```

执行上面的代码结果如下面所示, 需要注意的是假如我们讲上面的函数改为指针接收器方式,函数签名改为```func (p *person) Info() string ```，则重新执行无法正常通过，由于传递的是变量，而调用的为指针类型的方法集。而正常函数调用则会由编译器自动转换实现自适应，由此可以看到模板中的调用方式和正常调用在代码处理上还是有一定区别的。

```go
Hi ID=1, Name=Tom,Age=26
Hi ID=2, Name=Bob,Age=24
Hi ID=3, Name=Alice,Age=25
Hi ID=4, Name=Mike,Age=29

```


#### 16.4 内置函数调用

在template包中已经提供了很多内置的函数可以直接使用，这里我们简单的介绍一些常见的使用实例，比如index,and,以及or的使用，其他的可以参考[官方文档](https://golang.org/pkg/text/template/#hdr-Functions)在此不一一给出使用实例。


首先我们定义一个结构体对象如下：

```go
type user struct {
	Name     string
	Phones   []string
	IsAdmin  bool
	IsLoggin bool
}
mike := user{"Mike", []string{"1234567890", "010-123456"}, true, true}

```
后面我们会结合这个结构体分别给出相关的实例。

##### index使用实例

在模板中使用index用于获取列表，数组中的某一条记录数据，因此需要给出具体的位置信息，比如下面的{{index .Phones 0 }}等价于Phones[0]的变量取值操作。

```go
indexTplStr := `
the first phone number is {{index .Phones 0 }}
`
indexFuncTpl, err := template.New("functest").Parse(indexTplStr)
if err != nil {
    panic(err)
}
err = indexFuncTpl.Execute(os.Stdout, mike)
if err != nil {
    panic(err)
}
// the first phone number is 1234567890
```
##### and使用实例

在模板中使用and用于对多个表达式执行与操作，比如下面的{{ if and .IsLoggin .IsAdmin }}等价于user.IsLoggin && user.IsAdmin, 并通过if else操作获得按条件执行。

```go
andTplStr := `
the user is{{ if and .IsLoggin .IsAdmin }} admin {{ else if .IsLoggin }} login user {{ else }} guest{{end}}
`
andFuncTpl, err := template.New("andtest").Parse(andTplStr)
if err != nil {
    panic(err)
}
err = andFuncTpl.Execute(os.Stdout, mike)
if err != nil {
    panic(err)
}
// the user is admin
```

##### or使用实例

在模板中使用or用于对多个表达式执行或操作，比如下面的{{ if or .IsLoggin .IsAdmin }} 等价于user.IsLoggin || user.IsAdmin

```go
orTplStr := `
the user is{{ if or .IsLoggin .IsAdmin }} login {{ else }} guest{{end}}
`
orFuncTpl, err := template.New("ortest").Parse(orTplStr)
if err != nil {
    panic(err)
}
err = orFuncTpl.Execute(os.Stdout, mike)
if err != nil {
    panic(err)
}
// the user is login
```

#### 16.5 自定义函数调用

除了上述介绍标准包直接包含的函数外，template包还支持自定义的函数定义，通过创建自定义的函数，并将函数传递到模板中调用，实现灵活的数据处理，减少复杂模板的编写。下面我们根据上面的实例来添加一个自定义的函数，这里需要借助于内置的template.Funcs以及template.FuncMap两个接口实现. 在模板语句中我们直接调用了一个自定义的函数hasPermission，后面会给出详细的创建函数的方式。

```go
customFuncTplStr := `
{{ if hasPermission . "admin" }} the user is admin {{ end }}
`
permissionFunc := func(user user, level string) bool {
    if level == "admin" {
        return user.IsAdmin
    }
    return false

}
customFuncTpl, err := template.New("customFuncTest").Funcs(template.FuncMap{
    "hasPermission": permissionFunc,
}).Parse(customFuncTplStr)
if err != nil {
    panic(err)
}
err = customFuncTpl.Execute(os.Stdout, mike)
if err != nil {
    panic(err)
}
// the user is admin
```

permissionFunc定义就是一个普通的函数定义，需要注意的是参数的设置，第一个参数应该是模板中传递的变量的值，后面是传递给函数的其他参数，使用template.FuncMap{}可以创建一系列的函数集合Map类型，键是函数名称，值为所对应的函数体，通过template.Funcs绑定到模板中，供模板中调用


