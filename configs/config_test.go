package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	cmdDir, err := filepath.Abs("../cmd/server/")
	assert.NoError(t, err)

	err = os.Chdir(cmdDir)
	assert.NoError(t, err)

	defer func() {
		err := os.Chdir(filepath.Dir(cmdDir))
		assert.NoError(t, err)
	}()

	err = LoadConfig(".")

	assert.NoError(t, err, "Error returned when trying to read .env file.")
}

func TestLoadConfig_Error(t *testing.T) {
	defer func() {
		assert.NotNil(t, recover(), "Success expected in error scenario")
	}()

	LoadConfig(".")
}
