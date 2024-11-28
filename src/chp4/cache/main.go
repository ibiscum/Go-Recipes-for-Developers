package main

import (
	"fmt"
	"sync"

	"golang.org/x/exp/rand"
)

type Object struct {
	ID   string
	Name string
	// Other fields
}

type ObjectCache struct {
	mutex  sync.RWMutex
	values map[string]*Object
}

// Initialize and return a new instance of the cache
func NewObjectCache() *ObjectCache {
	return &ObjectCache{
		values: make(map[string]*Object),
	}
}

// Get an object from the cache
func (cache *ObjectCache) Get(key string) (*Object, bool) {
	cache.mutex.RLock()
	obj, exists := cache.values[key]
	cache.mutex.RUnlock()
	return obj, exists
}

// Put an object into the cache with the given key
func (cache *ObjectCache) Put(key string, value *Object) {
	cache.mutex.Lock()
	cache.values[key] = value
	cache.mutex.Unlock()
}

func main() {
	cache := NewObjectCache()
	objects := make([]Object, 0)
	// Create some objects
	for i := 0; i < 1000; i++ {
		objects = append(objects, Object{
			ID:   fmt.Sprint(i),
			Name: fmt.Sprintf("Name: %d", i),
		})
	}

	// Ten goroutines add objects to the cache randomly
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				// Pick a random object and add it to the cache
				n := rand.Intn(len(objects))
				cache.Put(objects[n].ID, &objects[n])
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Cache has %d objects\n", len(cache.values))
}
