/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return head
	}
	// 3

    var prev *ListNode

	var curr *ListNode = head
	var next *ListNode = curr.Next

	for curr != nil {		
		curr.Next = prev
		prev = curr
		curr =  next
		if curr == nil {
			break
		}
		next = curr.Next
	}

	return prev



}
