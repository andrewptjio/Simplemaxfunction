package max

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionMax(t *testing.T) {
	t.Run("Test function Max", func(t *testing.T) {
		res1 := Max([]int{1, 2, 3, 8, 9, 3, 2, 1})
		res2 := Max([]int{1, 2, 1, 2, 2, 1})
		res3 := Max([]int{7, 1, 2, 9, 7, 2, 1})

		assert.Equal(t, 3, res1)
		assert.Equal(t, 2, res2)
		assert.Equal(t, 2, res3)
	})
}
