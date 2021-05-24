## go-skeleton

`go-skeleton` 是基于golang的web骨架项目，可基于该骨架项目快速搭建web服务。

## 项目结构

##### （1）routes

路由创建与统一管理。

##### （2）api

控制器，控制程序执行流程，将路由转交给后面的业务逻辑层进行处理，并将处理结果返回给api接口调用者。

##### （3）service

主体业务逻辑层，目前仅包含用户模块，可根据自身需求添加其它模块。

##### （4）model

模型层，负责数据库CRUD操作。

##### （5）middleware

中间件，包含跨域中间件，jwt中间件，参数校验中间件，api路由日志中间件。

##### （6）serializer

序列化器，定义返回的json结构体，对返回结果进行序列化。

##### （7）utils

主要包含通用的模块，函数以及相关结构体。如日志处理模块，错误处理模块，请求参数与返回参数结构体定义，加解密方法等。

##### （8）config

存放配置文件。

##### （9）log

存放日志文件。

##### （10）data

存放项目相关的文件，如sql脚本文件。

##### （11）api-docs

接口文档。

## 技术栈

- gin (https://github.com/gin-gonic/gin) Go web 框架
- JWT (https://github.com/dgrijalva/jwt-go) JWT Middleware for Gin framework
- gorm (http://gorm.io/) Go 语言 orm 框架
- jio (https://github.com/faceair/jio) 参数校验框架
- logrus (https://github.com/sirupsen/logrus) 日志框架
- gjson (https://github.com/tidwall/gjson) golang json库
- cors (https://github.com/gin-contrib/cors) Gin框架的跨域中间件
- ini (https://github.com/go-ini/ini) golang操作ini文件的第三方库

## Docker部署

```shell
# 构建镜像
docker build -t go-skeleton:v1.0 .
# 运行
docker run --name web-demo -d -p 9116:9116 go-skeleton:v1.0
# 进入容器查看
docker exec -it [容器ID] /bin/sh
```

## 附录

[Golang 中文学习资料汇总](http://go.wuhaolin.cn/)