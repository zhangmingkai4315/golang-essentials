# golang-essentials

Go语言起源于Google，2007年的时候Google当时在自己的生产环境中主要使用的语言以是C++, Java和Python为主，但是发现随着软硬件系统的不断升级，部署规模的增加，这些语言都或多或少的存在一些问题，比如构建时间过长，依赖管理复杂，并发处理不够灵活，服务更新和升级繁琐等。因此为了解决这些问题，Google内部的技术人员开始尝试去开发一个新的编程语言去逐步解决这些问题，这就是Go语言的最初的设计初衷。


简单的来介绍Go语言的话就是： Go语言是一个编译型，并发，自动垃圾回收的静态语言，尽管由Google最初开发，但是已开源并获的大规模的使用。

在过去十年间，Go语言获得了越来越多的人的关注。据StackOverflow的[2018年技术调查报告](https://insights.stackoverflow.com/survey/2018/)显示，2018年Go语言与Python和JavaScript编程语言一起成为众多技术人员最想学习的编程语言。作为一名DevOps工程师，当前线上部署的大规模运维支撑系统越来越多的被Go语言编写的项目所占据，比如Docker, Kubernetes, Prometheus, Grafana，CoreDNS等等，这些项目已经深入到我们的日常工作之中。


Go语言的开发团队成员也都是一些知名的计算机行业专家其中包括：Rob Pike ，曾参与Unix 以及UTF-8的开发，Robert Griesemer 早期Pascal语言的开发人员，Ken Thompson曾参与Unix，B语言及C语言设计工作。这些技术专家的加入以及Google公司的背景，使得对于Go语言自诞生之初，就获得了众多的关注，经过十几年的发展，Go语言也其稳定可靠，干净效率的编程发生赢多了很多工程师的喜爱。

>  I think Node is not the best system to build a massive server web. I would use Go for that. And honestly, that’s the reason why I left Node. It was the realization that: oh, actually, this is not the best server-side system ever.                
                                    **- Ryan Dahl（NodeJS作者）**


使用Go语言可以进行系统及网络方面的程序设计开发，比如开发分布式系统，命令行程序，web后台系统，以及一些图像处理，加密算法等需求。



## 主要内容

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
### [4.基本类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/04-%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B)

- [4.1 number数值类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/04-%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B#41-number%E6%95%B0%E5%80%BC%E7%B1%BB%E5%9E%8B)
- [4.2 string字符类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/04-%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B#42-string%E5%AD%97%E7%AC%A6%E7%B1%BB%E5%9E%8B)
- [4.3const常量类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/04-%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B#43-const%E5%B8%B8%E9%87%8F%E7%B1%BB%E5%9E%8B)
- [4.3 类型转换](https://github.com/zhangmingkai4315/golang-essentials/tree/master/04-%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B#44-%E7%B1%BB%E5%9E%8B%E8%BD%AC%E6%8D%A2)

### [5.interface类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/05-interface%E7%B1%BB%E5%9E%8B)

- [5.1 接口Interface](https://github.com/zhangmingkai4315/golang-essentials/tree/master/05-interface%E7%B1%BB%E5%9E%8B#51-%E6%8E%A5%E5%8F%A3interface)
- [5.2 类型断言](https://github.com/zhangmingkai4315/golang-essentials/tree/master/05-interface%E7%B1%BB%E5%9E%8B#52-%E7%B1%BB%E5%9E%8B%E6%96%AD%E8%A8%80)
- [5.3 实现sort接口](https://github.com/zhangmingkai4315/golang-essentials/tree/master/05-interface%E7%B1%BB%E5%9E%8B#53-%E5%AE%9E%E7%8E%B0sort%E6%8E%A5%E5%8F%A3)

### [6.数组切片和Map类型](https://github.com/zhangmingkai4315/golang-essentials/tree/master/06-%E6%95%B0%E7%BB%84%E5%88%87%E7%89%87%E5%92%8CMap%E7%B1%BB%E5%9E%8B)
- []()
- []()
- []()
- []()
- []()
### [7.结构体](https://github.com/zhangmingkai4315/golang-essentials/tree/master/07-%E7%BB%93%E6%9E%84%E4%BD%93)
### [8.函数](https://github.com/zhangmingkai4315/golang-essentials/tree/master/08-%E5%87%BD%E6%95%B0)
### [9.指针](https://github.com/zhangmingkai4315/golang-essentials/tree/master/09-%E6%8C%87%E9%92%88)
### [10.编写应用程序](https://github.com/zhangmingkai4315/golang-essentials/tree/master/10-%E7%BC%96%E5%86%99%E5%BA%94%E7%94%A8)
### [11.并发编程](https://github.com/zhangmingkai4315/golang-essentials/tree/master/11-%E5%B9%B6%E5%8F%91%E7%BC%96%E7%A8%8B)
### [12.Channel](https://github.com/zhangmingkai4315/golang-essentials/tree/master/12-Channel)
### [13.错误处理](https://github.com/zhangmingkai4315/golang-essentials/tree/master/13-%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86)
### [14.程序文档](https://github.com/zhangmingkai4315/golang-essentials/tree/master/14-%E7%A8%8B%E5%BA%8F%E6%96%87%E6%A1%A3)
### [15.测试](https://github.com/zhangmingkai4315/golang-essentials/tree/master/15-%E6%B5%8B%E8%AF%95)


## 其他参考资料

##### 学习网站
- [Exaples实例](https://gobyexample.com/)
- [语言规范](https://golang.org/ref/spec)
- [Effictive-Go](https://golang.org/doc/effective_go.html)
- [Awesome-Go](https://github.com/avelino/awesome-go)

##### 技术博客

https://www.ardanlabs.com/blog/
http://www.doxsey.net

##### 书籍推荐：
《Go语言程序设计》
《Go In Action》
《Go web编程》

##### 社区论坛

https://forum.golangbridge.org/
