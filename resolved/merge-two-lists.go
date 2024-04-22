package resolved

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	current := dummy

	for list1 != nil && list2 != nil {
		current.Next = &ListNode{}
		current = current.Next

		if list1.Val < list2.Val {
			current.Val = list1.Val
			list1 = list1.Next
		} else {
			current.Val = list2.Val
			list2 = list2.Next
		}

	}

	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}

	return dummy.Next
}
