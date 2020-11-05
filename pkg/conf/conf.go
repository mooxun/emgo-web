package conf

import (
	"fmt"
	"github.com/mooxun/emgo-web/pkg/gofile"
	"github.com/mooxun/emgo-web/pkg/gopath"
	"github.com/spf13/viper"
)

var App Config

type AppPath struct {
	ExecPath string
}

type Config struct {
	Redis Redis
}

type Redis struct {
	Host     string
	Port     int
	Password string
}

func Init() {
	// 获取程序运行目录
	dir, err := gopath.CurrentPath()
	if err != nil {
		fmt.Println("get os path error: ", err)
	}

	cfgFile := fmt.Sprintf("%s/config/app.yaml", dir)
	if ok := gofile.FileExists(cfgFile); !ok {
		cfgFile = "./config/app.yaml"
	}

	//读取yaml文件
	v := viper.New()
	v.SetConfigFile(cfgFile)

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("config read error: ", err.Error())
	}

	if err := v.Unmarshal(&App); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("loading config success")
}
