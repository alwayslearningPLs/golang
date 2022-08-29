package initapproach

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	var (
		wg sync.WaitGroup

		want int32 = 100000
	)

	defer func() {
		wg.Wait()
		assert.Equal(t, want, GetInstance().count)
	}()

	for i := int32(0); i < want; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetInstance().AddOne()
		}()
	}
}
