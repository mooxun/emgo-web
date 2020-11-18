package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/mooxun/emgo-web/pkg/cfg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type rotateHandler struct {
	path     string
	linkName string
}

var zapLogger = zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()), zapcore.AddSync(os.Stdout), zap.DebugLevel))
var SugaredLogger = zapLogger.Sugar()
var Logs = zapLogger

var printCaller = false

// Init 初始化日志库
func Init() {
	cfg := &cfg.App.Logger
	if err := newLog(cfg); err != nil {
		fmt.Println("logger init error")
	}
}

func newLog(conf *cfg.Logger) error {
	// 默认保存最近15天日志
	if conf.MaxAge == 0 {
		conf.MaxAge = time.Hour * 24 * 15
	}
	if conf.RotationTime == 0 {
		conf.RotationTime = time.Hour
	}
	printCaller = conf.Caller

	// 建立日志目录
	if err := os.MkdirAll(conf.Path+"/", os.ModePerm); err != nil {
		fmt.Println("init log path error.")
		return err
	}
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	level := new(zapcore.Level)
	err := level.Set(conf.Level)
	if err != nil {
		return err
	}

	lv := *level

	var cores []zapcore.Core

	if conf.SplitLevel {
		if lv <= zapcore.DebugLevel {
			debugLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl <= zapcore.DebugLevel
			})

			debugWriter, err := getWriter(conf.Path+"/"+conf.Name+"_debug", conf)
			if err != nil {
				return err
			}
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(debugWriter), debugLevel))
		}
		if lv <= zapcore.InfoLevel {
			infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl < zapcore.WarnLevel && lvl > zapcore.DebugLevel
			})

			infoWriter, err := getWriter(conf.Path+"/"+conf.Name+"_info", conf)
			if err != nil {
				return err
			}

			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel))
		}

		if lv <= zapcore.WarnLevel {
			warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
				return lvl >= zapcore.WarnLevel
			})
			warnWriter, err := getWriter(conf.Path+"/"+conf.Name+"_error", conf)
			if err != nil {
				return err
			}
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(warnWriter), warnLevel))
		} else {
			// 级别是 error 以上
			errorLevel := lv
			errorWriter, err := getWriter(conf.Path+"/"+conf.Name+"_error", conf)
			if err != nil {
				return err
			}
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel))
		}
	} else {
		writer, err := getWriter(conf.Path+"/"+conf.Name, conf)
		if err != nil {
			return err
		}
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(writer), lv))
	}

	if conf.Console {
		consoleWriter := os.Stdout
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(consoleWriter), lv))
	}

	// 最后创建具体的Logger
	core := zapcore.NewTee(cores...)

	Logs = zap.New(core)
	SugaredLogger = Logs.Sugar()

	return nil
}

func getWriter(filename string, conf *cfg.Logger) (io.Writer, error) {
	hook, err := rotatelogs.New(
		filename+"-%Y-%m-%d.log",
		rotatelogs.WithHandler(&rotateHandler{
			path:     conf.Path,
			linkName: filename + ".log",
		}),
		rotatelogs.WithMaxAge(conf.MaxAge),
		rotatelogs.WithRotationTime(conf.RotationTime),
	)

	if err != nil {
		return nil, err
	}
	return hook, nil
}

// 创建一个符号链接文件，链接到最新的日志文件，方便查看最新日志
func (r *rotateHandler) Handle(e rotatelogs.Event) {
	ev, ok := e.(*rotatelogs.FileRotatedEvent)
	if ok {
		_ = os.Remove(r.linkName)
		current := ev.CurrentFile()
		rel, _ := filepath.Rel(r.path, current)
		err := os.Symlink(rel, r.linkName)
		if err != nil {
			// 如果是windows，其实通常都是失败的，所以干脆不要在 win 显示错误了
			if runtime.GOOS != "windows" {
				fmt.Println(err)
			}
			return
		}
	}
}
