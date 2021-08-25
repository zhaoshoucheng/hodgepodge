package zlog

import (
	"io"
	"os"
	"strings"
	"time"
)

var cwd, _         = os.Getwd()

// FormatWriter 格式化接口
type FormatWriter interface {
	FormatCaller(string) string
	FormatTimestamp() string
	FormatMessage(string) string
	FormatLevel(string) string
}

// LevelWriter Writer 接口，定义格式化、按照日志级别进行日志写入，核心Writer
type LevelWriter interface {
	FormatWriter
	io.Writer
	WriteLevel(level Level, p []byte) (n int, err error)
}

//--------以下几种 writer 实现 LevelWriter 接口---------------------

//Writer 默认适配器，非LevelWriter的对象利用levelWriterAdapter进行封装
type levelWriterAdapter struct {
	io.Writer
}

func (lw levelWriterAdapter) WriteLevel(l Level, p []byte) (n int, err error) {
	//finalwrite
	return lw.Write(p)
}

func (lw levelWriterAdapter) FormatCaller(i string) string {
	return defaultFormatCaller(i)
}

func (lw levelWriterAdapter) FormatTimestamp() string {
	return defaultFormatTimestamp()
}

func (lw levelWriterAdapter) FormatMessage(i string) string {
	return defaultFormatMessage(i)
}

func (lw levelWriterAdapter) FormatLevel(i string) string {
	return defaultFormatLevel(i)
}


//LevelWriter数组，相同信息可以同时写入多个LevelWriter 中
type multiLevelWriter struct {
	writers []LevelWriter
}

func (t multiLevelWriter) Write(p []byte) (n int, err error) {
	for _, w := range t.writers {
		n, err = w.Write(p)
		if err != nil {
			return
		}
		if n != len(p) {
			err = io.ErrShortWrite
			return
		}
	}
	return len(p), nil
}

func (t multiLevelWriter) WriteLevel(l Level, p []byte) (n int, err error) {
	for _, w := range t.writers {
		n, err = w.WriteLevel(l, p)
		if err != nil {
			return
		}
		if n != len(p) {
			err = io.ErrShortWrite
			return
		}
	}
	return len(p), nil
}

func (t multiLevelWriter) FormatCaller(i string) string {
	return defaultFormatCaller(i)
}

func (t multiLevelWriter) FormatTimestamp() string {
	return defaultFormatTimestamp()
}

func (t multiLevelWriter) FormatMessage(i string) string {
	return defaultFormatMessage(i)
}

func (t multiLevelWriter) FormatLevel(i string) string {
	return defaultFormatLevel(i)
}

//注册multiLevelWriter， 不是LevelWriter的对象可以通过封装levelWriterAdapter实现
func MultiLevelWriter(writers ...io.Writer) LevelWriter {
	lwriters := make([]LevelWriter, 0, len(writers))
	for _, w := range writers {
		if lw, ok := w.(LevelWriter); ok {
			lwriters = append(lwriters, lw)
		} else {
			lwriters = append(lwriters, levelWriterAdapter{w})
		}
	}
	return multiLevelWriter{lwriters}
}

// 默认格式化解析器
var (
	defaultFormatTimestamp = func() string {
		t := time.Now().Format("2006-01-02T15:04:05.000-0700")
		t = "[" + t + "]"
		return t
	}

	defaultFormatLevel = func(i string) string {
		var l string

		switch i {
		case "debug":
			l = "[DEBUG]"
		case "info":
			l = "[INFO]"
		case "warn":
			l = "[WARNING]"
		case "error":
			l = "[ERROR]"
		case "fatal":
			l = "[FATAL]"
		case "panic":
			l = "[PANIC]"
		default:
			l = "[???]"
		}
		return l
	}

	defaultFormatCaller = func(i string) string {
		var c string

		if len(i) > 0 {
			if cwd != "" {
				c = strings.TrimPrefix(i, cwd)
				c = strings.TrimPrefix(c, "/")
			}
			c = "[" + c + "]"
		}
		return c
	}

	defaultFormatMessage = func(i string) string {
		return i
	}
)

