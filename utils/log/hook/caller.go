package hook

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

type CallerHook struct {
}

func (hook *CallerHook) Fire(entry *logrus.Entry) error {
	if pc, file, line, ok := runtime.Caller(9); ok {
		fName := runtime.FuncForPC(pc).Name()

		entry.Data["file"] = file
		entry.Data["line"] = line
		entry.Data["func"] = fName
	}
	return nil
}

func (hook *CallerHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func NewCallerHook() *CallerHook {
	return &CallerHook{}
}
