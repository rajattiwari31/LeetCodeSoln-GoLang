package Easy

/*
	Algorithm:
		- Create dummy node
		- And check one by one and reference the pointer to the other node
		- No need to create a new ListNode
	TimeComplexity
		- O(n)
	SpaceComplexity
		- O(1) - as no creation of new Node, only referencing
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}

	head := &ListNode{}

	dummy := head

	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			dummy.Next = list1
			list1 = list1.Next
		} else {
			dummy.Next = list2
			list2 = list2.Next
		}
	}

	if list1 == nil && list2 != nil {
		dummy.Next = list2
	} else if list2 == nil && list1 != nil {
		dummy.Next = list1
	}

	return head.Next

}
