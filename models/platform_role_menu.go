package models

import (
	"fmt"

	"github.com/mooxun/emgo-web/pkg/cfg"
)

type PlatformRoleMenu struct {
	RoleId    int64  `gorm:"index;comment:角色ID" json:"role_id"`
	RoutePath string `gorm:"index;comment:权限菜单路由" json:"route_path"`
}

// 迁移文件配置
func PlatformRoleMenuTableComment() (o, c string) {
	return "gorm:table_options", fmt.Sprintf("ENGINE=%s comment '平台角色权限表'", cfg.App.Mysql.Engine)
}

// 表名
func (PlatformRoleMenu) TableName() string {
	return fmt.Sprintf("%splatform_role_menus", cfg.App.Mysql.Prefix)
}
