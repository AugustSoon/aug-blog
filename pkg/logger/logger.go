package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type ZapLogger struct {
	Self  *zap.Logger
	Sugar *zap.SugaredLogger
}

var Logger *ZapLogger

func InitZapWithConfig() *zap.Logger {
	hook := lumberjack.Logger{
		Filename:   viper.GetString("log.file_name"), // 日志文件路径
		MaxSize:    viper.GetInt("log.max_size"),     // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: viper.GetInt("log.max_backup"),   // 日志文件最多保存多少个备份
		MaxAge:     viper.GetInt("log.max_age"),      // 文件最多保存多少天
		Compress:   viper.GetBool("log.compress"),    // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zap.NewAtomicLevelAt(getLevel(viper.GetString("log.level"))),                    // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()

	// 开启文件及行号
	development := zap.Development()

	// 构造日志
	return zap.New(core, caller, development)

}

func getLevel(conf string) zapcore.Level {
	level := zap.InfoLevel

	switch conf {
	case "DEBUG":
		level = zap.DebugLevel
	case "WARN":
		level = zap.WarnLevel
	case "ERROR":
		level = zap.ErrorLevel
	}

	return level
}

func (zl *ZapLogger) Init() {
	log := InitZapWithConfig()

	Logger = &ZapLogger{
		Self:  log,
		Sugar: log.Sugar(),
	}
}

func (zl *ZapLogger) Close() {
	_ = Logger.Self.Sync()
}
