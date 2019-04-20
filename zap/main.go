package main

import "fmt"
import "time"
import "go.uber.org/zap"
import "go.uber.org/zap/zapcore"

func main() {
	c1 := &Logger{
		lv: zapcore.DebugLevel,
	}
	c2 := &Logger{
		lv: zapcore.WarnLevel,
	}
	
	l := zap.New(zapcore.NewTee(c1, c2))

	l.Debug("Hi")
	l.Info("Hi")
	l.Warn("Hi")
	l.Error("Hi")
}

type Logger struct {
	lv zapcore.Level
	fields []zapcore.Field
}

func (l *Logger) Enabled(lv zapcore.Level) bool {
	return l.lv.Enabled(lv)
}

func (l *Logger) With(fields []zapcore.Field) zapcore.Core {
	if len(fields) <= 0 {
		return l
	}
	
	l.fields = append(l.fields, fields...)
	return l
}

func (l *Logger) Check(e zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if l.Enabled(e.Level) {
		return ce.AddCore(e, l)
	}
	return ce
}

func (l *Logger) Write(e zapcore.Entry, fields []zapcore.Field) error {
	_, err := fmt.Println(fmt.Sprintf("[%s][%s] %s", e.Time.Format(time.RFC3339), e.Level.String(), e.Message))
	return err
}

func (l *Logger) Sync() error {
	return nil
}
