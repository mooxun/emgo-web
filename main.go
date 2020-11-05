package main

import (
	"github.com/mooxun/emgo-web/pkg/conf"
	"github.com/mooxun/emgo-web/pkg/redis"
	"github.com/mooxun/emgo-web/routes"
)

func init() {
	conf.Init()
	redis.Init()
}

func main() {
	// 路由初始化
	r := routes.Init()
	r.Run(":8000")
}
