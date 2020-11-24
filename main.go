package main

import (
	"github.com/mooxun/emgo-web/pkg/cfg"
	"github.com/mooxun/emgo-web/pkg/logger"
	"github.com/mooxun/emgo-web/pkg/mongodb"
	"github.com/mooxun/emgo-web/pkg/orm"
	"github.com/mooxun/emgo-web/pkg/redis"
	"github.com/mooxun/emgo-web/routes"
)

func init() {
	cfg.Init()
	logger.Init()
	redis.Init()
	mongodb.Init()
	orm.Init()
}

func main() {
	// 路由初始化
	r := routes.Init()
	r.Run(":5009")
}
