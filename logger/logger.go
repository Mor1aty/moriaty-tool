package logger

import (
	"fmt"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"runtime"
	"time"
)

var lg *zap.Logger

type LogConfig struct {
	Console      bool
	File         bool
	Level        string
	Filename     string
	RotationTime int
	MaxAges      int
}

// InitLogger 初始化 Logger
func InitLogger(cfg *LogConfig) error {
	if !cfg.Console && !cfg.File {
		log.Println("console and file is false, Logs need to be output in one place")
		return fmt.Errorf("console and file is false, Logs need to be output in one place")
	}

	encoder := getEncoder()
	var l = new(zapcore.Level)
	err := l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		log.Printf("unmarshal level failed, err: %v", err)
		return err
	}

	writeSyncerList := make([]zapcore.WriteSyncer, 0)
	if cfg.File {
		fileWriteSyncer, err := getLogWriter(cfg.Filename, cfg.RotationTime, cfg.MaxAges)
		if err != nil {
			log.Printf("rotate logs failed, err: %v", err)
			return err
		}
		writeSyncerList = append(writeSyncerList, fileWriteSyncer)
	}
	if cfg.Console {
		writeSyncerList = append(writeSyncerList, zapcore.AddSync(os.Stdout))
	}

	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncerList...), l)
	lg = zap.New(core, zap.AddCaller())

	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(lg)
	return nil
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, rotationTime, maxAge int) (zapcore.WriteSyncer, error) {
	// 生成 rotateLogs 的 Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log 是指向最新日志的链接
	// 保存 7 天内的日志，每 1 小时(整点)分割一次日志
	// 这里框架设置的日期格式
	var hook *rotateLogs.RotateLogs
	var err error
	if runtime.GOOS == "linux" {
		hook, err = rotateLogs.New(
			filename+".%Y%m%d%H",
			rotateLogs.WithMaxAge(time.Hour*time.Duration(maxAge)),
			rotateLogs.WithRotationTime(time.Hour*time.Duration(rotationTime)),
			rotateLogs.WithLinkName(filename),
		)
	} else {
		hook, err = rotateLogs.New(
			filename+".%Y%m%d%H",
			rotateLogs.WithMaxAge(time.Hour*time.Duration(maxAge)),
			rotateLogs.WithRotationTime(time.Hour*time.Duration(rotationTime)),
		)
	}
	if err != nil {
		return nil, err
	}
	return zapcore.AddSync(hook), nil
}

