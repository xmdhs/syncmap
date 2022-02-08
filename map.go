package syncmap

import "sync"

type SyncMap[K, V any] struct {
	m sync.Map
}

func (s *SyncMap[K, V]) Delete(key K) {
	s.m.Delete(key)
}

func (s *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	val, ok := s.m.Load(key)
	if ok {
		value = val.(V)
	}
	return value, ok
}

func (s *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	val, loaded := s.m.LoadAndDelete(key)
	if loaded {
		value = val.(V)
	}
	return value, loaded
}

func (s *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	val, loaded := s.m.LoadOrStore(key, value)
	if loaded {
		actual = val.(V)
	}
	return actual, loaded
}

func (s *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	s.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (s *SyncMap[K, V]) Store(key K, value V) {
	s.m.Store(key, value)
}
