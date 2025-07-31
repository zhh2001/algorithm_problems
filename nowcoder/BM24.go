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
func inorderTraversal(root *TreeNode) []int {
	vals := make([]int, 0)
	if root == nil {
		return vals
	}
	vals = append(vals, inorderTraversal(root.Left)...)
	vals = append(vals, root.Val)
	vals = append(vals, inorderTraversal(root.Right)...)
	return vals
}

func test1() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: nil,
	}
	vals := inorderTraversal(root)
	fmt.Println(vals)
}

func test2() {
	root := (*TreeNode)(nil)
	vals := inorderTraversal(root)
	fmt.Println(vals)
}

func test3() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	vals := inorderTraversal(root)
	fmt.Println(vals)
}

func test4() {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	vals := inorderTraversal(root)
	fmt.Println(vals)
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
