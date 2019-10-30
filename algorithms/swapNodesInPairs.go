package main

type ListNode struct {
 	Val int
 	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	pointer := head
	if head == nil {
		return head
	} else if head.Next == nil {
		return head
	} else if head.Next.Next == nil {
		head = head.Next
		pointer.Next = nil
		head.Next = pointer
	} else {
		head = head.Next
		pointer.Next = swapPairs(pointer.Next.Next)
		head.Next = pointer
	}
	return head	
}

// testing
func main() {
	first := &ListNode{ Val: 1,}
	second := &ListNode{ Val: 2,}
	third := &ListNode{ Val: 3,}
	fourth := &ListNode{ Val: 4,}
	fifth := &ListNode{ Val: 5,}
	sixth := &ListNode{ Val: 6,}

	first.Next = second
	second.Next = third
	third.Next = fourth
	fourth.Next = fifth
	fifth.Next = sixth

	swapPairs(first)
	pointer := second
	for i:=0; i<6; i++ {
		print(pointer.Val)
		pointer = pointer.Next
	}
}
