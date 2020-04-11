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

package logger

import (
	"github.com/rezamirz/myalgos/configurator"
	"os"
	"strconv"
	"strings"
	"sync"
)

type MemLog struct {
	filename     string // If it exist, it will append data to it during rotation
	file         *os.File
	mutex        *sync.RWMutex
	loggers      map[string]Logger
	buf          []byte
	total        int64 // Total number of bytes written to the log
	logSize      int64
	configLevels string // All configuration levels obtained from config file
	defaultLevel LogLevel
}

func newMemLog(configurator configurator.Configurator) (*MemLog, error) {
	filename, ok := configurator.Get(FILENAME)
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

	configLevels, ok := configurator.Get(LEVEL)

	return &MemLog{
		filename:     filename,
		mutex:        &sync.RWMutex{},
		loggers:      map[string]Logger{},
		logSize:      logSize,
		configLevels: configLevels,
	}, nil
}

func (mem *MemLog) Open() error {
	mem.mutex.Lock()
	defer mem.mutex.Unlock()

	var file *os.File
	var err error
	if len(mem.filename) > 0 {
		file, err = os.OpenFile(mem.filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			return err
		}
	}
	mem.file = file
	mem.defaultLevel = LevelError
	err = mem.setLevels()
	if err != nil {
		return err
	}
	mem.buf = make([]byte, mem.logSize)
	return nil
}

func (mem *MemLog) setLevels() error {
	if len(mem.configLevels) == 0 {
		return nil
	}

	cfgLevels := strings.Split(mem.configLevels, ",")
	for _, cfgLevel := range cfgLevels {
		s := strings.Split(cfgLevel, ":")
		if len(s) != 2 {
			continue
		}
		section := strings.Trim(s[0], " ")
		levelStr := strings.Trim(s[1], " ")
		level, err := GetLevelFromString(levelStr)
		if err != nil {
			return err
		}

		if strings.Compare(section, "ALL") == 0 {
			mem.defaultLevel = level
			continue
		}

		logger := mem.getLogger(section)
		SetLevel(logger, level)
	}

	return nil
}

func (mem *MemLog) Close() error {
	mem.mutex.Lock()
	defer mem.mutex.Unlock()

	if len(mem.filename) > 0 {
		mem.file.Write(mem.buf)
	}

	return nil
}

func (mem *MemLog) Write(msg string) (int, error) {
	mem.mutex.Lock()
	defer mem.mutex.Unlock()

	if mem.total+int64(len(msg)) > mem.logSize {
		mem.Rotate()
	}
	copy(mem.buf[mem.total:], []byte(msg))
	mem.total += int64(len(msg))
	return len(msg), nil
}

func (mem *MemLog) SetRotation(logSize int64, nRotation int) {
	mem.logSize = logSize
}

func (mem *MemLog) GetRotation() int {
	return 0
}

func (mem *MemLog) Rotate() (interface{}, error) {
	oldBuf := mem.buf
	total := mem.total
	mem.total = 0
	mem.buf = make([]byte, mem.logSize)

	if len(mem.filename) > 0 {
		go func(buf []byte) {
			mem.file.Write(buf)
		}(oldBuf[:total])
	}

	return oldBuf[:total], nil
}

func (mem *MemLog) GetLogger(section string) Logger {
	mem.mutex.Lock()
	defer mem.mutex.Unlock()
	return mem.getLogger(section)
}

func (mem *MemLog) getLogger(section string) Logger {
	logger, ok := mem.loggers[section]
	if ok {
		return logger
	}

	logger = newLogger(mem, section)
	mem.loggers[section] = logger
	return logger
}
