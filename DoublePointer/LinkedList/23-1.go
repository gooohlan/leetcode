package linked_list

func mergeKLists1(lists []*ListNode) *ListNode {
	pre := &ListNode{}
	cur := &ListNode{}
	if len(lists) == 0 {
		return nil
	}
	pre = lists[0]
	for i := 1; i < len(lists); i++ {
		cur = lists[i]
		pre = mergeTwoLists(pre, cur)
	}
	return pre
}
