package initapproach

import "sync/atomic"

type singleton struct {
	count int32
}

var s *singleton

func init() {
	s = new(singleton)
}

// GetInstance returns the singleton instance that 'init' function have just created for us
func GetInstance() *singleton {
	return s
}

// AddOne is used to add one to the counter in a thread-safe manner
func (s *singleton) AddOne() int32 {
	return atomic.AddInt32(&s.count, 1)
}
