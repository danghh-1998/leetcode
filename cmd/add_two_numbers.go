package main

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1:
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.

Example 2:
Input: l1 = [0], l2 = [0]
Output: [0]

Example 3:
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]

Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sumList := &ListNode{}
	sumItem := sumList
	l1Item := l1
	l2Item := l2
	for {
		sum := 0
		if l1Item != nil {
			sum += l1Item.Val
		}
		if l2Item != nil {
			sum += l2Item.Val
		}

		if l1Item != nil {
			l1Item = l1Item.Next
		}
		if l2Item != nil {
			l2Item = l2Item.Next
		}
		if sum >= 10 {
			if l1Item != nil {
				l1Item.Val += 1
			} else if l2Item != nil {
				l2Item.Val += 1
			} else {
				l1Item = &ListNode{Val: 1}
			}
		}
		sumItem.Val = sum % 10
		if l1Item == nil && l2Item == nil {
			return sumList
		}
		sumItem.Next = &ListNode{}
		sumItem = sumItem.Next
	}
}
