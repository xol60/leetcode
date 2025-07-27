package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first := head
	for i := 0; i < n; i++ {
		first = first.Next
	}
	if first == nil {
		return head.Next
	}
	first = first.Next
	second := head.Next
	prevSecond := head
	for first != nil {
		first = first.Next
		prevSecond = second
		second = second.Next
	}
	//fmt.Println(prevSecond.Val)
	prevSecond.Next = second.Next
	return head
}
