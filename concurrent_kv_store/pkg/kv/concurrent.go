package kv

import (
	"hash/fnv"
	"sync"
)

const NumShards = 64

type entry map[string]string

type shard struct {
	mu   sync.RWMutex
	data entry
}

type ConcurrentMap struct {
	shards [NumShards]*shard
}

func NewConcurrentMap() *ConcurrentMap {
	m := &ConcurrentMap{}
	for i := 0; i < NumShards; i++ {
		m.shards[i] = &shard{
			data: make(entry),
		}
	}
	return m
}

func (m *ConcurrentMap) getShard(key string) *shard {
	h := fnv.New32a()
	h.Write([]byte(key))
	hashVal := h.Sum32()
	return m.shards[hashVal%NumShards]
}

func (m *ConcurrentMap) Put(key, value string) {
	s := m.getShard(key)
	s.mu.Lock()
	s.data[key] = value
	s.mu.Unlock()
}

func (m *ConcurrentMap) Get(key string) (string, bool) {
	s := m.getShard(key)
	s.mu.RLock()
	val, ok := s.data[key]
	s.mu.RUnlock()
	return val, ok
}

func (m *ConcurrentMap) Delete(key string) {
	s := m.getShard(key)
	s.mu.Lock()
	delete(s.data, key)
	s.mu.Unlock()
}
