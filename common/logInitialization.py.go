package common

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

//
// LogInitialization
//  @Description: 初始化日志
//  @return *logrus.Logger
//
func LogInitialization() {
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	core := zapcore.NewTee(
		zapcore.NewCore(getFileEncoder(), getLogWriter(), zapcore.ErrorLevel),        // 控制台输出
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel), // 日志文件输出
	)
	logger := zap.New(
		core,
		zap.AddCaller(),                   // 配置调用者的文件名、行号和函数名来注释每条消息
		zap.AddStacktrace(zap.ErrorLevel)) // 配置为记录给定级别或高于给定级别的所有消息的堆栈跟踪

	Logger = logger.Sugar()
}

// getFileEncoder 文件编码器
func getFileEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter 保存文件日志切割
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   Configuration().LogPath,
		MaxSize:    Configuration().LogMaxSize,
		MaxBackups: Configuration().LogMaxBackups,
		MaxAge:     Configuration().LogMaxAge,
		Compress:   true, // 文件是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}
