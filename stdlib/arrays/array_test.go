package arrays

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	i interface {
		foo() string
	}

	s struct {
		string
		int
		float64
		i
	}

	props struct {
		len, cap int
		content  string
	}
)

func arrProps[T int | *int | string | s](arr [5]T) props {
	return props{
		len:     len(arr),
		cap:     cap(arr),
		content: fmt.Sprintf("%v", arr),
	}
}

func TestArrays(t *testing.T) {
	var wantProps = func(content string) props {
		return props{
			len:     5,
			cap:     5,
			content: content,
		}
	}

	t.Run("int array", func(t *testing.T) {
		var (
			i [5]int

			want = wantProps("[0 0 0 0 0]")
		)

		got := arrProps(i)

		assert.Equal(t, want, got)
	})

	t.Run("int* array", func(t *testing.T) {
		var (
			i [5]*int

			want = wantProps("[<nil> <nil> <nil> <nil> <nil>]")
		)

		got := arrProps(i)

		assert.Equal(t, want, got)
	})

	t.Run("string array", func(t *testing.T) {
		var (
			i [5]string

			want = wantProps("[    ]")
		)

		got := arrProps(i)

		assert.Equal(t, want, got)
	})

	t.Run("struct array", func(t *testing.T) {
		var (
			i [5]s

			want = wantProps("[{ 0 0 <nil>} { 0 0 <nil>} { 0 0 <nil>} { 0 0 <nil>} { 0 0 <nil>}]")
		)

		got := arrProps(i)

		assert.Equal(t, want, got)
	})
}

func TestWaysOfInitializingArrays(t *testing.T) {
	for _, each := range []struct {
		description string
		input       [5]int
		want        props
	}{
		{
			description: "init arrays with fixed size splicitly wrote",
			input:       [5]int{1, 2, 3, 4, 5},
			want: props{
				len:     5,
				cap:     5,
				content: "[1 2 3 4 5]",
			},
		},
		{
			description: "init arrays without writting the size",
			input:       [...]int{1, 2, 3, 4, 5},
			want: props{
				len:     5,
				cap:     5,
				content: "[1 2 3 4 5]",
			},
		},
	} {
		t.Run(each.description, func(t *testing.T) {
			assert.Equal(t, each.want, arrProps(each.input))
		})
	}

	t.Run("panic when converting interface{} of type []int to [5]int even when the value has 5 of size (it is an slice instead of an array)", func(t *testing.T) {
		var tmp any = []int{1, 2, 3, 4, 5}

		assert.Panics(t, func() {
			_ = tmp.([5]int) // This will break
		})
	})

	t.Run("reflecting arrays", func(t *testing.T) {
		var (
			arr = reflect.ArrayOf(5, reflect.TypeOf(1))

			want = props{
				len: 5,
				cap: 5,
			}
		)

		got := props{
			len: arr.Len(),
			cap: arr.Len(),
		}

		assert.Equal(t, reflect.Array, arr.Kind())
		assert.NotEqual(t, reflect.Slice, arr.Kind())
		assert.True(t, arr.ConvertibleTo(reflect.TypeOf([5]int{1, 2, 3, 4, 5})))
		assert.Equal(t, want, got)
	})
}
