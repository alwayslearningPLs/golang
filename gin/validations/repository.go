package main

import "sync"

type hashMap[K comparable, V any] struct {
	l sync.Mutex
	m map[K]V
}

func newHashMap[K comparable, V any]() hashMap[K, V] {
	return hashMap[K, V]{
		m: make(map[K]V),
	}
}

func (h *hashMap[K, V]) get(key K) (V, bool) {
	h.l.Lock()
	defer h.l.Unlock()
	var v V
	if v, ok := h.m[key]; ok {
		return v, true
	}
	return v, false
}

func (h *hashMap[K, V]) getAll() map[K]V {
	h.l.Lock()
	defer h.l.Unlock()
	return h.m
}

func (h *hashMap[K, V]) put(key K, v V) V {
	h.l.Lock()
	defer h.l.Unlock()
	h.m[key] = v
	return v
}

func (h *hashMap[K, V]) del(key K) {
	h.l.Lock()
	defer h.l.Unlock()
	delete(h.m, key)
}

func (h *hashMap[K, V]) toArray() []V {
	h.l.Lock()
	defer h.l.Unlock()
	var (
		result = make([]V, len(h.m))
		count  int
	)
	for _, v := range h.m {
		result[count] = v
		count++
	}
	return result
}
