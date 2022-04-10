package Medium

/*
	Algorithm:
		- For Method 1, it is pretty straight forward
			- calcaulate the length first
			- And then find out the index to be deleted
		- For Method 2, it is little tricky to come up with
			- We will use SlowFast pointer stratesgy to solve this
			- The idea is first move the fast pointer
			- This will make sure that the fast pointer is reached the differenc of the elem
			- Now move the slow pointer and fast pointer untill fast pointer reaches end of the linked list
			- Slow pointer has reached to the required position
			- IMPORTANT point here is the corener case, where we have want to delete the first element.
	TimeComplexity:
		- Method 1
			- o(n) - Two Pass
		- Method 2
			- o(n) - one pass
	SpaceComplexity
		- o(1)
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

//=====>>> Method 1 <<<<========

/*func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := head

	l := calculate(temp)

	deleteIndex := l - n

	temp = head

	for deleteIndex > 1 {
		temp = temp.Next
		deleteIndex--
	}

	temp.Next = temp.Next.Next

	return head
}

func calculate(head *ListNode) int {
	len := 0

	for head != nil {
		len++
		head = head.Next
	}

	return len
}*/

//====>>>>> Method 2  <<<<=======

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//var slow, fast *ListNode
	slow, fast := head, head

	for i := 1; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// This is the most important corner case, it means if we want to delete the first elem. Will fail when we have only one elem
	if fast == nil {
		return head.Next
	}

	slow.Next = slow.Next.Next

	return head
}
