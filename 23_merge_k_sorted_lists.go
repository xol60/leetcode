package main

import "math"

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	res := &ListNode{Val: math.MinInt32}
	for _, v := range lists {
		res = mergeTwoLists(res, v)
	}
	return res.Next
}
