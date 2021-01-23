package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var ZapLogger *zap.Logger
var SugarLogger *zap.SugaredLogger

func InitLogger() {
	hook := getLogWriter()
	logWriter := []zapcore.WriteSyncer{}
	logWriter = append(logWriter, hook)

	encoder := getEncoder()
	core := zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), hook),
		zapcore.DebugLevel)

	ZapLogger = zap.New(core, zap.AddCaller()) //印出log的位置
	// ZapLogger.Debug("POK")                     // ZapLogger sample
	// ZapLogger.Info("failed to fetch URL",
	// 	zap.String("url", "uedf"),
	// 	zap.Int("attempt", 3),
	// 	zap.Duration("backoff", time.Second),
	// )
	SugarLogger = ZapLogger.Sugar()
	//SugarLogger.Infof("Success! statusCode = %s for URL %s", "OK", "OK")  // SugarLogger sample

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
