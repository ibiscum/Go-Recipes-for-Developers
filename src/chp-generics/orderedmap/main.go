package main

import (
	"fmt"
	"strings"
)

// An OrderedMap maps keys of type Key to values of type Value while
// maintaining insertion order
type OrderedMap[Key comparable, Value any] struct {
	m     map[Key]Value
	slice []Key
}

// NewOrderedMap returns a new ordere map with key type Key and value
// type Value
func NewOrderedMap[Key comparable, Value any]() *OrderedMap[Key, Value] {
	return &OrderedMap[Key, Value]{
		m:     make(map[Key]Value),
		slice: make([]Key, 0),
	}
}

// Add key:value to the map
func (m *OrderedMap[Key, Value]) Add(key Key, value Value) {
	_, exists := m.m[key]
	if exists {
		m.m[key] = value
	} else {
		m.slice = append(m.slice, key)
		m.m[key] = value
	}
}

// ValueAt returns the value at the given index
func (m *OrderedMap[Key, Value]) ValueAt(index int) Value {
	return m.m[m.slice[index]]
}

// KeyAt returns the key at the given index
func (m *OrderedMap[Key, Value]) KeyAt(index int) Key {
	return m.slice[index]
}

// Get returns the value corresponding to the key, and whether or not
// key exists
func (m *OrderedMap[Key, Value]) Get(key Key) (Value, bool) {
	v, bool := m.m[key]
	return v, bool
}

func (m *OrderedMap[Key, Value]) Len() int { return len(m.slice) }

func main() {
	m := NewOrderedMap[int, string]()
	for i, s := range strings.Split("This container preserves the insertion order", " ") {
		m.Add(i, s)
	}
	for i := 0; i < m.Len(); i++ {
		fmt.Println(m.ValueAt(i))
	}
}
