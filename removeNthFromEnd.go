package main

func main()  {
	
}

/**
哨兵节点，其实就是一个附加在原链表最前面用来简化边界条件的附加节点，
它的值域不存储任何东 西，只是为了操作方便而引入。
比如原链表为 a -> b -> c，则加了哨兵节点的链表即为 x -> a -> b > c，

那我们为什么需要引入哨兵节点呢?举个例子，
比如我们要删除某链表的第一个元素，常见的删除链表
的操作是找到要删元素的前一个元素，假如我们记为 pre。我们通过：
	`pre.Next = pre.Next.Next`

来进行删除链表的操作。但是此时若是删除第一个元素的话，你就很难进行了，
因为按道理来讲，此时 第一个元素的前一个元素就是nil（空的），
如果使用pre就会报错。那如果此时你设置了哨兵节点的 话，
此时的pre就是哨兵节点了。这样对于链表中的任何一个元素，
你要删除都可以通过
pre.Next = pre.Next.Next 的方式来进行，这就是哨兵节点的作用。

 */
type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first := &ListNode{}
	first.Next = head

	var pre *ListNode
	cur := first
	i := 1
	for head != nil{
		if i >= n {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}

	pre.Next = pre.Next.Next

	return first.Next
}
