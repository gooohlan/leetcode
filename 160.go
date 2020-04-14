package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	Pa, Pb := headA,headB
	for Pa != Pb {
		if Pa == nil {
			Pa = headB
			continue
		}
		if Pb == nil {
			Pb = headA
		}
		Pa = Pa.Next
		Pb = Pb.Next
	}
	return Pa
}
