syslog 日志探针
-----

> **Glang版**

日志探针是一个日志采集组件，通过 syslog 采集系统日志。

### 态势感知平台-日志探针源码包含:

* log-prober: 包含一下处理（Golang ）
    * input: 采集解析日志服务 (基于开源[go-syslog](https://github.com/mcuadros/go-syslog)二次开发)
    * output: 输出 kafka 大数据接入
    * web: 接收态势感知平台下发采集配置及解析规则

### 目录结构

```
log-prober-go
    ├── README.md           说明文件
    ├── application.go      启动文件
    ├── common              公共方法
    ├── config.yaml         配置文件
    ├── domain              dto、实体类 
    ├── go.mod              依赖管理
    ├── handler             各处理解析文件
    ├── libs                第三方包
    ├── logs                日志
    ├── output              数据输出 kafka
    ├── utils               工具类
    └── test                测试文件
```

### Golang环境安装

**官方文档**
[https://golang.google.cn/doc/install
](https://golang.google.cn/doc/install
)

**mac OS安装Go**

下载并安装[Go for Mac](https://dl.google.com/go/go1.15.1.darwin-amd64.pkg)

验证安装结果

```
$ go version
go version go1.15.1 darwin/amd64
```

**linux 安装Go**

下载[Go for Linux](https://golang.org/dl/go1.15.8.linux-amd64.tar.gz)

```text
解压压缩包至/usr/local
$ tar -C /usr/local -xzf go1.15.8.linux-amd64.tar.gz
添加/usr/local/go/bin到环境变量
$ vi $HOME/.profile
$ export PATH=$PATH:/usr/local/go/bin
$ source $HOME/.profile
验证安装结果
$ go version
go version go1.15.1 linux/amd64
```

**Windows安装Go**

下载并安装[Go for Windows](https://golang.org/dl/go1.15.8.windows-amd64.msi)

验证安装结果

```text
$ go version
go version go1.15.1 windows/amd64
```

**其他**

更多操作系统安装见[https://golang.org/dl/]()

### Go Module设置

**Go Module介绍**

即Go Module是Golang管理依赖性的方式，像Java中的Maven，Android中的Gradle类似。

**MODULE配置**

查看`GO111MODULE`开启情况

```text
$ go env GO111MODULE
on
```

开启`GO111MODULE`，如果已开启（即执行`go env GO111MODULE`结果为`on`）请跳过。

`$ go env -w GO111MODULE="on"`

设置GOPROXY

`$ go env -w GOPROXY=https://goproxy.cn`

设置查看GOMODCACHE

`$ go env GOMODCACHE`

如果目录不为空或者`/dev/null`，请跳过。

`go env -w GOMODCACHE=$GOPATH/pkg/mod`

[1][参看文档Go Modules Reference](https://golang.google.cn/ref/mod)

### 部署

1. 进入项目目录

   `cd log-prober-go/`
2. 修改`config.yaml`文件等
    * kafka配置
    * 数据库
    * 运行端口
3. 下载依赖

   `go mod tidy`
4. 运行编译 检测编译是否通过

    1. `go build` 编译通过后 直接运行编译后文件

    2. `go build run ./` 编译运行
