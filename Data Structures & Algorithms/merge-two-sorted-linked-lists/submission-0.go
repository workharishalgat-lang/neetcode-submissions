/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

   out :=  &ListNode{}
   head :=out

   for list1 != nil || list2!= nil {
	  

	  if list1 != nil  && list2!= nil {
	  if list1.Val <  list2.Val {
		// then create a new node
		node := &ListNode{Val:list1.Val}
		out.Next = node
		out = node
		list1 = list1.Next
	  }else if list2.Val < list1.Val {
		node := &ListNode{Val:list2.Val}
		out.Next = node
		out = node
		list2 = list2.Next
	  }else{
		node1 := &ListNode{Val:list2.Val}
		out.Next = node1
		out =  node1
		node2 := &ListNode{Val:list2.Val}
		out.Next = node2
		out =  node2
		list1 =list1.Next
		list2 = list2.Next
	  }
	  }else if list1==nil {
	out.Next = list2
	break;
   } else {
	out.Next = list1
	break;
   }
   }
   return head.Next  

	
}
