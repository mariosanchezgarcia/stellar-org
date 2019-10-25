package env_test

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"testing"
	"time"

	"github.com/stellar/go/support/env"
	"github.com/stellar/go/support/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func randomStr(length int) string {
	raw := make([]byte, (length+1)/2)
	_, err := rand.Read(raw)
	if err != nil {
		err = errors.Wrap(err, "read from crypto/rand failed")
		panic(err)
	}
	return hex.EncodeToString(raw)[:length]
}

// TestString_set tests that env.String will return the value of the
// environment variable when the environment variable is set.
func TestString_set(t *testing.T) {
	envVar := "TestString_set_" + randomStr(10)
	err := os.Setenv(envVar, "value")
	require.NoError(t, err)
	defer os.Unsetenv(envVar)

	value := env.String(envVar, "default")
	assert.Equal(t, "value", value)
}

// TestString_set tests that env.String will return the default value given
// when the environment variable is not set.
func TestString_notSet(t *testing.T) {
	envVar := "TestString_notSet_" + randomStr(10)
	value := env.String(envVar, "default")
	assert.Equal(t, "default", value)
}

// TestInt_set tests that env.Int will return the value of the environment
// variable as an int when the environment variable is set.
func TestInt_set(t *testing.T) {
	envVar := "TestInt_set_" + randomStr(10)
	err := os.Setenv(envVar, "12345")
	require.NoError(t, err)
	defer os.Unsetenv(envVar)

	value := env.Int(envVar, 67890)
	assert.Equal(t, 12345, value)
}

// TestInt_set tests that env.Int will return the default value given when the
// environment variable is not set.
func TestInt_notSet(t *testing.T) {
	envVar := "TestInt_notSet_" + randomStr(10)
	value := env.Int(envVar, 67890)
	assert.Equal(t, 67890, value)
}

// TestDuration_set tests that env.Duration will return the value of the
// environment variable as a time.Duration when the environment variable is
// set to a duration string.
func TestDuration_set(t *testing.T) {
	envVar := "TestDuration_set_" + randomStr(10)
	err := os.Setenv(envVar, "5m30s")
	require.NoError(t, err)
	defer os.Unsetenv(envVar)

	setValue := time.Duration(330000000000)
	defaultValue := 2 * time.Minute
	value := env.Duration(envVar, defaultValue)
	assert.Equal(t, setValue, value)
}

// TestDuration_set tests that env.Duration will return the default value given
// when the environment variable is not set.
func TestDuration_notSet(t *testing.T) {
	envVar := "TestDuration_notSet_" + randomStr(10)
	defaultValue := 5*time.Minute + 30*time.Second
	value := env.Duration(envVar, defaultValue)
	assert.Equal(t, defaultValue, value)
}
