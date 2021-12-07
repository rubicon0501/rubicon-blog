package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	default: //默认是Debug
		return "debug"
	}
	return ""
}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}
func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}

	for k, v := range f {
		ll.fields[k] = v
	}

	return ll
}

func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s:%d %s", file, line, f.Name())}
	}

	return ll
}

func (l *Logger) WithCallersFrames() *Logger {
	minCallerDepth := 1
	maxCallerDepth := 25
	callers := make([]string, 0)
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}

func (l *Logger) JSONFormat(level Level, msg string) Fields {
	res := make(Fields, len(l.fields)+4)
	res["level"] = level.String()
	res["time"] = time.Now().Local().UnixNano()
	res["msg"] = msg
	res["callers"] = l.callers

	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := res[k]; !ok {
				res[k] = v
			}
		}
	}

	return res
}

func (l *Logger) OutputFormat(level Level, msg string) {
	jsonMap := l.JSONFormat(level, msg)
	content := fmt.Sprintf("[%s] [%s] [%v] %s",
		jsonMap["time"], jsonMap["level"], jsonMap["callers"], jsonMap["msg"])

	switch level {
	case LevelDebug, LevelInfo, LevelWarn, LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	default:
		l.newLogger.Printf(content)
	}
}

func (l *Logger) OutputJSON(level Level, msg string) {
	jsonMap := l.JSONFormat(level, msg)
	body, _ := json.Marshal(jsonMap)
	content := string(body)

	switch level {
	case LevelDebug, LevelInfo, LevelWarn, LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	default:
		l.newLogger.Printf(content)
	}
}

func (l *Logger) Info(v ...interface{}) {
	l.OutputFormat(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.OutputFormat(LevelInfo, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	if logger != nil {
		logger.Info(v...)
	} else {
		panic("Invalid logger pointer")
	}
}
func Infof(format string, v ...interface{}) {
	if logger != nil {
		logger.Infof(format, v...)
	} else {
		panic("Invalid logger pointer")
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.OutputFormat(LevelDebug, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.OutputFormat(LevelDebug, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	if logger != nil {
		logger.Debug(v...)
	} else {
		panic("Invalid logger pointer")
	}
}
func Debugf(format string, v ...interface{}) {
	if logger != nil {
		logger.Debugf(format, v...)
	} else {
		panic("Invalid logger pointer")
	}
}

func (l *Logger) Warn(v ...interface{}) {
	l.OutputFormat(LevelWarn, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.OutputFormat(LevelWarn, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	if logger != nil {
		logger.Warn(v...)
	} else {
		panic("Invalid logger pointer")
	}
}
func Warnf(format string, v ...interface{}) {
	if logger != nil {
		logger.Warnf(format, v...)
	} else {
		panic("Invalid logger pointer")
	}
}

func (l *Logger) Error(v ...interface{}) {
	l.OutputFormat(LevelError, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.OutputFormat(LevelError, fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	if logger != nil {
		logger.Error(v...)
	} else {
		panic("Invalid logger pointer")
	}
}
func Errorf(format string, v ...interface{}) {
	if logger != nil {
		logger.Errorf(format, v...)
	} else {
		panic("Invalid logger pointer")
	}
}

func (l *Logger) Fatal(v ...interface{}) {
	l.OutputFormat(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.OutputFormat(LevelFatal, fmt.Sprintf(format, v...))
}

func Fatal(v ...interface{}) {
	if logger != nil {
		logger.Fatal(v...)
	} else {
		panic("Invalid logger pointer")
	}
}
func Fatalf(format string, v ...interface{}) {
	if logger != nil {
		logger.Fatalf(format, v...)
	} else {
		panic("Invalid logger pointer")
	}
}

func (l *Logger) Panic(v ...interface{}) {
	l.OutputFormat(LevelPanic, fmt.Sprint(v...))
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.OutputFormat(LevelPanic, fmt.Sprintf(format, v...))
}

func Panic(v ...interface{}) {
	if logger != nil {
		logger.Panic(v...)
	} else {
		panic("Invalid logger pointer")
	}
}
func Panicf(format string, v ...interface{}) {
	if logger != nil {
		logger.Panicf(format, v...)
	} else {
		panic("Invalid logger pointer")
	}
}
