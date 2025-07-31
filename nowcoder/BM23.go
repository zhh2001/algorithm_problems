package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 *
 * @param root TreeNode类
 * @return int整型一维数组
 */
func preorderTraversal(root *TreeNode) []int {
	vals := make([]int, 0)
	if root == nil {
		return vals
	}
	vals = append(vals, root.Val)
	vals = append(vals, preorderTraversal(root.Left)...)
	vals = append(vals, preorderTraversal(root.Right)...)
	return vals
}

func main() {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	vals := preorderTraversal(root)
	fmt.Println(vals)
}
