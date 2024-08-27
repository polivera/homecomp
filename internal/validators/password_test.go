package validators_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"homecomp/internal/validators"
)

func TestPasswordValidator(t *testing.T) {

	t.Run("", func(t *testing.T) {
		invalidPasswd := []string{
			"Test123",
			"Test.test",
			"testonga",
			"TESTONGA",
		}

		for _, pass := range invalidPasswd {
			assert.False(t, validators.IsPasswordCharsValid(pass))
		}
	})

}
