package BinaryTree

type Node1650 struct {
	Val    int
	Left   *Node1650
	Right  *Node1650
	Parent *Node1650
}

// 这题乍一很难,但是细想下来,把这个父节点看为下一个节点,这道题就是160
// 这题就变成了寻找单链表相交的问题

func lowestCommonAncestor3(p, q *Node1650) *Node1650 {
	P, Q := p, q
	for P != Q {
		if P == nil { // p走完开始走q
			P = q
		}

		if Q == nil { // q走完开始走p
			Q = p
		}
		P = P.Parent
		Q = Q.Parent
	}

	return P
}
