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
func postorderTraversal(root *TreeNode) []int {
	vals := make([]int, 0)
	if root == nil {
		return vals
	}
	vals = append(vals, postorderTraversal(root.Left)...)
	vals = append(vals, postorderTraversal(root.Right)...)
	vals = append(vals, root.Val)
	return vals
}

func test1() {
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
	vals := postorderTraversal(root)
	fmt.Println(vals)
}

func test2() {
	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	vals := postorderTraversal(root)
	fmt.Println(vals)
}

func main() {
	test1()
	test2()
}
