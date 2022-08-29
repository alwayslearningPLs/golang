package syncdoapproach

import (
	"sync"
	"sync/atomic"
)

type singleton struct {
	count int32
}

var (
	counter *singleton

	counterOnce sync.Once
)

// GetInstance is used to create a singleton struct in a thread-safe manner
func GetInstance() *singleton {
	counterOnce.Do(func() {
		counter = new(singleton)
	})
	return counter
}

// AddOne is used to add one to the counter in a thread-safe manner
func (s *singleton) AddOne() int32 {
	return atomic.AddInt32(&s.count, 1)
}
