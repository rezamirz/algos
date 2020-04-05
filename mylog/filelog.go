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
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type FileLog struct {
	filename     string
	base         string		// base = basename + '.' + baseExt
	basename     string
	baseExt      string
	dir          string
	file         *os.File
	mutex        *sync.RWMutex
	loggers      map[string]Logger
	total        int64
	logSize      int64
	nRotation    int
	nextRotation int
}

/* Creates a FileLog with specified filename and determined log level */
func (flog *FileLog) Open() error {
	file, err := flog.doOpen()
	if err != nil {
		return err
	}

	flog.nextRotation = flog.findMaxRotationNumber()
	if flog.nextRotation > flog.nRotation {
		flog.nextRotation = 1
	}
	fmt.Printf("NextRotationNum=%d\n", flog.nextRotation)

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

func (flog *FileLog) findMaxRotationNumber() int {
	var n int

	s := strings.Split(flog.base, ".")
	flog.basename = s[0]
	flog.baseExt = ""
	if len(s) > 1 {
		flog.baseExt = s[1]
	}

	filepath.Walk(flog.dir, func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Walk path=%s, base=%s\n", path, flog.basename)
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

	return n+1
}

func (flog *FileLog) Close() error {
	return flog.file.Close()
}

func (flog *FileLog) Write(msg string) error {
	flog.mutex.Lock()
	n, err := flog.file.Write([]byte(msg))
	flog.total += int64(n)
	if flog.total >= flog.logSize {
		flog.Rotate()
	}
	flog.mutex.Unlock()
	return err
}

func (flog *FileLog) SetRotation(logSize int64, nRotation int) {
	flog.logSize = logSize
	flog.nRotation = nRotation
}

func (flog *FileLog) Rotate() error {
	flog.mutex.Lock()
	defer flog.mutex.Unlock()

	flog.file.Close()

	newFilename := flog.basename
	os.Rename(flog.filename, newFilename)

	file, err := flog.doOpen()
	if err != nil {
		return err
	}

	flog.file = file
	return nil
}

func (flog *FileLog) GetLogger(section string) Logger {
	flog.mutex.Lock()
	defer flog.mutex.Unlock()

	logger, ok := flog.loggers[section]
	if ok {
		return logger
	}

	logger = newLogger(flog, section)
	flog.loggers[section] = logger
	return logger
}
