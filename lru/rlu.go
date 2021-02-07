package main

import "fmt"

type LRUCache struct {
	head, tail *Node
	Keys       map[int]*Node
	Cap        int
}

type Node struct {
	Key, Val   int
	Prev, Next *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Keys: make(map[int]*Node), Cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Val = value
		this.Remove(node)
		this.Add(node)
		return
	} else {
		node = &Node{Key: key, Val: value}
		this.Keys[key] = node
		this.Add(node)
	}
	if len(this.Keys) > this.Cap {
		delete(this.Keys, this.tail.Key)
		this.Remove(this.tail)
	}
}

func (this *LRUCache) Add(node *Node) {
	node.Prev = nil
	node.Next = this.head
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
		this.tail.Next = nil
	}
}

func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		this.head = node.Next
		node.Next = nil
		return
	}
	if node == this.tail {
		this.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func main() {
	obj := Constructor(2)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.Put(1, 1)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	obj.Put(2, 2)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	param1 := obj.Get(1)
	fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
	obj.Put(3, 3)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	param1 = obj.Get(2)
	fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
	obj.Put(4, 4)
	fmt.Printf("obj = %v\n", MList2Ints(&obj))
	param1 = obj.Get(1)
	fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
	param1 = obj.Get(3)
	fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
	param1 = obj.Get(4)
	fmt.Printf("param_1 = %v obj = %v\n", param1, MList2Ints(&obj))
}

func MList2Ints(lru *LRUCache) [][]int {
	res := [][]int{}
	head := lru.head
	for head != nil {
		tmp := []int{head.Key, head.Val}
		res = append(res, tmp)
		head = head.Next
	}
	return res
}
