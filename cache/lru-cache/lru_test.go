package cache

import "testing"

func TestLRUCache(t *testing.T) {
	lru_cache := New[int](3)
	val, found := lru_cache.Get("foo")
	if found {
		t.Error("incorrect result: found is true for non-existent key")
	}
	lru_cache.Update("foo", 10)
	val, found = lru_cache.Get("foo")
	if !found {
		t.Error("incorrect result: found is false for  an existent key")
	}
	if val != 10 {
		t.Errorf("incorrect result: expected cache value to be %d, got %d", 10, val)
	}
	lru_cache.Update("foo", 100)
	val, found = lru_cache.Get("foo")
	if !found {
		t.Error("incorrect result: found is false for  an existent key")
	}
	if val != 100 {
		t.Errorf("incorrect result: expected cache value to be %d, got %d", 100, val)
	}
	lru_cache.Update("bar", 420)
	val, _ = lru_cache.Get("bar")
	if val != 420 {
		t.Errorf("incorrect result: expected cache value to be %d, got %d", 420, val)
	}
	lru_cache.Update("baz", 1337)
	lru_cache.Update("ball", 69420)
	val, found = lru_cache.Get("ball")
	if !found {
		t.Error("incorrect result: found is false for an existent key")
	}
	if val != 69420 {
		t.Errorf("incorrect result: expected cache value to be %d, got %d", 69420, val)
	}
	val, found = lru_cache.Get("foo")
	if found {
		t.Error("incorrect result: found is true for an non-existent key")
	}
	if val != 0 {
		t.Errorf("incorrect result: expected cache value to be %d, got %d", 0, val)
	}

}
