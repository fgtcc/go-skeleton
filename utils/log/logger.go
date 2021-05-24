package log

import (
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func InitLogger() {
	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	// logger.SetReportCaller(true)

	logger.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logPathPrefix := "log/log"
	logPath := logPathPrefix + "%Y%m%d.log"
	logWriter, _ := rotatelogs.New(
		logPath,
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfsHook := lfshook.NewHook(writeMap, &nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		NoColors:        true,
	})

	logger.AddHook(lfsHook)
	// logger.AddHook(hook.NewCallerHook())
}

func GetLogger() *logrus.Logger {
	return logger
}

// DecorateRuntimeContext appends line, file and function context to the logger
// func DecorateRuntimeContext(logger *logrus.Logger) *logrus.Entry {
// 	if pc, file, line, ok := runtime.Caller(2); ok {
// 		fName := runtime.FuncForPC(pc).Name()
// 		return logger.WithField("file", file).WithField("line", line).WithField("func", fName)
// 	} else {
// 		return logrus.NewEntry(logger)
// 	}
// }

// logger methods
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	logger.Fatalf(template, args...)
}

func Panic(args ...interface{}) {
	logger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logger.Panicf(template, args...)
}
