package orm

import (
	"fmt"

	"github.com/mooxun/emgo-web/pkg/cfg"
	"github.com/mooxun/emgo-web/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.App.Mysql.Username, cfg.App.Mysql.Password,
		cfg.App.Mysql.Host, cfg.App.Mysql.Port,
		cfg.App.Mysql.Database, cfg.App.Mysql.Charset)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Error("mysql connect error: ", err)
	}

	if DB.Error != nil {
		logger.Error("database error: ", err)
	}

	DB.Logger.LogMode(1)
}