package main

type ListNodes struct {
	Val  int
	Next *ListNodes
}

//给入头节点
func deleteDuplicates(head *ListNodes) *ListNodes {
	var pre, cur *ListNodes = head, nil
	if head != nil {
		cur = head.Next
	}
	for cur != nil {
		if pre.Val == cur.Val {
			pre.Next = cur.Next
			cur = pre.Next
		} else {
			pre = pre.Next
			cur = cur.Next
		}
	}
	return head
}
