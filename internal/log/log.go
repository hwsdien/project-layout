package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"project-layout/internal/config"
	"sync"
)

type Logger struct {
	Log *zap.SugaredLogger
}

var log *Logger
var once sync.Once

// 目前只支持按文件大小切割，原因是按时间切割效率低且不能保证日志数据不被破坏
// 详情: https://github.com/natefinch/lumberjack/issues/54。
func getLogWriter() zapcore.WriteSyncer {
	conf := config.GetConfig()
	lumberJackLogger := &lumberjack.Logger{
		Filename:   conf.Viper.GetString("log.filename"), // 日志文件的位置
		MaxSize:    conf.Viper.GetInt("log.max_size"),    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: conf.Viper.GetInt("log.max_backups"), // 保留旧文件的最大个数
		MaxAge:     conf.Viper.GetInt("log.max_age"),     // 保留旧文件的最大天数
		Compress:   conf.Viper.GetBool("log.compress"),   // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func GetLogger() *Logger {
	once.Do(func() {
		log = &Logger{}
		writeSyncer := getLogWriter()
		encoder := getEncoder()
		core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
		log.Log = zap.New(core, zap.AddCaller()).Sugar()
		defer func() {
			_ = log.Log.Sync()
		}()
	})

	return log
}
