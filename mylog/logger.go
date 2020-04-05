/*

mylogger.go

Logs a section of a program.

MIT License

Copyright (c) 2019 rezamirz

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package mylog

import (
	"fmt"
	"time"
)

type Logger interface {
	SetLevel(level LogLevel)
	GetLevel() LogLevel

	// Ataches the logger to a new logging sink
	Attach(log Log)

	Error(format string, v ...interface{}) (int, error)
	Warn(format string, v ...interface{}) (int, error)
	Info(format string, v ...interface{}) (int, error)
	Debug(format string, v ...interface{}) (int, error)
}

type LoggerImpl struct {
	log     Log
	level   LogLevel
	levels  map[string]LogLevel
	section string
}

func newLogger(log Log, section string) Logger {
	return &LoggerImpl{
		log:     log,
		level:   LevelError,
		section: section,
	}
}

func (logger *LoggerImpl) SetLevel(level LogLevel) {
	logger.level = level
}

func (logger *LoggerImpl) GetLevel() LogLevel {
	return logger.level
}

func (logger *LoggerImpl) Attach(log Log) {
	logger.log = log
}

func (logger *LoggerImpl) write(level string, format string, v ...interface{}) (int, error) {
	t := time.Now()
	s := fmt.Sprintf(format, v...)
	s2 := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-%03d:%03d %s %s\t%s\n",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(),
		t.Second(), t.Nanosecond()/1000000, (t.Nanosecond()/1000)%1000,
		level, logger.section, s)
	return logger.log.Write(s2)
}

func (logger *LoggerImpl) Error(format string, v ...interface{}) (int, error) {
	if logger.level < LevelError {
		return 0, nil
	}

	return logger.write("ERR", format, v...)
}

func (logger *LoggerImpl) Warn(format string, v ...interface{}) (int, error) {
	if logger.level < LevelWarn {
		return 0, nil
	}

	return logger.write("WARN", format, v...)
}

func (logger *LoggerImpl) Info(format string, v ...interface{}) (int, error) {
	if logger.level < LevelInfo {
		return 0, nil
	}

	return logger.write("INFO", format, v...)
}

func (logger *LoggerImpl) Debug(format string, v ...interface{}) (int, error) {
	if logger.level < LevelDebug {
		return 0, nil
	}

	return logger.write("DBG", format, v...)
}
