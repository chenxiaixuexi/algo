package ics4

type LinkNode struct {
	key, val  int
	pre, next *LinkNode
}

type LRUCache struct {
	hmap       map[int]*LinkNode
	cap        int
	head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	hmap := make(map[int]*LinkNode)
	head := &LinkNode{
		key:  0,
		val:  0,
		pre:  nil,
		next: nil,
	}
	tail := &LinkNode{
		key:  0,
		val:  0,
		pre:  nil,
		next: nil,
	}
	head.next = tail
	tail.pre = head
	cache := LRUCache{
		hmap: hmap,
		cap:  capacity,
		head: head,
		tail: tail,
	}
	return cache
}

func (this *LRUCache) Get(key int) int {
	return 0
}

func (this *LRUCache) Put(key int, value int) {

}
