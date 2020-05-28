package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{}
	if root == nil{
		return nil
	}
	rs:=[][]int{}
	queue = append(queue,root)

	for len(queue)>0{
		tmpRes := []int{}
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		rs = append(rs, tmpRes)
	}
	return rs
}


func main() {
	node1 := new(TreeNode)
	node1.Val = 3

	node2 := new(TreeNode)
	node2.Val = 9

	node3 := new(TreeNode)
	node3.Val = 20

	node4 := new(TreeNode)
	node4.Val = 15

	node5 := new(TreeNode)
	node5.Val = 7

	node1.Left = node2
	node1.Right = node3

	node3.Left = node4
	node3.Right = node5

	fmt.Println(levelOrder(node1))
}