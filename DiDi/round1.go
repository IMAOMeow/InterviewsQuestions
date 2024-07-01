// 以下均为面试时所写，未经运行测试，仅供参考

// 单链表翻转，
// 1->2->3->4->5->6->7->8,
// 8->7->6->5->4->3->2->1

type ListNode struct {
	Next *ListNode
	Val  int
}

func reverse(head *ListNode) *ListNode {
	// input validation
	if head == nil {
		return nil
	}

	dummy := &ListNode{}

	p := head

	for p != nil {
		// insert to dummy
		temp := p.Next
		p.Next = dummy.Next
		dummy.Next = p

		p = temp
	}

	return dummy.Next
}

// 每k个节点进行一次反转
// 1->2->3->4->5->6->7->8,k=3
// 3->2->1->6->5->4->7->8

func reverseK(head *ListNode, k int) *ListNode {
	// input validation
	if head == nil || k == 0 {
		return head
	}

	p := head
	length := 0

	for p != nil {
		length++
		p = p.Next
	}

	if length < k {
		return head
	}

	dummy := &ListNode{}
	p = head
	pre := head
	count := 0
	for count < k {
		temp := p.Next
		p.Next = dummy.Next
		dummy.Next = p

		pre = p
		p = temp

		count++
	}

	q := reverseK(p, k)
	pre.Next = q

	return dummy.Next
}
