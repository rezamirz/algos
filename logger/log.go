/*

log.go

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
	"errors"
	"github.com/rezamirz/myalgos/configurator"
	"strings"
)

var ErrNoLogInConfigurator = errors.New("No log in configurator")
var ErrNoFilenameInConfigurator = errors.New("No filename in configurator")
var ErrNoLevelInConfigurator = errors.New("No level in configurator")
var ErrInvalidLogLevel = errors.New("Invalid log level in configurator")
var ErrInvalidLogType = errors.New("Invalid log type")
var ErrInvalidLogSize = errors.New("Invalid log size")
var ErrInvalidLogRotation = errors.New("Invalid log rotation")

const (
	LOGTYPE      = "logtype"
	FILE_LOG     = "file_log"
	STDOUT_LOG   = "stdout_log"
	MEM_LOG      = "mem_log"
	FILENAME     = "filename"
	LOGFILE_SIZE = "log_size"
	LOG_ROTATION = "log_rotation"
	LEVEL        = "level"

	DefaultLogSize     = 100 * 1024 * 1024
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
	// Opens the log sink (file/mem/device)
	Open() error

	// Closes the log sink
	Close() error

	// Returns the number of bytes that are written to the log and error if there is any
	Write(msg string) (int, error)

	// Sets the rotation size and number of log rotations
	SetRotation(rotationSize int64, nRotation int)

	// Returns the next rotation number
	GetRotation() int

	// Write() calls Rotate() as soon as the log_file_size >= rotation_size
	// Rotate() returns a reference to the rotated log file.
	// If the log file if FILE_LOG, it returns the filename that the log is rotated to.
	// If the log file is MEM_LOG, it returns the memory that contains all the log data.
	Rotate() (interface{}, error)

	// Returns a section logger that associated with this log
	GetLogger(section string) Logger
}

// Opens a logger with log method Log
func New(configurator configurator.Configurator) (Log, error) {
	logtype, ok := configurator.Get(LOGTYPE)
	if !ok {
		return nil, ErrNoLogInConfigurator
	}

	switch logtype {
	case FILE_LOG:
		return newFileLog(configurator)
	case STDOUT_LOG:
		return newStdOutLog(configurator)
	case MEM_LOG:
		return newMemLog(configurator)
	}

	return nil, ErrInvalidLogType
}

func GetLevelFromString(levelStr string) (LogLevel, error) {
	levelStr = strings.ToUpper(levelStr)
	if strings.Compare(levelStr, "INFO") == 0 {
		return LevelInfo, nil
	}

	if strings.Compare(levelStr, "DEBUG") == 0 || strings.Compare(levelStr, "DBG") == 0 {
		return LevelDebug, nil
	}

	if strings.Compare(levelStr, "ERROR") == 0 || strings.Compare(levelStr, "ERR") == 0 {
		return LevelError, nil
	}

	if strings.Compare(levelStr, "WARNING") == 0 || strings.Compare(levelStr, "WARN") == 0 {
		return LevelWarn, nil
	}

	if strings.Compare(levelStr, "FATAL") == 0 {
		return LevelFatal, nil
	}

	return LevelFatal, ErrInvalidLogLevel
}
