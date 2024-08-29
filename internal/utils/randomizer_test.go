//go:build unit

package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"homecomp/internal/utils"
)

func TestRandomizer(t *testing.T) {
	t.Run("Test randomizer", func(t *testing.T) {
		res := utils.RandomStrOfLen(5)
		assert.Len(t, res, 5)

		res = utils.RandomStrOfLen(10)
		assert.Len(t, res, 10)

		res = utils.RandomStrOfLen(25)
		assert.Len(t, res, 25)
	})
}
