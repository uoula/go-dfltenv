package dfltenv

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetenv(t *testing.T) {
	os.Setenv("TEST_KEY", "SET VALUES")
	defer os.Unsetenv("TEST_KEY")
	assert.Equal(t, "SET VALUES", Getenv("TEST_KEY", "DEFAULT VALUE"))
}

func TestSetEmptyStringEnv(t *testing.T) {
	os.Setenv("TEST_KEY", "")
	defer os.Unsetenv("TEST_KEY")
	assert.Equal(t, "", Getenv("TEST_KEY", "DEFAULT VALUE"))
}

func TestNotSetEnv(t *testing.T) {
	assert.Equal(t, "DEFAULT VALUE", Getenv("TEST_KEY", "DEFAULT VALUE"))
}
