
package mylog

import (
	"github.com/rezamirz/myalgos/configurator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinPQ(t *testing.T) {
	configurator := configurator.NewConfigurator()
	configurator.Put(LOGTYPE, "file")
	configurator.Put(FILENAME, "mylog.log")
	configurator.Put(LOGFILE_SIZE, "1000")
	configurator.Put(LOG_ROTATION, "10")
	log, err := New(configurator)
	assert.NoError(t, err)

	err = log.Open()
	assert.NoError(t, err)

	logger := log.GetLogger("test1")
	logger.SetLevel(LevelInfo)
	for i:=0; i<100; i++ {
		logger.Info("Loop %d", i)
	}

	log.Close()
}