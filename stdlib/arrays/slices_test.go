package arrays

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlices(t *testing.T) {
	t.Run("slices are not arrays", func(t *testing.T) {
		var (
			slice = []int{1, 2, 3, 4}
			array = [...]int{1, 2, 3, 4}
		)

		assert.NotEqual(t, array, slice)
		assert.NotEqual(t, reflect.TypeOf(array).Kind(), reflect.TypeOf(slice).Kind())
	})

	t.Run("making slices with n length and n capacity", func(t *testing.T) {
		var slice = make([]byte, 10)

		for i := 0; i < len(slice); i++ {
			assert.Equal(t, byte(0), slice[i])
		}
	})

	t.Run("making slices with 0 length and n capacity", func(t *testing.T) {
		var slice = make([]byte, 0, 10)

		for i := 0; i < cap(slice); i++ {
			assert.Panics(t, func() { _ = slice[i] })
		}
	})

	t.Run("making slices with 0 length and n capacity appending values on demand", func(t *testing.T) {
		var (
			slice = make([]byte, 0, 10)

			wantLen = 10
			wantStr = "[0 1 2 3 4 5 6 7 8 9]"
		)

		for i := 0; i < cap(slice); i++ {
			slice = append(slice, byte(i))
		}

		assert.Equal(t, wantLen, len(slice))
		assert.Equal(t, cap(slice), len(slice))
		assert.Equal(t, wantStr, fmt.Sprintf("%v", slice))
	})

	t.Run("nil slice", func(t *testing.T) {
		var (
			input []int  // nil slice
			arr   [5]int // array with 5 len/cap
		)

		assert.Equal(t, reflect.Slice, reflect.TypeOf(input).Kind())
		assert.Equal(t, 0, len(input))
		assert.Equal(t, 0, cap(input))

		assert.Equal(t, reflect.Array, reflect.TypeOf(arr).Kind())
		assert.Equal(t, 5, len(arr))
		assert.Equal(t, 5, cap(arr))
	})
}

func TestCompArr(t *testing.T) {
	var slice = []string{"g", "o", "l", "a", "n", "g"}

	for i := 0; i < len(slice); i++ {
		t.Run(fmt.Sprintf("slicing in %d", i), func(t *testing.T) {
			var (
				wantLen = i
				wantCap = cap(slice)
			)

			s := slice[:i]

			assert.Equal(t, wantLen, len(s))
			assert.Equal(t, wantCap, cap(s))
		})
	}

	t.Run("slicing from array/slice", func(t *testing.T) {
		var arr = []int{1, 2, 3, 4, 5}

		assert.Equal(t, arr, arr[:])
	})

}
