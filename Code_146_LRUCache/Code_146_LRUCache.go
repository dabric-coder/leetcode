package main

import "fmt"

type Node struct {
	val interface{}
	last *Node
	next *Node
}

type NodeDoubleLinkedList struct {
	head *Node
	tail *Node
}

func NewNodeDoubleLinkedList() NodeDoubleLinkedList {
	return NodeDoubleLinkedList{nil, nil}
}

func (this *NodeDoubleLinkedList) addNode(newNode *Node) {
	if newNode == nil {
		return
	}
	if this.head == nil {
		this.head = newNode
		this.tail = newNode
	} else {
		this.tail.next = newNode
		newNode.last = this.tail
		this.tail = newNode
	}

}

func (this *NodeDoubleLinkedList) moveNodeToTail(node *Node) {
	if node == this.tail {
		return
	}

	if node == this.head {
		this.head = this.head.next
		this.head.last = nil
	} else {
		node.last.next = node.next
		node.next.last = node.last
	}
	node.last = this.tail
	node.next = nil
	this.tail.next = node
	this.tail = node
}

func (this *NodeDoubleLinkedList) removeHeadNode() *Node {
	if this.head == nil {
		return nil
	}

	res := this.head

	if this.head == this.tail {
		this.head = nil
		this.tail = nil
	} else {
		this.head = res.next
		res.next = nil
		this.head.last = nil
	}
	return res
}



type LRUCache struct {
	keyNodeMap map[interface{}]*Node
	nodeKeyMap map[*Node]interface{}
	capacity int
	nodeList NodeDoubleLinkedList
}


func Constructor(capacity int) LRUCache {
	if capacity < 1 {
		panic("should be more than 0.")
	}
	keyNodeMap := make(map[interface{}]*Node, 0)
	nodeKeyMap := make(map[*Node]interface{}, 0)
	return LRUCache{keyNodeMap, nodeKeyMap,capacity, NewNodeDoubleLinkedList()}
}

func (this *LRUCache) Get(key int) int {
	res, ok := this.keyNodeMap[key]
	if ok {
		//res := this.keyNodeMap[key]
		this.nodeList.moveNodeToTail(res)
		return res.val.(int)
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	node, ok := this.keyNodeMap[key]
	if ok {
		//node := this.keyNodeMap[key]
		node.val = value
		this.nodeList.moveNodeToTail(node)
	} else {
		newNode := &Node{value, nil, nil}
		this.keyNodeMap[key] = newNode
		this.nodeKeyMap[newNode] = key
		this.nodeList.addNode(newNode)
		if len(this.keyNodeMap) == this.capacity + 1{
			this.removeMostUnusedCache()
		}
	}
}



func (this *LRUCache) removeMostUnusedCache() {
	removeNode := this.nodeList.removeHeadNode()
	removeNodeKey := this.nodeKeyMap[removeNode]
	delete(this.keyNodeMap, removeNodeKey)
	delete(this.nodeKeyMap, removeNode)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	myCache := Constructor(3)
	myCache.Put(10, 1)
	myCache.Put(11,2)
	myCache.Put(12, 3)
	fmt.Println(myCache.Get(11))
	fmt.Println(myCache.Get(10))
	myCache.Put(13, 4)
	fmt.Println(myCache.Get(13))
	fmt.Println(myCache.Get(12))
}