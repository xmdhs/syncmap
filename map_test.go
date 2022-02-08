package syncmap

import (
	"strconv"
	"testing"
)

func TestNewSyncMap(t *testing.T) {
	_ = SyncMap[string, int]{}
}

func TestSyncMapStore(t *testing.T) {
	s := SyncMap[string, int]{}
	s.Store("foo", 1)
	if v, _ := s.Load("foo"); v != 1 {
		t.Fatal("Expected 1, got ", v)
	}
}

func TestSyncMapLoadAndDelete(t *testing.T) {
	s := SyncMap[string, int]{}
	s.Store("foo", 1)
	if v, _ := s.LoadAndDelete("foo"); v != 1 {
		t.Fatal("Expected 1, got ", v)
	}
	if _, ok := s.Load("foo"); ok {
		t.Fatal("Expected false, got true")
	}
}

func TestSyncMapLoadOrStore(t *testing.T) {
	s := SyncMap[string, int]{}
	if _, ok := s.LoadOrStore("foo", 1); ok {
		t.Fatal("Expected false, got true")
	}

	if v, _ := s.Load("foo"); v != 1 {
		t.Fatal("Expected 1, got ", v)
	}

	if _, ok := s.LoadOrStore("foo", 1); !ok {
		t.Fatal("Expected true, got false")
	}
	if v, _ := s.LoadOrStore("foo", 1); v != 1 {
		t.Fatal("Expected 1, got ", v)
	}
}

func TestSyncMapRange(t *testing.T) {
	s := SyncMap[string, int]{}
	for i := 0; i < 10; i++ {
		s.Store("foo"+strconv.Itoa(i), i)
	}

	var count int
	s.Range(func(key string, value int) bool {
		count++
		if v, _ := s.Load(key); v != value {
			t.Fatal("Expected ", value, ", got ", v)
		}
		return true
	})
	if count != 10 {
		t.Fatal("Expected 10, got ", count)
	}
}

func TestSyncMapDelete(t *testing.T) {
	s := SyncMap[string, int]{}
	s.Store("foo", 1)
	s.Delete("foo")
	if _, ok := s.Load("foo"); ok {
		t.Fatal("Expected false, got true")
	}
}
