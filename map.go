package gensync

import (
	"iter"
	"sync"
)

type Map[K comparable, V any] struct {
	m sync.Map
}

func (m *Map[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

func (m *Map[K, V]) Load(key K) (V, bool) {
	v, ok := m.m.Load(key)
	if !ok {
		var zero V
		return zero, false
	}

	return v.(V), true
}

func (m *Map[K, V]) Delete(key K) {
	m.m.Delete(key)
}

func (m *Map[K, V]) CompareAndDelete(key K, value V) bool {
	return m.m.CompareAndDelete(key, value)
}

func (m *Map[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m.m.CompareAndSwap(key, old, new)
}

func (m *Map[K, V]) LoadAndDelete(key K) (V, bool) {
	v, ok := m.m.LoadAndDelete(key)
	if !ok {
		var zero V
		return zero, ok
	}
	return v.(V), ok
}

func (m *Map[K, V]) LoadOrStore(key K, value V) (V, bool) {
	actual, loaded := m.m.LoadOrStore(key, value)
	return actual.(V), loaded
}

func (m *Map[K, V]) Range(fn func(K, V) bool) {
	m.m.Range(func(k, v any) bool {
		return fn(k.(K), v.(V))
	})
}

 func (m *Map[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		m.m.Range(func(k, v any) bool {
			return yield(k.(K), v.(V))
		})
	}
}

func (m *Map[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		m.m.Range(func(k, _ any) bool {
			return yield(k.(K))
		})
	}
}

func (m *Map[K, V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		m.m.Range(func(_, v any) bool {
			return yield(v.(V))
		})
	}
}
