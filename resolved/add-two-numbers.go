package resolved

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{Val: 0, Next: nil}
	currentNode := l3
	var carry = 0
	for l1 != nil || l2 != nil || carry > 0 {
		currentNode.Next = &ListNode{}
		currentNode = currentNode.Next

		var (
			l1Value = 0
			l2Value = 0
		)

		if l1 != nil {
			l1Value = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Value = l2.Val
			l2 = l2.Next
		}

		sum := l1Value + l2Value + carry

		currentNode.Val = sum % 10
		carry = sum / 10
	}
	return l3.Next
}
