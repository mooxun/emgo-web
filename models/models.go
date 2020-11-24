package models

import (
	"github.com/mooxun/emgo-web/pkg/orm"
)

// 查找一条信息
func Find(id int64, model interface{}) error {
	err := orm.DB.Where("id=?", id).First(model).Error
	return err
}

// 创建
func Create(model interface{}) error {
	err := orm.DB.Create(model).Error
	return err
}

// 更新
func Update(id int64, model interface{}) error {
	err := orm.DB.Where("id=?", id).Updates(model).Error
	return err
}

// 删除
func Delete(id int64, model interface{}) error {
	err := orm.DB.Where("id=?", id).Delete(model).Error
	return err
}
