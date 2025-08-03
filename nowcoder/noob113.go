package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param head ListNode类
 * @return ListNode类
 */
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	node := &ListNode{
		Val:  head.Val,
		Next: nil,
	}
	for p := head.Next; p != nil; p = p.Next {
		node = &ListNode{
			Val:  p.Val,
			Next: node,
		}
	}
	return node
}

func PrintList(head *ListNode) {
	if head == nil {
		fmt.Println(head)
		return
	}

	for p := head; p != nil; p = p.Next {
		fmt.Printf("%d\t", p.Val)
	}
	fmt.Println()
}

func test1() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	newHead := ReverseList(head)
	PrintList(newHead)
}

func test2() {
	head := (*ListNode)(nil)
	newHead := ReverseList(head)
	PrintList(newHead)
}

func main() {
	test1()
	test2()
}
