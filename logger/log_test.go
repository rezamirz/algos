package logger

import (
	"github.com/rezamirz/myalgos/configurator"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func cleanLogs() {
	files, err := filepath.Glob("mylog*.log")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
}

func TestSimpleLogLevel(t *testing.T) {
	cleanLogs()

	configurator := configurator.NewConfigurator()
	configurator.Put(LOGTYPE, FILE_LOG)
	configurator.Put(FILENAME, "mylog.log")
	configurator.Put(LOGFILE_SIZE, "1000")
	configurator.Put(LOG_ROTATION, "10")
	log, err := New(configurator)
	assert.NoError(t, err)

	err = log.Open()
	assert.NoError(t, err)
	assert.Equal(t, 1, log.GetRotation())

	tlogger := log.GetLogger("test1")
	SetLevel(tlogger, LevelInfo)
	for i := 0; i < 10; i++ {
		n, err := Debug(tlogger,"Loop %d", i)
		assert.NoError(t, err)
		assert.Equal(t, 0, n)

		n, err = Error(tlogger,"Loop %d", i)
		assert.NoError(t, err)
		assert.Equal(t, true, n > 0)
	}

	SetLevel(tlogger, LevelError)
	for i := 0; i < 10; i++ {
		n, err := Debug(tlogger, "Loop %d", i)
		assert.NoError(t, err)
		assert.Equal(t, 0, n)

		n, err = Info(tlogger, "Loop %d", i)
		assert.NoError(t, err)
		assert.Equal(t, 0, n)

	}

	log.Close()
}

func TestLogLevelByConf(t *testing.T) {
	cleanLogs()

	configurator := configurator.NewConfigurator()
	configurator.Put(LOGTYPE, FILE_LOG)
	configurator.Put(FILENAME, "mylog.log")
	configurator.Put(LOGFILE_SIZE, "1000")
	configurator.Put(LOG_ROTATION, "10")

	logConfLevels := "ALL:INFO, test1:ERR, test2:DBG"
	configurator.Put(LEVEL, logConfLevels)
	log, err := New(configurator)

	err = log.Open()
	assert.NoError(t, err)
	assert.Equal(t, 1, log.GetRotation())

	// Section test1 has LevelError
	tlogger := log.GetLogger("test1")
	assert.Equal(t, LevelError, GetLevel(tlogger))
	for i := 0; i < 10; i++ {
		n, err := Debug(tlogger, "Loop %d", i)
		assert.NoError(t, err)
		assert.Equal(t, 0, n)

		n, err = Error(tlogger, "Loop %d", i)
		assert.NoError(t, err)
		assert.Equal(t, true, n > 0)
	}

	// Section test2 has LevelDebug
	tlogger = log.GetLogger("test2")
	for i := 0; i < 10; i++ {
		n, err := Debug(tlogger, "Loop %d", i)
		assert.NoError(t, err)
		assert.Equal(t, true, n > 0)

		n, err = Error(tlogger, "Loop %d", i)
		assert.NoError(t, err)
		assert.Equal(t, true, n > 0)
	}

	// Section test 3 has LevelInfo (default)
	tlogger = log.GetLogger("test3")
	assert.Equal(t, LevelInfo, GetLevel(tlogger))
	for i := 0; i < 10; i++ {
		n, err := Debug(tlogger, "Loop %d", i)
		assert.NoError(t, err)
		assert.Equal(t, 0, n)

		n, err = Error(tlogger, "Loop %d", i)
		assert.NoError(t, err)
		assert.Equal(t, true, n > 0)
	}

	assert.NoError(t, err)

}

func TestSimpleLogRotation(t *testing.T) {
	cleanLogs()

	configurator := configurator.NewConfigurator()
	configurator.Put(LOGTYPE, FILE_LOG)
	configurator.Put(FILENAME, "mylog.log")
	configurator.Put(LOGFILE_SIZE, "1000")
	configurator.Put(LOG_ROTATION, "10")
	log, err := New(configurator)
	assert.NoError(t, err)

	err = log.Open()
	assert.NoError(t, err)
	assert.Equal(t, 1, log.GetRotation())

	tlogger := log.GetLogger("test1")
	SetLevel(tlogger, LevelInfo)
	for i := 0; i < 100; i++ {
		Info(tlogger, "Loop %d", i)
	}

	assert.Equal(t, 5, log.GetRotation())

	log.Close()
}

func TestSimpleLogRotationMultithreaded(t *testing.T) {
	cleanLogs()

	configurator := configurator.NewConfigurator()
	configurator.Put(LOGTYPE, FILE_LOG)
	configurator.Put(FILENAME, "mylog.log")
	configurator.Put(LOGFILE_SIZE, "1000")
	configurator.Put(LOG_ROTATION, "10")
	log, err := New(configurator)
	assert.NoError(t, err)

	err = log.Open()
	assert.NoError(t, err)
	assert.Equal(t, 1, log.GetRotation())

	tlogger := log.GetLogger("test1")

	done := make(chan bool)
	go func(tlogger Logger) {
		SetLevel(tlogger, LevelInfo)
		for i := 0; i < 50; i++ {
			Info(tlogger, "Loop %d", i)
		}
		done <- true
	}(tlogger)

	go func(tlogger Logger) {
		SetLevel(tlogger, LevelInfo)
		for i := 50; i < 100; i++ {
			Info(tlogger, "Loop %d", i)
		}
		done <- true
	}(tlogger)

	<-done
	<-done

	assert.Equal(t, 5, log.GetRotation())

	log.Close()
}

func TestFullLogRotation(t *testing.T) {
	cleanLogs()

	configurator := configurator.NewConfigurator()
	configurator.Put(LOGTYPE, FILE_LOG)
	configurator.Put(FILENAME, "mylog.log")
	configurator.Put(LOGFILE_SIZE, "1000")
	configurator.Put(LOG_ROTATION, "10")
	log, err := New(configurator)
	assert.NoError(t, err)

	err = log.Open()
	assert.NoError(t, err)
	assert.Equal(t, 1, log.GetRotation())

	tlogger := log.GetLogger("test1")
	SetLevel(tlogger, LevelInfo)
	for i := 0; i < 200; i++ {
		Info(tlogger, "Loop %d", i)
	}

	assert.Equal(t, 10, log.GetRotation())
	log.Close()

	// Open the log again and test it

	log, err = New(configurator)
	assert.NoError(t, err)

	err = log.Open()
	assert.NoError(t, err)
	assert.Equal(t, 10, log.GetRotation())

	tlogger = log.GetLogger("test1")
	SetLevel(tlogger, LevelInfo)
	assert.Equal(t, 10, log.GetRotation())

	for i := 200; i < 220; i++ {
		Info(tlogger, "Loop %d", i)
	}
	assert.Equal(t, 1, log.GetRotation())

	log.Close()
}

func TestFullLogRotationMultithreaded(t *testing.T) {
	cleanLogs()

	configurator := configurator.NewConfigurator()
	configurator.Put(LOGTYPE, FILE_LOG)
	configurator.Put(FILENAME, "mylog.log")
	configurator.Put(LOGFILE_SIZE, "1000")
	configurator.Put(LOG_ROTATION, "10")
	log, err := New(configurator)
	assert.NoError(t, err)

	err = log.Open()
	assert.NoError(t, err)
	assert.Equal(t, 1, log.GetRotation())

	tlogger := log.GetLogger("test1")
	SetLevel(tlogger, LevelInfo)

	done := make(chan bool)
	go func(tlogger Logger) {
		for i := 0; i < 100; i++ {
			Info(tlogger, "Loop %d", i)
		}
		done <- true
	}(tlogger)

	go func(tlogger Logger) {
		for i := 100; i < 200; i++ {
			Info(tlogger, "Loop %d", i)
		}
		done <- true
	}(tlogger)

	<-done
	<-done

	assert.Equal(t, 10, log.GetRotation())
	log.Close()
}
