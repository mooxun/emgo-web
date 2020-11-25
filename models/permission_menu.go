package models

import (
	"fmt"

	"github.com/mooxun/emgo-web/pkg/cfg"
)

type PermissionMenu struct {
	Id           uint64   `gorm:"primaryKey" json:"id" title:"ID"`
	AuthProvider string   `gorm:"size:50;comment:系统权限认证模块" json:"auth_provider"`
	Path         string   `gorm:"comment:前端路由地址" json:"path"`
	Name         string   `gorm:"comment:菜单名称" json:"name"`
	Rules        []string `gorm:"type:json;comment:权限api" json:"rules"`
}

// 迁移文件配置
func PermissionMenuTableComment() (o, c string) {
	return "gorm:table_options", fmt.Sprintf("ENGINE=%s comment '权限菜单表'", cfg.App.Mysql.Engine)
}

// 表名
func (PermissionMenu) TableName() string {
	return fmt.Sprintf("%spermission_menu", cfg.App.Mysql.Prefix)
}
