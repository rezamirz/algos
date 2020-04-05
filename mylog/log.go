/*

mylog.go

MIT License

Copyright (c) 2018 rezamirz

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
	"errors"
	"github.com/rezamirz/myalgos/configurator"
	"path/filepath"
	"strconv"
	"sync"
)

var ErrNoLogInConfigurator = errors.New("No log in configurator")
var ErrNoFilenameInConfigurator = errors.New("No filename in configurator")
var ErrNoLevelInConfigurator = errors.New("No level in configurator")
var ErrInvalidLogLevel = errors.New("Invalid log level in configurator")
var ErrInvalidLogType = errors.New("Invalid log type")
var ErrInvalidLogSize = errors.New("Invalid log size")
var ErrInvalidLogRotation = errors.New("Invalid log rotation")


const (
	LOGTYPE  = "logtype"
	FILENAME = "filename"
	LOGFILE_SIZE = "log_size"
	LOG_ROTATION = "log_rotation"
	LEVEL    = "level"

	DefaultLogSize = 100 * 1024 * 1024
	DefaultLogRotation = 10
)

type LogLevel int

const (
	LevelFatal LogLevel = iota
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

type Log interface {
	Open() error
	Close() error
	Write(msg string) error

	SetRotation(rotationSize int64, nRotation int)

	// Rotate() must be called as soon as the log_file_size >= rotation_size
	Rotate() error

	GetLogger(section string) Logger
}

// Opens a logger with log method Log
func New(configurator configurator.Configurator) (Log, error) {
	logtype, ok := configurator.Get(LOGTYPE)
	if !ok {
		return nil, ErrNoLogInConfigurator
	}

	switch logtype {
	case "file":
		filename, ok := configurator.Get(FILENAME)
		if !ok {
			return nil, ErrNoFilenameInConfigurator
		}
		base := filepath.Base(filename)
		dir := filepath.Dir(filename)

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

		logRotationStr, ok := configurator.Get(LOG_ROTATION)
		var logRotation int
		if ok {
			logRotation, err = strconv.Atoi(logRotationStr)
			if err != nil {
				return nil, ErrInvalidLogRotation
			}
		} else {
			logRotation = DefaultLogRotation
		}

		return &FileLog{
			filename: filename,
			base: base,
			dir: dir,
			mutex:    &sync.RWMutex{},
			loggers:  map[string]Logger{},
			logSize: logSize,
			nRotation: logRotation,
		}, nil
	}

	return nil, ErrInvalidLogType
}
