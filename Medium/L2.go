package Medium

/*
* Approach -
		- Add the number from the left-right side
		- If carry is there add the new struct after the last element
		- Honour carry everytime we do addition
		- If the one LinkedList is travered fully, 	then handle the carry
		- In this soln we are using the same linked list
		- create new structure only when we reached at the end of the either of the linked list

* Example - 9 9 9 9 9
			9 9 9
			----------
			8 9 9 0 0 1

* TimeComplexity  - O(n)
* SpaceComplexity - O(len(l1)-len(l2))

* Tags - LinkedList
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	l3 := l2
	newHead := l2

	for l1 != nil && l2 != nil {

		sum := l1.Val + l2.Val + carry
		carry = sum / 10

		l3.Val = sum % 10
		l1 = l1.Next
		l2 = l2.Next
		if l3.Next != nil {
			l3 = l3.Next
		}
	}

	if l2 == nil && l1 != nil {
		for l1 != nil {
			sum := l1.Val + carry
			carry = sum / 10

			l3.Next = &ListNode{
				Val:  sum % 10,
				Next: nil,
			}
			l1 = l1.Next
			l3 = l3.Next
		}
	} else if l1 == nil && l2 != nil {
		for l2 != nil {
			sum := l2.Val + carry
			carry = sum / 10

			l3.Val = sum % 10
			l2 = l2.Next
			if l3.Next != nil {
				l3 = l3.Next
			}
		}
	}

	if carry == 1 {
		l3.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return newHead
}
