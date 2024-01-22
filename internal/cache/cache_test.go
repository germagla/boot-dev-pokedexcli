package cache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	c := NewCache(5 * time.Second)
	c.Add("foo", []byte("bar"))
	val, ok := c.Get("foo")
	if !ok {
		t.Errorf("expected to find key foo after adding it")
	}
	if string(val) != "bar" {
		t.Errorf("expected val to be bar, got %v", string(val))
	}
}

func TestReap(t *testing.T) {
	c := NewCache(5 * time.Second)
	c.Add("foo", []byte("bar"))
	time.Sleep(10 * time.Second)
	_, ok := c.Get("foo")
	if ok {
		t.Errorf("expected to not find key foo after reaping")
	}
}
