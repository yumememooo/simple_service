package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLogger() {
	hook := getLogWriter()
	logWriter := []zapcore.WriteSyncer{}
	logWriter = append(logWriter, hook)

	encoder := getEncoder()
	core := zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), hook),
		zapcore.DebugLevel)

	Logger = zap.New(core, zap.AddCaller()) //印出log的位置
	Logger.Debug("POK")
	Logger.Info("failed to fetch URL",
		zap.String("url", "uedf"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	sugarLogger := Logger.Sugar()
	sugarLogger.Infof("Success! statusCode = %s for URL %s", "OK", "OK")

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
