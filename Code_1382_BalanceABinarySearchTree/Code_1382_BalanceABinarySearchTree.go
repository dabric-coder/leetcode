package main

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func balanceBST(root *TreeNode) *TreeNode {
	arr := make([]int, 0)
	inorder(root, &arr)
	return buildTree(arr, 0, len(arr)-1)
}

func inorder(root *TreeNode, arr *[]int) {
	if root != nil {
		inorder(root.Left, arr)
		*arr = append(*arr, root.Val)
		inorder(root.Right, arr)
	}
}

func buildTree(arr []int, start, end int) *TreeNode{
	if start > end {
		return nil
	}

	mid := start + (end-start) >> 1
	root := &TreeNode{arr[mid], nil, nil}
	root.Left = buildTree(arr, start, mid-1)
	root.Right = buildTree(arr, mid+1, end)
	return root
}
