package main

import (
	"github.com/mooxun/emgo-web/models"
	"github.com/mooxun/emgo-web/pkg/cfg"
	"github.com/mooxun/emgo-web/pkg/logger"
	"github.com/mooxun/emgo-web/pkg/orm"
)

func main() {
	cfg.Init()
	logger.Init()
	orm.Init()

	orm.DB.Set(models.PlatformAdminTableComment()).AutoMigrate(&models.PlatformAdmin{})
	orm.DB.Set(models.PlatformRoleTableComment()).AutoMigrate(&models.PlatformRole{})
	orm.DB.Set(models.PlatformRoleMenuTableComment()).AutoMigrate(&models.PlatformRoleMenu{})
	orm.DB.Set(models.PermissionMenuTableComment()).AutoMigrate(&models.PermissionMenu{})
}
