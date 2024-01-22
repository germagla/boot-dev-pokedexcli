package cache

import (
	"encoding/json"
	"fmt"
	"github.com/germagla/boot-dev-pokedexcli/internal/pokeapi"
	"sync"
	"time"
)

type cacheEntry struct {
	lastAccessed time.Time
	val          []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[key] = cacheEntry{
		lastAccessed: time.Now(),
		val:          val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	entry.lastAccessed = time.Now()
	fmt.Println("Cache hit!")
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		c.reap()
	}

}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.entries {
		if time.Since(entry.lastAccessed) > c.interval {
			delete(c.entries, key)
		}
	}
}

func (c *Cache) AddAreaList(key string, val pokeapi.LocationAreaList) error {
	bytes, err := json.Marshal(val)
	if err != nil {
		return err
	}
	c.Add(key, bytes)
	return nil
}

func (c *Cache) GetAreaList(key string) (pokeapi.LocationAreaList, bool) {
	bytes, ok := c.Get(key)
	if !ok {
		return pokeapi.LocationAreaList{}, false
	}
	var list pokeapi.LocationAreaList
	err := json.Unmarshal(bytes, &list)
	if err != nil {
		return pokeapi.LocationAreaList{}, false
	}
	return list, true
}

func (c *Cache) AddLocationArea(key string, val pokeapi.LocationArea) error {
	bytes, err := json.Marshal(val)
	if err != nil {
		return err
	}
	c.Add(key, bytes)
	return nil
}
func (c *Cache) GetLocationArea(key string) (pokeapi.LocationArea, bool) {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	key = baseURL + key
	bytes, ok := c.Get(key)
	if !ok {
		return pokeapi.LocationArea{}, false
	}
	var area pokeapi.LocationArea
	err := json.Unmarshal(bytes, &area)
	if err != nil {
		return pokeapi.LocationArea{}, false
	}
	return area, true
}
