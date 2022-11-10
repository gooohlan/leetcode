package BinarySearchTree

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

func Test230(t *testing.T) {
	node := deserializeLevelOrder("3,1,4,-,2")
	fmt.Println(kthSmallest2(node, 3))
}

func Test538(t *testing.T) {
	root := deserializeLevelOrder("4,1,6,0,2,5,7,null,null,null,3,null,null,null,8")
	convertBST(root)
}

func Test98(t *testing.T) {
	root := deserializeLevelOrder("3,1,5,0,2,4,6")
	fmt.Println(isValidBST(root))
}

func Test700(t *testing.T) {
	root := deserializeLevelOrder("4,2,7,1,3")
	fmt.Println(searchBST(root, 2))
}
