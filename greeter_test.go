package go_unit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

// Assert
func TestStartGreeterAssert(t *testing.T) {
	result := StartGreeter("John")

	assert.Equal(t, "Hello John", result, "Result must be `Hello John`")
}

// Require
func TestStartGreeterRequire(t *testing.T) {
	result := StartGreeter("John")

	require.Equal(t, "Hello John", result, "Result must be 'Hello John'")
}

// Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on mac os")
	}

	result := StartGreeter("John")

	require.Equal(t, "Hello John", result, "Result must be 'Hello John'")
}

// Before and After (Test Main only run per package)
func TestMain(m *testing.M) {
	fmt.Println("Before")
	m.Run()
	fmt.Println("After")
}

// Sub Test
func TestSubTest(t *testing.T) {
	t.Run("John", func(t *testing.T) {
		result := StartGreeter("John")

		require.Equal(t, "Hello John", result, "Result must be 'Hello John'")
	})

	t.Run("Doe", func(t *testing.T) {
		result := StartGreeter("Doe")

		require.Equal(t, "Hello Doe", result, "Result must be 'Hello Doe'")
	})
}