package main

import (
	"errors"
	"fmt"
	"sync"

	"golang.org/x/exp/rand"
)

type cacheItem struct {
	sync.Once
	object *Object
}

type Object struct {
	ID   string
	Name string
	// Other fields
}

type ObjectCache struct {
	mutex         sync.RWMutex
	values        map[string]*cacheItem
	getObjectFunc func(string) (*Object, error)
}

// Initialize and return a new instance of the cache
func NewObjectCache(getObjectFunc func(string) (*Object, error)) *ObjectCache {
	return &ObjectCache{
		values:        make(map[string]*cacheItem),
		getObjectFunc: getObjectFunc,
	}
}

func (item *cacheItem) get(key string, cache *ObjectCache) (err error) {
	// Calling item.Once.Do
	item.Do(func() {
		item.object, err = cache.getObjectFunc(key)
	})
	return
}

func (cache *ObjectCache) Get(key string) (*Object, error) {
	cache.mutex.RLock()
	object, exists := cache.values[key]
	cache.mutex.RUnlock()
	if exists {
		return object.object, nil
	}
	cache.mutex.Lock()
	object, exists = cache.values[key]
	if !exists {
		object = &cacheItem{}
		cache.values[key] = object
	}
	cache.mutex.Unlock()
	err := object.get(key, cache)
	return object.object, err
}

var ErrNotFound = errors.New("not found")

func main() {

	objects := make([]Object, 0)
	// Create some objects
	for i := 0; i < 1000; i++ {
		objects = append(objects, Object{
			ID:   fmt.Sprint(i),
			Name: fmt.Sprintf("Name: %d", i),
		})
	}

	calledNTimes := 0
	cache := NewObjectCache(func(id string) (*Object, error) {
		calledNTimes++
		for i := range objects {
			if objects[i].ID == id {
				return &objects[i], nil
			}
		}
		return nil, ErrNotFound
	})

	// Ten goroutines get objects from the cache randomly
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				// Pick a random object and add it to the cache
				n := rand.Intn(len(objects))
				_, err := cache.Get(objects[n].ID)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}
	wg.Wait()
	fmt.Printf("GetObject func called %d times\n", calledNTimes)
}
