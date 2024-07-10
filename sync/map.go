package sync

import "sync"

type Map[K comparable, V any] struct {
	m *sync.Map
}

func New[K comparable, V any]() Map[K, V] {
	return Map[K, V]{m: &sync.Map{}}
}

func (m Map[K, V]) Load(key K) (V, bool) {
	val, ok := m.m.Load(key)
	if ok {
		return val.(V), ok
	} else {
		var dummy V
		return dummy, ok
	}
}

func (m Map[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

func (m Map[K, V]) Delete(key K) {
	m.m.Delete(key)
}

func (m Map[K, V]) Range(f func(K, V) bool) {
	m.m.Range(func(key any, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (m Map[K, V]) LoadOrCreate(key K, createFunc func() V) (value V, loaded bool) {
	v, ok := m.m.Load(key)
	if ok {
		return v, false
	}

	v = createFunc()
	return m.m.LoadOrStore(key, v)
}
