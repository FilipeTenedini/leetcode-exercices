package resolved

import (
	"fmt"
)

func PrintResolved() {
	fmt.Printf("Two sum -> The position of slice when the sum result in target is %v\n\n", TwoSum([]int{1, 3, 4, 5, 6}, 6))
	fmt.Printf("Search Standard on Word -> The word has equal qtt of each letter or error qntt less second param? %v\n\n", SearchStandardOnWord("pepesq", 1))
	fmt.Printf("Valid Parentheses -> The string has valid brackets? %v\n\n", ValidParentheses(")()"))
	fmt.Printf("Add two numbers -> retun linked list with sum of numbers: If you want to see, print one for in value %v\n\n", AddTwoNumbers(returnTwomNodeLists()[0], returnTwomNodeLists()[1]))
	fmt.Printf("Roman to int -> Translate roman number to int? %v\n\n", RomanToInt("XVIII"))
	fmt.Printf("Longest Common Prefix -> Return the largest prefix in common of slice of strings or '' if no has common prefix%v\n\n", LongestCommonPrefix([]string{"prefix", "prefeito"}))
	fmt.Printf("merge two lists -> retun linked list with sorted and merged two list. If you want to see, print one for in value %v\n\n", MergeTwoLists(returnTwomNodeLists()[0], returnTwomNodeLists()[1]))
}

func returnTwomNodeLists() [2]*ListNode {
	listNode1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4, Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	listNode2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list := [2]*ListNode{listNode1, listNode2}

	return list
}
