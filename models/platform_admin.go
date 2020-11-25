package models

import (
	"fmt"
	"time"

	"github.com/mooxun/emgo-web/pkg/cfg"
	"github.com/mooxun/emgo-web/pkg/orm"
)

type PlatformAdmin struct {
	Id        uint64    `gorm:"primaryKey" json:"id" title:"ID" show:"detail,lists"`
	RoleId    int8      `gorm:"index;comment:用户角色ID" json:"role_id" show:"detail,lists"`
	Username  string    `gorm:"size:20;unique;comment:用户名" json:"username" validate:"required" title:"用户名" show:"detail,lists"`
	Password  string    `gorm:"varchar:150;comment:用户密码" json:"password" validate:"required"`
	Status    int       `gorm:"type:tinyint;comment:是否启用" json:"status" title:"是否启用" show:"detail,lists"`
	IsRoot    int       `gorm:"type:tinyint;comment:是否是超级管理员" json:"is_root" title:"是否是超级管理员" show:"detail,lists"`
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

func FindAdmin(username string) (*PlatformAdmin, error) {
	var admin PlatformAdmin
	if err := orm.DB.Where("username=?", username).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

// jwt auth 验证账号是否正确
func CheckAdmin(id uint64, username string) bool {
	var admin PlatformAdmin
	if err := orm.DB.Where("id=?", id).Where("username=?", username).First(&admin).Error; err != nil {
		return false
	}
	return true
}
