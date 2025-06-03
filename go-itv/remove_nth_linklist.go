package main

type ListNode struct {
	Val int

	Next *ListNode
}

func removeNthLastNode(head *ListNode, n int) *ListNode {

	left, right := head, head

	for i := 0; i < n && right != nil; i++ {
		right = right.Next
	}
	
	if right == nil {
		return head
	}

	for right.Next != nil {
		right = right.Next
		left = left.Next
	}

	left.Next = left.Next.Next

	return head

}
