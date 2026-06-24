# gopl.io
《Go程序设计语言》(The Go Programming Language)书中代码
* 本书网站：<https://www.gopl.io/>
* 源代码：<https://github.com/adonovan/gopl.io/>
* 笔记：<https://zzy979.github.io/posts/gopl-note-index/>

## 依赖
Go 1.25

## 运行方式
将本仓库克隆到$GOPATH/src目录中，目录结构如下：

```
$GOPATH/
  src/
    gopl.io/
      ch1/
        helloworld/
          main.go
        ...
  pkg/
    windows_amd64/
  bin/
    helloworld
    ...
```

在gopl.io目录中使用`go run`命令运行示例程序。例如：

```shell
$ go run gopl.io/ch1/helloworld
Hello, 世界
```

## 单元测试
在项目根目录中执行以下命令运行所有单元测试：

```shell
$ make test
```

运行单个章节的单元测试：

```shell
$ go test -v ch1
```

## 代码目录
### 第1章 入门
* [gopl.io/ch1/helloworld](ch1/helloworld/main.go) Hello World
* [gopl.io/ch1/echo1](ch1/echo1/main.go) 打印命令行参数1
* [gopl.io/ch1/echo2](ch1/echo2/main.go) 打印命令行参数2
* [gopl.io/ch1/echo3](ch1/echo3/main.go) 打印命令行参数3
* [gopl.io/ch1/exec1-1](ch1/exec1-1/main.go) 练习1.1
* [gopl.io/ch1/exec1-2](ch1/exec1-2/main.go) 练习1.2
* [gopl.io/ch1/exec1-3](ch1/exec1-3/main.go) 练习1.3
* [gopl.io/ch1/dup1](ch1/dup1/main.go) 查找重复的行1
* [gopl.io/ch1/dup2](ch1/dup2/main.go) 查找重复的行2
* [gopl.io/ch1/dup3](ch1/dup3/main.go) 查找重复的行3
* [gopl.io/ch1/exec1-4](ch1/exec1-4/main.go) 练习1.4
* [gopl.io/ch1/lissajous](ch1/lissajous/main.go) 李萨如图形
* [gopl.io/ch1/exec1-5](ch1/exec1-5/main.go) 练习1.5
* [gopl.io/ch1/fetch](ch1/fetch/main.go) 获取URL
* [gopl.io/ch1/exec1-7](ch1/exec1-7/main.go) 练习1.7~1.9
* [gopl.io/ch1/fetchall](ch1/fetchall/main.go) 并发获取URL
* [gopl.io/ch1/server1](ch1/server1/main.go) Web服务器1
* [gopl.io/ch1/server2](ch1/server2/main.go) Web服务器2
* [gopl.io/ch1/server3](ch1/server3/main.go) Web服务器3
* [gopl.io/ch1/exec1-12](ch1/exec1-12/main.go) 练习1.12

### 第2章 程序结构
* [gopl.io/ch2/boiling](ch2/boiling/main.go) 打印温度
* [gopl.io/ch2/ftoc](ch2/ftoc/main.go) 打印温度转换
* [gopl.io/ch2/echo4](ch2/echo4/main.go) 打印命令行参数4（支持标志）
* [gopl.io/ch2/tempconv0](ch2/tempconv0/celsius.go) 温度转换
* [gopl.io/ch2/tempconv](ch2/tempconv) 温度转换、练习2.1
* [gopl.io/ch2/cf](ch2/cf/main.go) 导入包
* [gopl.io/ch2/exec2-2](ch2/exec2-2/main.go) 练习2.2
* [gopl.io/ch2/popcount](ch2/popcount/main.go) 二进制位计数、练习2.3~2.5
