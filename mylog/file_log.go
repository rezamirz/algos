/*

file_log.go

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
	"github.com/rezamirz/myalgos/configurator"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type FileLog struct {
	filename     string // filename = dir + '/' + base
	base         string // base = basename + '.' + baseExt
	basename     string
	baseExt      string
	dir          string
	file         *os.File
	mutex        *sync.RWMutex
	loggers      map[string]Logger
	total        int64  // Total number of bytes written to the log
	logSize      int64  // Size of the log when rotation happens
	nRotation    int    // Number of rotation files
	nextRotation int    // Next rotation number
	configLevels string // All configuration levels obtained from config file
	defaultLevel LogLevel
}

func newFileLog(configurator configurator.Configurator) (*FileLog, error) {
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

	configLevels, ok := configurator.Get(LEVEL)

	return &FileLog{
		filename:     filename,
		base:         base,
		dir:          dir,
		mutex:        &sync.RWMutex{},
		loggers:      map[string]Logger{},
		logSize:      logSize,
		nRotation:    logRotation,
		configLevels: configLevels,
	}, nil
}

/* Creates a FileLog with specified filename and determined log level */
func (flog *FileLog) Open() error {
	flog.mutex.Lock()
	defer flog.mutex.Unlock()

	file, err := flog.doOpen()
	if err != nil {
		return err
	}

	flog.nextRotation = flog.findMaxRotationNumber()
	if flog.nextRotation > flog.nRotation {
		flog.nextRotation = 1
	}

	fileInfo, err := file.Stat()
	if err == nil {
		flog.total += fileInfo.Size()
	}

	// Set the default log level to LevelError.
	// This might be overwritten by setLevels(), which enforces users configured log levels.
	flog.defaultLevel = LevelError
	err = flog.setLevels()
	if err != nil {
		return err
	}
	fmt.Printf("NextRotationNum=%d, size=%d, total=%d\n", flog.nextRotation, fileInfo.Size(), flog.total)

	flog.file = file
	return nil
}

func (flog *FileLog) doOpen() (*os.File, error) {
	file, err := os.OpenFile(flog.filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}

	return file, err
}

func (flog *FileLog) setLevels() error {
	if len(flog.configLevels) == 0 {
		return nil
	}

	cfgLevels := strings.Split(flog.configLevels, ",")
	for _, cfgLevel := range cfgLevels {
		s := strings.Split(cfgLevel, ":")
		if len(s) != 2 {
			continue
		}
		section := s[0]
		levelStr := s[1]
		level, err := GetLevelFromString(levelStr)
		if err != nil {
			return err
		}

		if strings.Compare(section, "ALL") == 0 {
			flog.defaultLevel = level
			continue
		}

		logger := flog.GetLogger(section)
		logger.SetLevel(level)
	}

	return nil
}

func (flog *FileLog) findMaxRotationNumber() int {
	var n int

	s := strings.Split(flog.base, ".")
	flog.basename = s[0]
	flog.baseExt = ""
	if len(s) > 1 {
		flog.baseExt = s[1]
	}

	filepath.Walk(flog.dir, func(path string, info os.FileInfo, err error) error {
		//fmt.Printf("Walk path=%s, base=%s\n", path, flog.basename)
		if len(path) < len(flog.base) {
			return nil
		}

		if strings.Compare(path[0:len(flog.basename)], flog.basename) != 0 {
			return nil
		}

		p := strings.Split(path, ".")
		if len(p) == 1 || len(p[0]) < len(flog.basename) {
			return nil
		}

		if strings.Compare(flog.baseExt, p[1]) != 0 {
			return nil
		}

		i, err := strconv.Atoi(p[0][len(flog.basename):])
		if err != nil {
			return nil
		}

		if i > n {
			n = i
		}

		return nil
	})

	return n + 1
}

func (flog *FileLog) Close() error {
	flog.mutex.Lock()
	defer flog.mutex.Unlock()

	return flog.file.Close()
}

func (flog *FileLog) Write(msg string) (int, error) {
	flog.mutex.Lock()
	defer flog.mutex.Unlock()

	n, err := flog.file.Write([]byte(msg))
	if err != nil {
		return n, err
	}
	flog.total += int64(n)

	//fmt.Printf("XXX total=%d, logSize=%d\n", flog.total, flog.logSize)

	if flog.total >= flog.logSize {
		_, err = flog.rotate()
		if err != nil {
			return n, err
		}
	}

	return n, err
}

func (flog *FileLog) SetRotation(logSize int64, nRotation int) {
	flog.logSize = logSize
	flog.nRotation = nRotation
}

func (flog *FileLog) GetRotation() int {
	return flog.nextRotation
}

func (flog *FileLog) Rotate() (interface{}, error) {
	flog.mutex.Lock()
	defer flog.mutex.Unlock()

	return flog.rotate()
}

func (flog *FileLog) rotate() (interface{}, error) {
	flog.file.Close()

	var newFilename string
	if flog.nextRotation > 0 {
		newFilename = fmt.Sprintf("%s/%s%d.%s", flog.dir, flog.basename, flog.nextRotation, flog.baseExt)
	} else {
		newFilename = fmt.Sprintf("%s/%s.%s", flog.dir, flog.basename, flog.baseExt)
	}
	fmt.Printf("XXX filename=%s, new=%s\n", flog.filename, newFilename)
	os.Rename(flog.filename, newFilename)

	file, err := flog.doOpen()
	if err != nil {
		return nil, err
	}

	flog.file = file
	flog.total = 0
	flog.nextRotation = flog.nextRotation%flog.nRotation + 1
	return newFilename, nil
}

func (flog *FileLog) GetLogger(section string) Logger {
	flog.mutex.Lock()
	defer flog.mutex.Unlock()

	logger, ok := flog.loggers[section]
	if ok {
		return logger
	}

	logger = newLogger(flog, section)
	logger.SetLevel(flog.defaultLevel)
	flog.loggers[section] = logger
	return logger
}
