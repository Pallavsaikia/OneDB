package cache

import (
    "container/list"
    "sync"
    "time"
)

// CacheItem represents an item stored in the cache.
type CacheItem struct {
    Key        string
    Value      interface{}
    ExpiryTime time.Time
}

// Cache is a simple in-memory cache.
type Cache struct {
    items      map[string]*list.Element
    eviction   *list.List
    mutex      sync.RWMutex
    maxSize    int // Maximum number of items in the cache
    avgItemSize int // Average size of items in KB
    currentSize int // Current size of the cache in KB
}

// NewCache creates a new instance of Cache with a specified maximum size in KB.
func NewCache(maxSizeKB int) *Cache {
    return &Cache{
        items:    make(map[string]*list.Element),
        eviction: list.New(),
        maxSize:  maxSizeKB,
    }
}

// Set adds an item to the cache with the specified key and expiration time.
func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    // Check if the key already exists in the cache
    if elem, found := c.items[key]; found {
        // If the key exists, update its value and move it to the front of the eviction list
        elem.Value.(*CacheItem).Value = value
        elem.Value.(*CacheItem).ExpiryTime = time.Now().Add(ttl)
        c.eviction.MoveToFront(elem)
        return
    }

    // If adding the new item exceeds the maximum size, evict the least recently used items until the cache size is within the limit
    for c.currentSize + c.getItemSize(value) > c.maxSize {
        c.evictLeastRecentlyUsed()
    }

    // Add the new item to the cache
    item := &CacheItem{
        Key:        key,
        Value:      value,
        ExpiryTime: time.Now().Add(ttl),
    }
    c.items[key] = c.eviction.PushFront(item)
    c.currentSize += c.getItemSize(value)
}

// Get retrieves an item from the cache by key.
func (c *Cache) Get(key string) (interface{}, bool) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if elem, found := c.items[key]; found {
        item := elem.Value.(*CacheItem)
        // Check if the item has expired
        if time.Now().After(item.ExpiryTime) {
            delete(c.items, key)
            c.eviction.Remove(elem)
            c.currentSize -= c.getItemSize(item.Value)
            return nil, false
        }
        // Move the accessed item to the front of the eviction list
        c.eviction.MoveToFront(elem)
        return item.Value, true
    }
    return nil, false
}

// evictLeastRecentlyUsed removes the least recently used item from the cache.
func (c *Cache) evictLeastRecentlyUsed() {
    if c.eviction.Len() == 0 {
        return
    }
    elem := c.eviction.Back()
    delete(c.items, elem.Value.(*CacheItem).Key)
    c.eviction.Remove(elem)
}

// getItemSize returns the size of an item in KB.
func (c *Cache) getItemSize(value interface{}) int {
    // For demonstration purposes, let's assume the average size of an item is 1 KB
    return 1
}

