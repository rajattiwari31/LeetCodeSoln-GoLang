package Hard

/*

* Approach -
	First Approach -
		- This is similar to Merge two sorted LinkdList
		- Will follow the same, take two list and then sort
	Second Approach -
		- In this we can reduce some iterations by paring all the List in one go
		- then follow the same for other iter

* TimeComplexity
	- O(N) - for sorting LinkedList
	- O(kN) - for k List
	- Second Approach -
		- o(nlog(k))
* SpaceComplexity
	- O(1)
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2nd Approach - Merge a pair of ListNode in first iteration and smililary in the next iteration
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	l := len(lists)

	increment := 1

	for increment < l {
		for i := 0; i < l-increment; i += 2 * increment {
			lists[i] = sortLists(lists[i], lists[i+increment])
		}
		increment *= 2
	}

	return lists[0]
}

/* First Approach - Mege two LinkedList at a time
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
}*/

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
