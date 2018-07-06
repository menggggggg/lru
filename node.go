package lru

type Node struct {
	key   string
	value interface{}
	pre   *Node
	next  *Node
}

func NewNode(key string,value interface{})*Node{
	return  &Node{
		key:key,
		value:value,
	}
}
