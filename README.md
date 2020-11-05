## emgo-web

emgo-web 是一套简单易用的Go语言业务框架，主要是提供 API 服务。

## Features
- HTTP服务基于[gin](https://github.com/gin-gonic/gin) 进行模块化设计，简单易用、核心足够轻量；
- 数据库组件 [GORM](https://github.com/jinzhu/gorm)
- 配置文件解析库 [Viper](https://github.com/spf13/viper)
- 使用 [JWT](https://jwt.io/) 进行身份鉴权认证
- 校验器使用 [Validator](https://github.com/go-playground/validator) 也是 Gin 框架默认的校验器
- 包管理工具 [Go Modules](https://github.com/golang/go/wiki/Modules)
- 使用 make 来管理 Go 工程