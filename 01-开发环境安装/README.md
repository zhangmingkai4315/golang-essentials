### 1. 开发环境准备

本章节主要介绍如何在Windows,Linux以及MacOS平台上部署Go语言的开发环境,为接下来的程序学习做准备。同时下面介绍的开发软件以及相关辅助工具绝大部分都提供三个平台通用的版本，因此即便是将来跨平台开发也不会有任何使用上的不便。

#### 1.1 Windows环境

Windows上部署开发环境，主要安装的软件清单如下：

- [Git-scm](https://git-scm.com/download/win): 提供代码的管理以及一个简单的GitBash工具，用于终端命令的执行输入
- [VSCode](https://code.visualstudio.com/): Go语言的开发IDE，编写代码的集成开发环境。
- [Go安装包](https://studygolang.com/dl): 包含Go二进制可执行文件，用于执行go程序。

一般这些软件会提供自动下载，如果需要手动选择的话，一定要选择Windows版本下载，由于都是安装包安装按照默认的配置即可，完成三个软件安装后，点击GitBash（来自于GitScm软件）进入终端运行环境，输入以下的命令确定是否可以正常显示：

```
$ go version
go version go1.11.5 windows/amd64
```

执行以下的命令来完成项目运行环境的搭建

```
$ cd ~ 
$ mkdir -p workspace/go/{bin, pkg, src/github.com}
```
接下来需要创建一个[github.com](https://github.com)的账号用于未来的代码版本管理，完成账号的创建, 按照如下方式创建目录：

```
$ mkdir workspace/go/src/github.com/{替换为自己的账号ID}/
$ cd workspace/go/src/github.com/{替换为自己的账号ID}/
```

该文件夹用来创建Go程序项目使用，所有的个人项目都放在该文件夹下面，此时还需要设置环境变量，用于go运行环境使用，具体的修改方式：**打开控制面板->所有控制面板项->系统->高级系统设置->环境变量**， 找到上半部分的用户变量并选择GOPATH，如果没有的话新增一条记录，变量名称为GOPATH，变量值为
```
%USERPROFILE%\workspace\go
```
同时环境变量Path变量中新建一条记录：
```
%USERPROFILE%\workspace\go\bin
```

设置GOPATH的主要目的用于Go语言找到开发项目的位置，进而用于管理所有的依赖项和编译程序使用。修改完毕后点击**确定**退出，并重新启动GitBash,执行下面的语句用于检查是否设置成功, 首先进到自己账号下的目录中执行如下的操作, 第一步下载一个第三方的库，该库下载完成后源代码会自动放入$GOPATH/src/github.com/jinzhu/now目录下面，将来编写的代码如果依赖了该库，会自动去该位置查找。
```
$ go get -u github.com/jinzhu/now
$ mkdir example
$ code example
```
执行最后一步应该会自动的打开VSCode编程环境，创建一个新的文件main.go,并将如下的代码写入文件中：

```go
package main

import(
	"fmt"
	"github.com/jinzhu/now"
)

func main(){
	fmt.Println(now.BeginningOfYear())
}
```

执行下面的命令运行go程序，如果能获得输出结果则代表运行环境搭建成功，可以继续下面的课程学习内容了。

```
go run example/main.go
2019-01-01 00:00:00 +0800 CST
```

#### 1.2 Linux环境



#### 1.3 MacOS环境

