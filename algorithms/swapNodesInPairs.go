// NOT YET READY!!!!

package main

type ListNode struct {
 	Val int
 	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	//print(head.Next.Next)
	pointer := head.Next
	head.Next = head
	head = pointer
	//print(head.Val)
	print(head.Next)
	if pointer.Next != nil {
		return swapPairs(pointer)
	}
	
	return head

// 1 -> 2 // 1 -> 1, 2,	
}





func main() {
	l := &ListNode{ Val: 0, Next: &ListNode{ Val: 1, }, }
			// Val : 1, Next :  {
			// 	Val : 2, Next : {
			// 		Val :3 },}}}
	swapPairs(l)
}