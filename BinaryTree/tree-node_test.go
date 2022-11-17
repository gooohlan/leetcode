package BinaryTree

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func deserializeLevelOrder(data string) *TreeNode {
	list := strings.Split(data, ",")
	val, err := strconv.Atoi(list[0])
	if err != nil {
		return nil
	}
	root := &TreeNode{Val: val}

	q := []*TreeNode{root}
	for i := 1; i < len(list)-1; {
		// 队列中存的都是父节点
		parent := q[0]
		q = q[1:]
		// 父节点对应的左侧子节点的值
		if list[i] != "null" {
			val, _ := strconv.Atoi(list[i])
			parent.Left = &TreeNode{Val: val}
			q = append(q, parent.Left)
		} else {
			parent.Left = nil
		}
		i++

		// 父节点对应的右侧子节点的值
		if list[i] != "null" {
			val, _ := strconv.Atoi(list[i])
			parent.Right = &TreeNode{Val: val}
			q = append(q, parent.Right)
		} else {
			parent.Right = nil
		}
		i++
	}
	return root
}

func TestMaxDepth(t *testing.T) {
	node := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: &TreeNode{2, nil, nil},
	}
	fmt.Println(maxDepth1(node))
	fmt.Println(maxDepth2(node))
}

func TestPreorderTraversal(t *testing.T) {
	node := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, &TreeNode{Val: 5}}},
		Right: &TreeNode{9, nil, nil},
	}
	fmt.Println(preorderTraversal1(node))
	// fmt.Println(preorderTraversal2(node))
}
func TestDiameterOfBinaryTree(t *testing.T) {
	node := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, &TreeNode{Val: 5}}},
		Right: &TreeNode{9, nil, nil},
	}
	fmt.Println(diameterOfBinaryTree(node))
}

func TestInvertTree(t *testing.T) {
	node := &TreeNode{
		Val:   4,
		Left:  &TreeNode{2, &TreeNode{Val: 1}, &TreeNode{Val: 3}},
		Right: &TreeNode{7, &TreeNode{Val: 6}, &TreeNode{Val: 9}},
	}
	node = invertTree1(node)
	node = invertTree2(node)

	fmt.Println(111)
}

func TestConnect(t *testing.T) {
	node := &Node{
		1,
		&Node{
			2,
			&Node{4, &Node{8, nil, nil, nil}, &Node{9, nil, nil, nil}, nil},
			&Node{5, &Node{10, nil, nil, nil}, &Node{11, nil, nil, nil}, nil},
			nil,
		},
		&Node{
			3,
			&Node{6, &Node{12, nil, nil, nil}, &Node{13, nil, nil, nil}, nil},
			&Node{7, &Node{14, nil, nil, nil}, &Node{15, nil, nil, nil}, nil},
			nil,
		},
		nil,
	}

	connect(node)
}

func Test114(t *testing.T) {
	node := &TreeNode{
		1,
		&TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{5, &TreeNode{6, nil, nil}, nil},
	}

	flatten(node)
	fmt.Println(node)
}

func Test654(t *testing.T) {
	nums := []int{3, 2, 1, 6, 0, 5}
	node := constructMaximumBinaryTree2(nums)
	fmt.Println(node)
}

func Test105(t *testing.T) {
	preorder := []int{1, 2}
	inorder := []int{2, 1}
	node := buildTree(preorder, inorder)
	fmt.Println(node)
}

func Test106(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	node := buildTree106(inorder, postorder)
	fmt.Println(node)
}

func Test889(t *testing.T) {
	preorder := []int{1, 2, 4, 5, 3, 6, 7}
	postorder := []int{4, 5, 2, 6, 7, 3, 1}
	node := constructFromPrePost(preorder, postorder)
	fmt.Println(node)
}

func Test297(t *testing.T) {
	// node := &TreeNode{}
	ser := Constructor()
	deser := Constructor()
	data := ser.serializeLevelOrder(nil)
	ans := deser.deserializeLevelOrder(data)
	fmt.Println(data)
	fmt.Println(ans)
}

func Test102(t *testing.T) {
	node := &TreeNode{
		1,
		&TreeNode{2, nil, &TreeNode{4, nil, nil}},
		&TreeNode{3, nil, nil},
	}

	fmt.Println(levelOrder(node))
}

func Test652(t *testing.T) {
	l := &TreeNode{4, nil, nil}
	ll := &TreeNode{2, l, nil}
	node := &TreeNode{
		Val:   1,
		Left:  ll,
		Right: &TreeNode{3, ll, &TreeNode{4, nil, nil}},
	}
	subtrees := findDuplicateSubtrees(node)
	fmt.Println(subtrees)
}

func Test236(t *testing.T) {
	p := deserializeLevelOrder("2,null")
	q := deserializeLevelOrder("1,4,null")

	node := &TreeNode{
		Val:   1,
		Left:  &TreeNode{2, p, nil},
		Right: &TreeNode{3, nil, q},
	}
	fmt.Println(lowestCommonAncestor(node, p, q))
}
