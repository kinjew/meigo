# MEIGO框架简介

MEIGO是基于gin框架，结合go语言特性设计的MVC风格的框架，使用了`Go Modules`包管理方式。 

## 一、MEIGO使用方法

1.安装go（mac环境下执行`brew install go`，其他环境参考[golang官网](<https://golang.google.cn/>)安装方法) 

```shell
brew install go
```
2.进入工作主目录（cd ~），设置环境变量   
```shell
cd ~
```
3.环境变量配置有两种方法，任选其一，推荐第一种方法。

3.1 第一种方法 (**推荐**)  
执行命令 `go version` ，确保 Go 的版本在 1.13 及以上，然后依次执行命令：

```bash
go env -w GOPATH=$HOME/go
go env -w GO111MODULE=on
go env -w GOPRIVATE=*.sprucetec.com
go env -w GOPROXY=https://goproxy.cn
```

然后执行命令 `go env` 查看是否修改成功。此修改保存在 go 的 环境变量 **`GOENV`** 中。

3.2 第二种方法   
在 `.bash_profile` 中增加以下环境变量(引入了私有包)

```shell
export GOPATH=$HOME/go
export GO111MODULE=on
export GOPRIVATE=*.sprucetec.com
export GOPROXY=https://goproxy.cn
```
然后执行`source .bash_profile`
```shell
source .bash_profile
```
4.新建工作目录，以meigo为例，并克隆项目到本地
```shell
cd go_workspace | mkdir meigo  
git clone 项目url meigo
```
将 conf/conf.toml.example 文件拷贝到 conf/conf.toml，并根据情况修改配置   
修改配置文件读取目录，具体为，找到library/config/configLoad.go文件，将viper.AddConfigPath("xxx")中的xxx替换为本地配置文件的绝对目录


5.编译项目  

本地环境中，开发环境、测试环境、发布环境 分别把 `conf/conf.toml` 中的 `runMode` 修改为 `debug` 、 `test`、 `release` 后，执行`bash scripts/dev.sh build`;  

**注：上线发布环境脚本为 `build.sh`  `start.sh`  `stop.sh` ，不用动**   

6.启动项目：在本地环境执行`bash scripts/dev.sh start` ；停止项目：执行`bash scripts/dev.sh stop`（线上项目启停可以让supervisor接管） 


7.配置  

数据库等相关配置文件位于 `conf/conf.toml`，框架默认mysql数据库账号为root，密码为空（root:@tcp(127.0.0.1:3306)，可根据情况适当修改。

8.执行如下命令向表中插入数据
```shell
  curl -v -X POST   http://127.0.0.1:8000/people   -H 'content-type: application/json'   -d '{ "first_name": "7777", "last_name": "胡88","city": "ufe"}'
```
9.访问网址测试  
```
http://127.0.0.1:8000、http://127.0.0.1:8000/people/
```

## 二、MEIGO目录结构说明
```shell
├── bin
├── cmd
├── conf
├── library
├── models
├── modules
├── release
├── routers
├── scripts
├── go.mod
├── go.sum
├── main.go
├── readme.md
```

1.main.go为框架入口文件  
2.cmd目录为命令行程序代码目录
3.conf目录存放数据库服务配置文件  
4.routers目录存放框架路由配置文件  
5.modules目录存放业务模块文件，不同业务模块放在不同子目录中  
6.models目录存放业务模型文件，不同业务模型放在不同子目录中  
7.scripts目录存放框架脚本   
8.libray目录存放框架库文件，非公共包，配置加载、服务初始化等位于该文件下   
9.bin目录存放编译成的可执行文件   
10.release目录存放应用文件，包含可执行文件、部署脚本以及必要的配置文件等

## 三、命令行
meigo引入cobra框架，该框架可方便编写命令和子命令程序，并可以接受命令行参数，可用于任务调度场景。执行框架统一的编译命令后，可执行如下命令查看帮助
```
./bin/cmd -h
```

执行如下命令进行一次函数执行

```
./bin/cmd -f index
```


## 四、其他说明

1.该框架需要go版本1.13及以上，框架使用go mod维护包版本及依赖。一个项目最好只有一个go.mod并放在项目根目录下  
2.为框架开发新的公共包需要定义go.mod文件（执行go mod init xxx，xxx为包名），同时可能需要在主目录的go.mod中定义替代（replace）路径，这样编译的时候才会从本地拉取包  
3.使用GoLand编辑器进行开发应该在GoLand下拉菜单（Mac中）找到`Preferences->Go->Go Modules(vgo)`中设置开启`Go Modules`支持并设置Proxy值为`https://goproxy.cn`  

## 异常问题解决方案  

1.如果第一次运行 meigo 项目出现以下错误：
`go: git.sprucetec.com/meigo/gin-context-ext@v0.0.2: reading git.sprucetec.com/meigo/gin-context-ext/go.mod at revision v0.0.2: unknown revision v0.0.2`  
该异常是 go get 机制引起，暂时解决方案为在本地终端执行以下 3 条命令：  

```bash
git config --global url.ssh://git@git.sprucetec.com:50022/golang-pkg/gin-context-ext.git.insteadof=https://git.sprucetec.com/golang-pkg/gin-context-ext.git  

git config --global url.ssh://git@go.uber.org/zap.insteadof=https://go.uber.org/zap  

git config --global url.ssh://git@git.sprucetec.com:50022.insteadof=https://git.sprucetec.com
```

