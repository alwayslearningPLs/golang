package mutexapproach

import (
	"sync"
	"sync/atomic"
)

type singleton struct {
	count int32
}

var (
	counter *singleton

	counterLock sync.Mutex
)

// GetInstance is used to create a singleton struct in a thread-safe manner
func GetInstance() *singleton {
	if counter == nil {
		counterLock.Lock()
		defer counterLock.Unlock()
		if counter == nil {
			counter = new(singleton)
		}
	}
	return counter
}

// AddOne is used to add one to the counter in a thread-safe manner
func (s *singleton) AddOne() int32 {
	atomic.AddInt32(&s.count, 1)
	return s.count
}
