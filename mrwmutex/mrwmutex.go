package mrwmutex
 
import (
  "math"
  "math/rand"
  "sync/atomic"
  "time"
)

// A MultiRWMutex is a reader/writer mutual exclusion lock.
// The lock can be held by an arbitrary number of readers, or
// an arbitray numbers of writers.

type MultiRWMutex struct {
  rwCounts atomic.Int64
}

func (m *MultiRWMutex) incReaders() bool {
  result := m.rwCounts.Add(1<<32)
  writers := result & math.MaxInt32
  success := writers == 0
  if !success {
    m.decReaders()
  }
  return success
}

func (m *MultiRWMutex) decReaders() {
  m.rwCounts.Add(-1<<32)
}

func (m *MultiRWMutex) RLock() {
  for !m.incReaders() {
    time.Sleep(time.Duration(rand.Intn(5))*time.Nanosecond)
  }
}

func (m *MultiRWMutex) RUnlock() {
  m.decReaders()
}

func (m *MultiRWMutex) incWriters() bool {
  result := m.rwCounts.Add(1)
  readers := result>>32
  success := readers == 0
  if !success {
    m.decWriters()
  }
  return success
}

func (m *MultiRWMutex) decWriters() {
  m.rwCounts.Add(-1)
}

func (m *MultiRWMutex) WLock() {
  for !m.incWriters() {
    time.Sleep(time.Duration(rand.Intn(5))*time.Nanosecond)
  }
}
func (m *MultiRWMutex) WUnlock() {
  m.decWriters()
}
