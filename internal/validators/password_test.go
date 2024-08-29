package validators_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"homecomp/internal/validators"
)

func TestPasswordValidator(t *testing.T) {
	t.Run("it must not validate invalid passwords", func(t *testing.T) {
		invalidPasswd := []string{
			"Test123",
			"Test.test",
			"testonga",
			"TESTONGA",
		}

		for _, pass := range invalidPasswd {
			assert.False(t, validators.IsPasswordCharsValid(pass), fmt.Sprintf("%s should be false", pass))
		}
	})

	t.Run("it must validate valid passwords", func(t *testing.T) {
		invalidPasswd := []string{
			"Test.123!",
			"Testonga!24",
			"G4r0mp3t4!",
			"abc123!@#ABC",
		}

		for _, pass := range invalidPasswd {
			assert.True(t, validators.IsPasswordCharsValid(pass), fmt.Sprintf("%s should be true", pass))
		}
	})

}
