package lru

import "fmt"

type LRUCache struct {
	capacity int
	maps     map[string]*Node
	head     *Node
	tail     *Node
}

func NewLRUCache(cap int) *LRUCache {
	if cap == 0 {
		cap = 1024
	}

	lru := &LRUCache{
		capacity: cap,
		maps:     make(map[string]*Node),
		head:     new(Node),
		tail:     new(Node),
	}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

func (c *LRUCache) Get(key string) (value interface{}, exist bool) {
	if node, ok := c.maps[key]; ok {
		c.moveTop(node)
		return node.value, true
	}
	return nil, false
}

func (c *LRUCache) Put(key string, value string) {
	node, ok := c.maps[key]
	if ok {
		node.value = value
		c.moveTop(node)
		return
	}
	if len(c.maps) >= c.capacity {
		c.moveTail()
	}
	node = NewNode(key, value)
	c.addTop(node)
}

func (c *LRUCache) PreList() {
	for node := c.head.next; node.next != nil; node = node.next {
		fmt.Println(node.key, " ", node.value)
	}
}

func (c *LRUCache) LastList() {
	for node := c.tail.pre; node.pre != nil; node = node.pre {
		fmt.Println(node.key, " ", node.value)
	}
}

func (c *LRUCache) Len() int {
	return len(c.maps)
}

func (c *LRUCache) moveTop(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre

	node.next = c.head.next
	node.pre = c.head

	c.head.next.pre = node
	c.head.next = node
}

func (c *LRUCache) moveTail() {
	fmt.Println(c.tail.pre.key)
	delete(c.maps, c.tail.pre.key)
	c.tail.pre.pre.next = c.tail
	c.tail.pre = c.tail.pre.pre
}

func (c *LRUCache) addTop(node *Node) {
	c.maps[node.key] = node
	node.next = c.head.next
	node.pre = c.head

	c.head.next.pre = node
	c.head.next = node
}
