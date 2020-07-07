package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
     Val int
     Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	pre := &ListNode{0, head}
	hint := pre

	for pre.Next != nil && pre.Next.Next != nil {
		a := pre.Next
		b := a.Next

		a.Next = b.Next
		b.Next = a

		pre.Next = b
		pre = a
	}
	return hint.Next
}


func main() {
	head := new(ListNode)
	head.Val = 1
	head.Next =  new(ListNode)

	head.Next.Val = 2
	head.Next.Next = new(ListNode)

	head.Next.Next.Val = 3
	head.Next.Next.Next = new(ListNode)

	head.Next.Next.Next.Val = 4
	head.Next.Next.Next.Next = nil
	cur := head
	for cur != nil {
		fmt.Print(cur.Val)
		cur = cur.Next

	}
	fmt.Println()

	newHead := swapPairs(head)
	//
	for newHead != nil {
		fmt.Print(newHead.Val)
		newHead = newHead.Next

	}
}
