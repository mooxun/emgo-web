package cfg

import "time"

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
