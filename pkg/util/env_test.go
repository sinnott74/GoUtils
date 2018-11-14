package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetEnvOrDefault tests that an environment variable can be read when set
func TestGetEnvOrDefault(t *testing.T) {
	name := "FOO"
	value := "BAR"
	dfault := "BAZ"
	os.Setenv(name, value)

	envVar := GetEnvOrDefault(name, dfault)
	assert.Equal(t, value, envVar)
}

// TestGetEnvOrDefaultNotSet tests that a default will be returned when an environment variable is not set
func TestGetEnvOrDefaultNotSet(t *testing.T) {
	name := "FOO"
	dfault := "BAZ"
	os.Unsetenv(name)

	envVar := GetEnvOrDefault(name, dfault)
	assert.Equal(t, dfault, envVar)
}
