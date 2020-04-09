/*

logger.go

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

package logger

import (
	"fmt"
	"time"
)

type Logger interface {
	setLevel(level LogLevel)
	getLevel() LogLevel

	// Ataches the logger to a new logging sink
	attach(log Log)

	error(format string, v ...interface{}) (int, error)
	warn(format string, v ...interface{}) (int, error)
	info(format string, v ...interface{}) (int, error)
	debug(format string, v ...interface{}) (int, error)
}

type loggerImpl struct {
	log     Log
	level   LogLevel
	levels  map[string]LogLevel
	section string
}

func newLogger(log Log, section string) Logger {
	return &loggerImpl{
		log:     log,
		section: section,
	}
}

func (logger *loggerImpl) setLevel(level LogLevel) {
	logger.level = level
}

func SetLevel(logger Logger, level LogLevel) {
	if logger == nil {
		return
	}

	logger.setLevel(level)
}

func (logger *loggerImpl) getLevel() LogLevel {
	return logger.level
}

func GetLevel(logger Logger) LogLevel {
	if logger == nil {
		return LevelError
	}

	return logger.getLevel()
}

func (logger *loggerImpl) attach(log Log) {
	logger.log = log
}

func Attach(logger Logger, log Log) {
	if logger == nil {
		return
	}

	logger.attach(log)
}

func (logger *loggerImpl) write(level string, format string, v ...interface{}) (int, error) {
	t := time.Now()
	s := fmt.Sprintf(format, v...)
	s2 := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d-%03d:%03d %s %s\t%s\n",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(),
		t.Second(), t.Nanosecond()/1000000, (t.Nanosecond()/1000)%1000,
		level, logger.section, s)
	return logger.log.Write(s2)
}

func (logger *loggerImpl) error(format string, v ...interface{}) (int, error) {
	if logger.level < LevelError {
		return 0, nil
	}

	return logger.write("ERR", format, v...)
}

func Error(logger Logger, format string, v ...interface{}) (int, error) {
	if logger == nil {
		return 0, nil
	}

	return logger.error(format, v ...)
}

func (logger *loggerImpl) warn(format string, v ...interface{}) (int, error) {
	if logger.level < LevelWarn {
		return 0, nil
	}

	return logger.write("WARN", format, v...)
}

func Warn(logger Logger, format string, v ...interface{}) (int, error) {
	if logger == nil {
		return 0, nil
	}

	return logger.warn(format, v ...)
}

func (logger *loggerImpl) info(format string, v ...interface{}) (int, error) {
	if logger.level < LevelInfo {
		return 0, nil
	}

	return logger.write("INFO", format, v...)
}

func Info(logger Logger, format string, v ...interface{}) (int, error) {
	if logger == nil {
		return 0, nil
	}

	return logger.info(format, v ...)
}

func (logger *loggerImpl) debug(format string, v ...interface{}) (int, error) {
	if logger.level < LevelDebug {
		return 0, nil
	}

	return logger.write("DBG", format, v...)
}

func Debug(logger Logger, format string, v ...interface{}) (int, error) {
	if logger == nil {
		return 0, nil
	}

	return logger.debug(format, v ...)
}

