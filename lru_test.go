package lru

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLRUCache_Put(t *testing.T) {
	cache := NewLRUCache(5)
	cache.Put("1","value1")
	cache.Put("2","value2")
	cache.Put("3","value3")
	cache.Put("4","value4")
	cache.Put("5","value5")
	cache.Put("6","value6")
	assert.Equal(t,5,cache.Len())
}

func TestLRUCache_Get(t *testing.T) {
	cache := NewLRUCache(5)
	cache.Put("1","value1")
	cache.Put("2","value2")
	cache.Put("3","value3")
	cache.Put("4","value4")
	cache.Put("5","value5")
	val,ok :=cache.Get("1")
	assert.Equal(t,true,ok)
	assert.Equal(t,"value1",val.(string))
	val,ok =cache.Get("2")
	assert.Equal(t,true,ok)
	assert.Equal(t,"value2",val.(string))
	val,ok =cache.Get("3")
	assert.Equal(t,true,ok)
	assert.Equal(t,"value3",val.(string))
	val,ok =cache.Get("4")
	assert.Equal(t,true,ok)
	assert.Equal(t,"value4",val.(string))
	val,ok =cache.Get("5")
	assert.Equal(t,true,ok)
	assert.Equal(t,"value5",val.(string))
}