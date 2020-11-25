package models

import (
	"fmt"
	"time"

	"github.com/mooxun/emgo-web/pkg/cfg"
)

type PlatformRole struct {
	Id        uint64    `gorm:"primaryKey" json:"id" title:"ID"`
	Name      string    `gorm:"size:50;comment:角色名称" json:"username" validate:"required" title:"角色名称"`
	Status    int       `gorm:"type:tinyint;default:1;comment:是否启用" json:"status"`
	Listorder int       `gorm:"type:tinyint;default:0;comment:是否启用" json:"listorder"`
	Menus     []string  `gorm:"type:json;comment:前端选中权限菜单" json:"menus"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at" json:"updated_at"`
}

// 迁移文件配置
func PlatformRoleTableComment() (o, c string) {
	return "gorm:table_options", fmt.Sprintf("ENGINE=%s comment '平台角色表'", cfg.App.Mysql.Engine)
}

// 表名
func (PlatformRole) TableName() string {
	return fmt.Sprintf("%splatform_role", cfg.App.Mysql.Prefix)
}
