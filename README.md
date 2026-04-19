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

## 代码目录
### 第1章 入门
* [Hello World](ch1/helloworld/main.go)
* [打印命令行参数1](ch1/echo1/main.go)
* [打印命令行参数2](ch1/echo2/main.go)
* [打印命令行参数3](ch1/echo3/main.go)
* [查找重复的行1](ch1/dup1/main.go)
* [查找重复的行2](ch1/dup2/main.go)
* [查找重复的行3](ch1/dup3/main.go)
* [李萨如图形](ch1/lissajous/main.go)
* [获取URL](ch1/fetch/main.go)
* [并发获取URL](ch1/fetchall/main.go)
* [Web服务器1](ch1/server1/main.go)
* [Web服务器2](ch1/server2/main.go)
* [Web服务器3](ch1/server3/main.go)
