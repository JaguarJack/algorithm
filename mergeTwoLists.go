package main

func main()  {

}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
		pre := &ListNode{}

		result := pre

		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				pre.Next = l1
				l1 = l1.Next
			} else {
				pre.Next = l2
				l2.Next = l2
			}

			pre = pre.Next
		}

		if l1 != nil {
			pre.Next = l1
		}

		if l2 != nil {
			pre.Next = l2
		}

		return result.Next
}