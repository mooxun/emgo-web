package models

import (
	"time"

	"github.com/mooxun/emgo-web/pkg/orm"
)

type IdField struct {
	Id uint64 `gorm:"primaryKey" json:"id" title:"ID"`
}

type TimeField struct {
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at" json:"updated_at"`
}

// 查找一条信息
func Find(id uint64, model interface{}) error {
	err := orm.DB.Where("id=?", id).First(model).Error
	return err
}

// 创建
func Create(model interface{}) error {
	err := orm.DB.Create(model).Error
	return err
}

// 更新
func Update(id uint64, model interface{}) error {
	err := orm.DB.Where("id=?", id).Updates(model).Error
	return err
}

// 删除
func Delete(id uint64, model interface{}) error {
	err := orm.DB.Where("id=?", id).Delete(model).Error
	return err
}
