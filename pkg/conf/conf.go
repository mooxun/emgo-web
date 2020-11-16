package conf

import (
	"fmt"
	"time"

	"github.com/mooxun/emgo-web/pkg/gofile"
	"github.com/mooxun/emgo-web/pkg/gopath"
	"github.com/spf13/viper"
)

var App Config

type Config struct {
	Logger Logger
	Redis  Redis
	Mongodb Mongodb
}

// Logger 日志配置
type Logger struct {
	Level        string        // 日志级别
	Path         string        // 路径
	Name         string        // 文件名称
	Console      bool          // 是否输出到控制台
	MaxAge       time.Duration // 保存多久的日志，默认15天
	RotationTime time.Duration // 多久分割一次日志
	Caller       bool          // 是否打印文件行号
	SplitLevel   bool          // 是否把不同级别的日志打到不同文件
}

// Redis 配置
type Redis struct {
	Host     string
	Port     int
	Password string
}

type Mongodb struct {
	Uri      string
	Database string
	Timeout  time.Duration
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
