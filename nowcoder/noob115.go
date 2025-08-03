package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param arr int整型一维数组
 * @return ListNode类
 */
func vectorToListnode(arr []int) *ListNode {
	head := &ListNode{}
	node := head
	for _, v := range arr {
		node.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		node = node.Next
	}
	return head.Next
}

func test1() {
	arr := []int{1, 3, 8}
	head := vectorToListnode(arr)
	node := head
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func main() {
	test1()
}
