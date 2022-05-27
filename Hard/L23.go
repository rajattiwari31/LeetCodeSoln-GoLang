package Hard

/*

* Approach -
	- This is similar to Merge two sorted LinkdList
	- Will follow the same, take two list and then sort

* TimeComplexity
	- O(N) - for sorting LinkedList
	- O(kN) - for k List
* SpaceComplexity
	- O(1)
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		l1 := lists[0]
		l2 := lists[1]

		//Removing the first two ListNode from the Lists
		lists = lists[2:]

		l3 := sortLists(l1, l2)

		lists = append(lists, l3)
	}

	return lists[0]
}

func sortLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.Val <= l2.Val {
		l1.Next = sortLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = sortLists(l1, l2.Next)
		return l2
	}
}
