/*

stdout_log.go

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
	"github.com/rezamirz/myalgos/configurator"
	"os"
	"strconv"
	"sync"
)

type StdOutLog struct {
	mutex   *sync.RWMutex
	loggers map[string]Logger
	file    *os.File
	total   int64 // Total number of bytes written to the log
	logSize int64
}

func newStdOutLog(configurator configurator.Configurator) (*StdOutLog, error) {
	logSizeStr, ok := configurator.Get(LOGFILE_SIZE)
	var logSize int64
	var err error
	if ok {
		logSize, err = strconv.ParseInt(logSizeStr, 10, 64)
		if err != nil {
			return nil, ErrInvalidLogSize
		}
	} else {
		logSize = DefaultLogSize
	}
	return &StdOutLog{
		mutex:   &sync.RWMutex{},
		loggers: map[string]Logger{},
		logSize: logSize,
	}, nil
}

func (std *StdOutLog) Open() error {
	std.mutex.Lock()
	defer std.mutex.Unlock()

	std.file = os.Stdout
	return nil
}

func (std *StdOutLog) Close() error {
	std.mutex.Lock()
	defer std.mutex.Unlock()

	std.file = nil
	return nil
}

func (std *StdOutLog) Write(msg string) (int, error) {
	std.mutex.Lock()
	defer std.mutex.Unlock()

	n, err := std.file.Write([]byte(msg))
	std.total += int64(n)
	if std.total >= std.logSize {
		std.Rotate()
	}
	return n, err
}

func (std *StdOutLog) SetRotation(logSize int64, nRotation int) {
	std.logSize = logSize
}

func (std *StdOutLog) GetRotation() int {
	return 0
}

func (std *StdOutLog) Rotate() (interface{}, error) {
	std.file.Sync()
	std.total = 0
	return nil, nil
}

func (std *StdOutLog) GetLogger(section string) Logger {
	std.mutex.Lock()
	defer std.mutex.Unlock()

	logger, ok := std.loggers[section]
	if ok {
		return logger
	}

	logger = newLogger(std, section)
	std.loggers[section] = logger
	return logger
}
