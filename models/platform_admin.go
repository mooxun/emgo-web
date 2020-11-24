package models

import (
	"fmt"
	"time"

	"github.com/mooxun/emgo-web/pkg/cfg"
)

type PlatformAdmin struct {
	ID        uint64    `gorm:"primaryKey" json:"id" title:"ID"`
	RoleId    int8      `gorm:"index;comment:用户角色ID" json:"role_id"`
	Username  string    `gorm:"size:20;unique;comment:用户名" json:"username" validate:"required" label:"用户名"`
	Password  string    `gorm:"varchar:150;comment:用户密码" json:"password" validate:"required" label:"用户密码"`
	Status    int       `gorm:"tinyint:3;comment:是否启用" json:"status"`
	IsRoot    int       `gorm:"tinyint:3;comment:是否是超级管理员" json:"is_root"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at" json:"updated_at"`
}

// 平台管理员表迁移文件配置
func PlatformAdminTableComment() (o, c string) {
	return "gorm:table_options", fmt.Sprintf("ENGINE=%s AUTO_INCREMENT = 10000 comment '平台管理员账户表'", cfg.App.Mysql.Engine)
}

// 表名
func (PlatformAdmin) TableName() string {
	return fmt.Sprintf("%splatform_admin", cfg.App.Mysql.Prefix)
}