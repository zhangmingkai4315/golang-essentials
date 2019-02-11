# Golang语言编程指南

![](https://github.com/zhangmingkai4315/golang-essentials/blob/master/public/golang.png?raw=true)

2007年, 当时Google在自己的生产环境中主要使用的语言以C++, Java和Python为主，但是发现随着系统规模越来越大，软硬件的不断升级，这些传统的编程语言本身的设计和使用或多或少的对于运行环境产生了一些负面的影响，比如编译构建时间过长，包依赖管理复杂，并发处理不够灵活，服务更新和升级繁琐等问题。因此Google的技术人员计划从语言层面去尝试解决这些问题，通过开发一种新的编程语言来彻底的解决这些问题，这就导致了最终Go语言的诞生。

Go语言的开发团队成员也都是一些知名的计算机行业专家其中包括：Rob Pike ，曾参与Unix 以及UTF-8的开发，Robert Griesemer 早期Pascal语言的开发人员，Ken Thompson曾参与Unix，B语言及C语言设计工作。这些技术专家的加入以及Google公司的背景，使得对于Go语言自诞生之初，就获得了众多的关注，经过十几年的发展，Go语言也其稳定可靠，干净效率的编程发生赢多了很多工程师的喜爱。

![](https://github.com/zhangmingkai4315/golang-essentials/blob/master/public/stackoverflow-golang-wanted.PNG?raw=true)
在过去十年间，Go语言获得了越来越多的人的关注。据StackOverflow的[2018年技术调查报告](https://insights.stackoverflow.com/survey/2018/)显示，2018年Go语言与Python和JavaScript编程语言一起成为众多技术人员最想学习的编程语言。作为一名DevOps工程师，当前线上部署的大规模运维支撑系统越来越多的被Go语言编写的项目所占据，比如Docker, Kubernetes, Prometheus, Grafana，CoreDNS等等，这些项目已经深入到我们的日常的运维开发工作之中。


>  I think Node is not the best system to build a massive server web. I would use Go for that. And honestly, that’s the reason why I left Node. It was the realization that: oh, actually, this is not the best server-side system ever.                
                                    **- Ryan Dahl（NodeJS作者）**

如果使用一句话来简单的介绍Go语言那就是: **Go语言是一个编译型，并发，自动垃圾回收的静态语言, 由Google最初开发，但已开源并获的大规模的使用**。使用Go语言可以进行系统及网络方面的程序设计开发，比如开发分布式系统，命令行程序，web后台系统，以及一些图像处理，加密算法等需求。

本书通过一步步的介绍Go语言的语法规范，语言特性来带领大家熟悉和理解Go语言，并能够运用到日常的工作之中，同时后期也会不断的维护和更新，保证代码和内容与当前最新版本一致。

## 目录

### [1.开发环境准备](https://github.com/zhangmingkai4315/golang-essentials/tree/master/01-%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83%E5%AE%89%E8%A3%85)
- [1.1 Windows环境安装](https://github.com/zhangmingkai4315/golang-essentials/tree/master/01-%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83%E5%AE%89%E8%A3%85#11-windows%E7%8E%AF%E5%A2%83)
- [1.2 Linux环境安装](https://github.com/zhangmingkai4315/golang-essentials/tree/master/01-%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83%E5%AE%89%E8%A3%85#12-linux%E7%8E%AF%E5%A2%83)
- [1.3 MacOS环境安装](https://github.com/zhangmingkai4315/golang-essentials/tree/master/01-%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83%E5%AE%89%E8%A3%85#12-linux%E7%8E%AF%E5%A2%83)

### [2.变量和类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/02-%E5%8F%98%E9%87%8F%E5%92%8C%E7%B1%BB%E5%9E%8B)
- [2.1 编写第一个Go语言程序](https://github.com/zhangmingkai4315/golang-essentials/tree/master/02-%E5%8F%98%E9%87%8F%E5%92%8C%E7%B1%BB%E5%9E%8B#21-%E7%BC%96%E5%86%99%E7%AC%AC%E4%B8%80%E4%B8%AAgo%E8%AF%AD%E8%A8%80%E7%A8%8B%E5%BA%8F)
- [2.2 使用标准包](https://github.com/zhangmingkai4315/golang-essentials/tree/master/02-%E5%8F%98%E9%87%8F%E5%92%8C%E7%B1%BB%E5%9E%8B#22-%E4%BD%BF%E7%94%A8%E6%A0%87%E5%87%86%E5%8C%85)
- [2.3 参数和类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/02-%E5%8F%98%E9%87%8F%E5%92%8C%E7%B1%BB%E5%9E%8B#23-%E5%8F%82%E6%95%B0%E5%92%8C%E7%B1%BB%E5%9E%8B)
- [2.4 变量定义](https://github.com/zhangmingkai4315/golang-essentials/tree/master/02-%E5%8F%98%E9%87%8F%E5%92%8C%E7%B1%BB%E5%9E%8B#24-go%E8%AF%AD%E8%A8%80%E4%B8%AD%E7%9A%84%E5%8F%98%E9%87%8F%E5%AE%9A%E4%B9%89)
- [2.5 语言类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/02-%E5%8F%98%E9%87%8F%E5%92%8C%E7%B1%BB%E5%9E%8B#25-%E7%B1%BB%E5%9E%8B)
### [3.控制流](https://github.com/zhangmingkai4315/golang-essentials/tree/master/03-%E6%8E%A7%E5%88%B6%E6%B5%81)

- [3.1 for循环语句](https://github.com/zhangmingkai4315/golang-essentials/tree/master/03-%E6%8E%A7%E5%88%B6%E6%B5%81#31-for%E5%BE%AA%E7%8E%AF%E8%AF%AD%E5%8F%A5)
- [3.2 break和continue](https://github.com/zhangmingkai4315/golang-essentials/tree/master/03-%E6%8E%A7%E5%88%B6%E6%B5%81#32-break%E5%92%8Ccontinue)
- [3.3 if判断语句](https://github.com/zhangmingkai4315/golang-essentials/tree/master/03-%E6%8E%A7%E5%88%B6%E6%B5%81#32-break%E5%92%8Ccontinue)
- [3.4 Switch选择语句](https://github.com/zhangmingkai4315/golang-essentials/tree/master/03-%E6%8E%A7%E5%88%B6%E6%B5%81#34-switch%E9%80%89%E6%8B%A9%E8%AF%AD%E5%8F%A5)

### [4.基本类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/04-%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B)

- [4.1 number数值类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/04-%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B#41-number%E6%95%B0%E5%80%BC%E7%B1%BB%E5%9E%8B)
- [4.2 string字符类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/04-%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B#42-string%E5%AD%97%E7%AC%A6%E7%B1%BB%E5%9E%8B)
- [4.3 const常量类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/04-%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B#43-const%E5%B8%B8%E9%87%8F%E7%B1%BB%E5%9E%8B)
- [4.4 类型转换](https://github.com/zhangmingkai4315/golang-essentials/tree/master/04-%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B#44-%E7%B1%BB%E5%9E%8B%E8%BD%AC%E6%8D%A2)

### [5.interface类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/05-interface%E7%B1%BB%E5%9E%8B)

- [5.1 接口Interface](https://github.com/zhangmingkai4315/golang-essentials/tree/master/05-interface%E7%B1%BB%E5%9E%8B#51-%E6%8E%A5%E5%8F%A3interface)
- [5.2 类型断言](https://github.com/zhangmingkai4315/golang-essentials/tree/master/05-interface%E7%B1%BB%E5%9E%8B#52-%E7%B1%BB%E5%9E%8B%E6%96%AD%E8%A8%80)
- [5.3 实现sort接口](https://github.com/zhangmingkai4315/golang-essentials/tree/master/05-interface%E7%B1%BB%E5%9E%8B#53-%E5%AE%9E%E7%8E%B0sort%E6%8E%A5%E5%8F%A3)
- [5.4 HTTP接口]()

### [6.数组切片和Map类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/06-%E6%95%B0%E7%BB%84%E5%88%87%E7%89%87%E5%92%8CMap%E7%B1%BB%E5%9E%8B)
- [6.1. Array数组类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/06-%E6%95%B0%E7%BB%84%E5%88%87%E7%89%87%E5%92%8CMap%E7%B1%BB%E5%9E%8B#61-array%E6%95%B0%E7%BB%84%E7%B1%BB%E5%9E%8B)
- [6.2 slice切片类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/06-%E6%95%B0%E7%BB%84%E5%88%87%E7%89%87%E5%92%8CMap%E7%B1%BB%E5%9E%8B#62-slice%E5%88%87%E7%89%87%E7%B1%BB%E5%9E%8B)
- [6.3 map类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/06-%E6%95%B0%E7%BB%84%E5%88%87%E7%89%87%E5%92%8CMap%E7%B1%BB%E5%9E%8B#63-map%E7%B1%BB%E5%9E%8B)
### [7.结构体](https://github.com/zhangmingkai4315/golang-essentials/tree/master/07-%E7%BB%93%E6%9E%84%E4%BD%93)
- [7.1 嵌套结构体](https://github.com/zhangmingkai4315/golang-essentials/tree/master/07-%E7%BB%93%E6%9E%84%E4%BD%93#71-%E5%B5%8C%E5%A5%97%E7%BB%93%E6%9E%84%E4%BD%93)
- [7.2 匿名结构体](https://github.com/zhangmingkai4315/golang-essentials/tree/master/07-%E7%BB%93%E6%9E%84%E4%BD%93#72-%E5%8C%BF%E5%90%8D%E7%BB%93%E6%9E%84%E4%BD%93)
- [7.3 使用new关键词](https://github.com/zhangmingkai4315/golang-essentials/tree/master/07-%E7%BB%93%E6%9E%84%E4%BD%93#73-%E4%BD%BF%E7%94%A8new%E5%85%B3%E9%94%AE%E8%AF%8D)

### [8.函数](https://github.com/zhangmingkai4315/golang-essentials/tree/master/08-%E5%87%BD%E6%95%B0)
- [8.1 函数的创建](https://github.com/zhangmingkai4315/golang-essentials/tree/master/08-%E5%87%BD%E6%95%B0#81-%E5%87%BD%E6%95%B0%E7%9A%84%E5%88%9B%E5%BB%BA)
- [8.2 可变参数](https://github.com/zhangmingkai4315/golang-essentials/tree/master/08-%E5%87%BD%E6%95%B0#82-%E5%8F%AF%E5%8F%98%E5%8F%82%E6%95%B0)
- [8.3 defer函数](https://github.com/zhangmingkai4315/golang-essentials/tree/master/08-%E5%87%BD%E6%95%B0#83-defer%E5%87%BD%E6%95%B0)
- [8.4 匿名函数](https://github.com/zhangmingkai4315/golang-essentials/tree/master/08-%E5%87%BD%E6%95%B0#84-%E5%8C%BF%E5%90%8D%E5%87%BD%E6%95%B0)
- [8.5 传递和返回函数](https://github.com/zhangmingkai4315/golang-essentials/tree/master/08-%E5%87%BD%E6%95%B0#85-%E4%BC%A0%E9%80%92%E5%92%8C%E8%BF%94%E5%9B%9E%E5%87%BD%E6%95%B0)
- [8.6 回调函数](https://github.com/zhangmingkai4315/golang-essentials/tree/master/08-%E5%87%BD%E6%95%B0#86-%E5%9B%9E%E8%B0%83%E5%87%BD%E6%95%B0)
- [8.7 闭包](https://github.com/zhangmingkai4315/golang-essentials/tree/master/08-%E5%87%BD%E6%95%B0#82-%E9%97%AD%E5%8C%85)
- [8.8 递归函数](https://github.com/zhangmingkai4315/golang-essentials/tree/master/08-%E5%87%BD%E6%95%B0#88-%E9%80%92%E5%BD%92%E5%87%BD%E6%95%B0)
- [8.9 init函数](https://github.com/zhangmingkai4315/golang-essentials/tree/master/08-%E5%87%BD%E6%95%B0#89-init%E5%87%BD%E6%95%B0)
### [9.指针](https://github.com/zhangmingkai4315/golang-essentials/tree/master/09-%E6%8C%87%E9%92%88)
- [9.1 地址和指针](https://github.com/zhangmingkai4315/golang-essentials/tree/master/09-%E6%8C%87%E9%92%88#91--%E5%9C%B0%E5%9D%80%E5%92%8C%E6%8C%87%E9%92%88)
- [9.2. 指针类型的使用](https://github.com/zhangmingkai4315/golang-essentials/tree/master/09-%E6%8C%87%E9%92%88#92-%E6%8C%87%E9%92%88%E7%B1%BB%E5%9E%8B%E7%9A%84%E4%BD%BF%E7%94%A8)
- [9.3 结构体函数指针](https://github.com/zhangmingkai4315/golang-essentials/tree/master/09-%E6%8C%87%E9%92%88#93-%E7%BB%93%E6%9E%84%E4%BD%93%E5%87%BD%E6%95%B0%E6%8C%87%E9%92%88)
### [10.编写应用程序](https://github.com/zhangmingkai4315/golang-essentials/tree/master/10-%E7%BC%96%E5%86%99%E5%BA%94%E7%94%A8)
- [10.1 使用Go语言编写JSON序列化应用](https://github.com/zhangmingkai4315/golang-essentials/tree/master/10-%E7%BC%96%E5%86%99%E5%BA%94%E7%94%A8#101-%E4%BD%BF%E7%94%A8go%E8%AF%AD%E8%A8%80%E7%BC%96%E5%86%99json%E5%BA%8F%E5%88%97%E5%8C%96%E5%BA%94%E7%94%A8)
- [10.2 官方源码程序](https://github.com/zhangmingkai4315/golang-essentials/tree/master/10-%E7%BC%96%E5%86%99%E5%BA%94%E7%94%A8#102-%E5%AE%98%E6%96%B9%E6%BA%90%E7%A0%81%E7%A8%8B%E5%BA%8F)
- [10.3 排序操作](https://github.com/zhangmingkai4315/golang-essentials/tree/master/10-%E7%BC%96%E5%86%99%E5%BA%94%E7%94%A8#103-%E6%8E%92%E5%BA%8F%E6%93%8D%E4%BD%9C)
- [10.4 使用Go语言编写密码加密和验证](https://github.com/zhangmingkai4315/golang-essentials/tree/master/10-%E7%BC%96%E5%86%99%E5%BA%94%E7%94%A8#104-%E4%BD%BF%E7%94%A8go%E8%AF%AD%E8%A8%80%E7%BC%96%E5%86%99%E5%AF%86%E7%A0%81%E5%8A%A0%E5%AF%86%E5%92%8C%E9%AA%8C%E8%AF%81)
- [10.5 应用参数传递]()

### [11.并发编程](https://github.com/zhangmingkai4315/golang-essentials/tree/master/11-%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B)

- [11.1 并发不是并行](https://github.com/zhangmingkai4315/golang-essentials/tree/master/11-%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B#111-%E5%B9%B6%E5%8F%91%E4%B8%8D%E6%98%AF%E5%B9%B6%E8%A1%8C)
- [11.2 goroutine](https://github.com/zhangmingkai4315/golang-essentials/tree/master/11-%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B#112-goroutine)
- [11.3 WaitGroup](https://github.com/zhangmingkai4315/golang-essentials/tree/master/11-%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B#113-waitgroup)
- [11.4 并发环境下的共享变量](https://github.com/zhangmingkai4315/golang-essentials/tree/master/11-%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B#114-%E5%B9%B6%E5%8F%91%E7%8E%AF%E5%A2%83%E4%B8%8B%E7%9A%84%E5%85%B1%E4%BA%AB%E5%8F%98%E9%87%8F)
- [11.5 Mutex锁机制](https://github.com/zhangmingkai4315/golang-essentials/tree/master/11-%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B#115-mutex%E9%94%81%E6%9C%BA%E5%88%B6)
- [11.6 Atomic原子操作](https://github.com/zhangmingkai4315/golang-essentials/tree/master/11-%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B#116-atomic%E5%8E%9F%E5%AD%90%E6%93%8D%E4%BD%9C)

### [12.Channel](https://github.com/zhangmingkai4315/golang-essentials/tree/master/12-Channel)
- [12.1 channel基础](https://github.com/zhangmingkai4315/golang-essentials/tree/master/12-Channel#121-channel%E5%9F%BA%E7%A1%80)
- [12.2 单向channel](https://github.com/zhangmingkai4315/golang-essentials/tree/master/12-Channel#122-%E5%8D%95%E5%90%91channel)
- [12.3 select语句](https://github.com/zhangmingkai4315/golang-essentials/tree/master/12-Channel#123-select%E8%AF%AD%E5%8F%A5)
- [12.4 Fan in和Fan out模型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/12-Channel#124-fan-in%E5%92%8Cfan-out%E6%A8%A1%E5%9E%8B)
- [12.5 context上下文管理](https://github.com/zhangmingkai4315/golang-essentials/tree/master/12-Channel#125-context%E4%B8%8A%E4%B8%8B%E6%96%87%E7%AE%A1%E7%90%86)
- [12.6 限速器的使用]()

### [13.错误处理](https://github.com/zhangmingkai4315/golang-essentials/tree/master/13-%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86)
- [13.1 error类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/13-%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86#131-error%E7%B1%BB%E5%9E%8B)
- [13.2 错误检查](https://github.com/zhangmingkai4315/golang-essentials/tree/master/13-%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86#132-%E9%94%99%E8%AF%AF%E6%A3%80%E6%9F%A5)
- [13.3 错误处理](https://github.com/zhangmingkai4315/golang-essentials/tree/master/13-%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86#133-%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86)
- [13.4 defer和recover](https://github.com/zhangmingkai4315/golang-essentials/tree/master/13-%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86#134-defer%E5%92%8Crecover)
- [13.5 自定义错误](https://github.com/zhangmingkai4315/golang-essentials/tree/master/13-%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86#135-%E8%87%AA%E5%AE%9A%E4%B9%89%E9%94%99%E8%AF%AF)
### [14.程序文档](https://github.com/zhangmingkai4315/golang-essentials/tree/master/14-%E7%A8%8B%E5%BA%8F%E6%96%87%E6%A1%A3)
- [14.1 go doc命令](https://github.com/zhangmingkai4315/golang-essentials/tree/master/14-%E7%A8%8B%E5%BA%8F%E6%96%87%E6%A1%A3#141-go-doc%E5%91%BD%E4%BB%A4)
- [14.2 godoc命令](https://github.com/zhangmingkai4315/golang-essentials/tree/master/14-%E7%A8%8B%E5%BA%8F%E6%96%87%E6%A1%A3#142-godoc%E5%91%BD%E4%BB%A4)
- [14.3 使用godoc索引](https://github.com/zhangmingkai4315/golang-essentials/tree/master/14-%E7%A8%8B%E5%BA%8F%E6%96%87%E6%A1%A3#143-%E4%BD%BF%E7%94%A8godoc%E7%B4%A2%E5%BC%95)
- [14.4 编写文档](https://github.com/zhangmingkai4315/golang-essentials/tree/master/14-%E7%A8%8B%E5%BA%8F%E6%96%87%E6%A1%A3#144-%E7%BC%96%E5%86%99%E6%96%87%E6%A1%A3)
### [15.测试](https://github.com/zhangmingkai4315/golang-essentials/tree/master/15-%E6%B5%8B%E8%AF%95)
- [15.1 测试模块testing.T](https://github.com/zhangmingkai4315/golang-essentials/tree/master/15-%E6%B5%8B%E8%AF%95#151-%E6%B5%8B%E8%AF%95%E6%A8%A1%E5%9D%97testingt)
- [15.2 编写测试](https://github.com/zhangmingkai4315/golang-essentials/tree/master/15-%E6%B5%8B%E8%AF%95#152-%E7%BC%96%E5%86%99%E6%B5%8B%E8%AF%95)
- [15.3 表格测试](https://github.com/zhangmingkai4315/golang-essentials/tree/master/15-%E6%B5%8B%E8%AF%95#153-%E8%A1%A8%E6%A0%BC%E6%B5%8B%E8%AF%95)
- [15.4 编写示例测试](https://github.com/zhangmingkai4315/golang-essentials/tree/master/15-%E6%B5%8B%E8%AF%95#154-%E7%BC%96%E5%86%99%E7%A4%BA%E4%BE%8B%E6%B5%8B%E8%AF%95)
- [15.5 语法检查](https://github.com/zhangmingkai4315/golang-essentials/tree/master/15-%E6%B5%8B%E8%AF%95#155-%E8%AF%AD%E6%B3%95%E6%A3%80%E6%9F%A5)
- [15.6 压力测试](https://github.com/zhangmingkai4315/golang-essentials/tree/master/15-%E6%B5%8B%E8%AF%95#156-%E5%8E%8B%E5%8A%9B%E6%B5%8B%E8%AF%95)
- [15.7 覆盖测试](https://github.com/zhangmingkai4315/golang-essentials/tree/master/15-%E6%B5%8B%E8%AF%95#157-%E8%A6%86%E7%9B%96%E6%B5%8B%E8%AF%95)

## 其他参考资料

##### 学习网站
- [Exaples实例](https://gobyexample.com/)
- [语言规范](https://golang.org/ref/spec)
- [Effictive-Go](https://golang.org/doc/effective_go.html)
- [Awesome-Go](https://github.com/avelino/awesome-go)

##### 技术博客

- [ardanlabs](https://www.ardanlabs.com/blog/)
- [doxsey](http://www.doxsey.net)

##### 书籍推荐：
- 《Go语言程序设计》
- 《Go In Action》
- 《Go web编程》

##### 社区论坛

- [Golang中文网](https://studygolang.com/)
- [Golangbridge](https://forum.golangbridge.org/)
