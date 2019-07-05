package linkedlist

//DeleteDuplicate :Write code to remove duplicates from an unsorted linked list.
func DeleteDuplicate(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	set := make(map[int]int)
	dummy := &ListNode{0, nil}
	pre := dummy
	p := head
	for p != nil {
		if _, ok := set[p.Val]; !ok {
			pre.Next = p
			pre = pre.Next
			set[p.Val] = 1
		}
		p = p.Next
	}
	pre.Next = p
	return dummy.Next
}
