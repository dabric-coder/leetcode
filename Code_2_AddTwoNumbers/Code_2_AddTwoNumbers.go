package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	n := 0
	ca := 0
	n1 := 0
	n2 := 0
	c1 := l1
	c2 := l2
	var cur = &ListNode{0, nil}
	var node *ListNode = nil
	newHead := cur

	for c1 != nil || c2 != nil {
		if c1 != nil {
			n1 = c1.Val
		} else {
			n1 = 0
		}

		if c2 != nil {
			n2 = c2.Val
		} else {
			n2 = 0
		}

		n = n1 + n2 + ca
		node = &ListNode{n % 10, nil}
		cur.Next = node
		cur = node

		if c1 != nil {
			c1 = c1.Next
		} else {
			c1 = nil
		}

		if c2 != nil {
			c2 = c2.Next
		} else {
			c2 = nil
		}
		ca = n / 10

	}

	if ca == 1 {
		node = &ListNode{1, nil}
		cur.Next = node
		cur = node
	}
	return newHead.Next
}

func main() {

}
