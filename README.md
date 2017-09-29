# penguin
go的web框架

## 基本介绍
* penguin主要用于使用go语言进行web开发，penguin里面集成了gin和gorm等开源项目

## 集成组件
* config  
* log  
* mysql  
* mongoDB
* redis

## 使用方法
* 默认配置文件是go二进制执行文件当前目录下的penguin.toml，若想采用自定义配置文件，需要 可执行文件 -f example.toml  
* 配置文件根据关键字映射，如[模板][模板]所示。

[模板]:./penguin.toml

## 配置文件详解
### server
* addr: 服务的监听地址
* mode： 设置gin的模式
* pprof：是否开启debug功能
* origin: 跨域请求的时候，允许的ip

### mysql
* driver: 引擎类型
* source: 写库的地址
* slave: 各个读库

### mongo
* address: 地址
* userName: 用户
* password: 密码
* database: 库

### redis
* address: 地址
* password: 密码

### log
* file: log文件配置地址，可以不填，penguin会采用默认的配置