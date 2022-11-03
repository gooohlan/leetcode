package BinaryTree

import (
	"bytes"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// 前序遍历

// 序列化与之前的做法一样,不同的是遇到空树加入特殊标记
func (c *Codec) serializePreorder(root *TreeNode) string {
	var buffer bytes.Buffer
	var dfs func(*TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			buffer.WriteString("-,")
			return
		}
		buffer.WriteString(strconv.Itoa(node.Val))
		buffer.WriteString(",")
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return buffer.String()
}

// 正常的前序遍历是无法恢复为二叉树的,但是这里我们存储了空指针信息
// 把字符串切割为数组,循环遍历,遇到空指针标记则表示当前为空树,直接返回
// 否则继续解析这棵树的左子树,然后是右子树
func (c *Codec) deserializePreorder(data string) *TreeNode {
	list := strings.Split(data, ",")

	var build func() *TreeNode
	build = func() *TreeNode {
		if list[0] == "-" { // 空指针直接返回
			list = list[1:]
			return nil
		}
		val, _ := strconv.Atoi(list[0])
		list = list[1:]
		root := &TreeNode{Val: val}
		root.Left = build() // 先解析左子树,左子树遇到空指针时表示左子树解析结束,然后开始解析右子树
		root.Right = build()
		return root
	}
	return build()
}

// 后续遍历
// 交换一下写入当前节点位置即可
func (c *Codec) serializePostorder(root *TreeNode) string {
	var buffer bytes.Buffer
	var dfs func(*TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			buffer.WriteString("-,")
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		buffer.WriteString(strconv.Itoa(node.Val))
		buffer.WriteString(",")
	}

	dfs(root)
	return buffer.String()
}

// 与前序解析类似,不同的是需要从后往前找,并且是先构造右子树
func (c *Codec) deserializePostorder(data string) *TreeNode {
	list := strings.Split(data, ",")
	n := len(list) - 2 // 构造时会多出一个,解析为数组后末尾会多一个空字符串
	var build func() *TreeNode
	build = func() *TreeNode {
		if list[n] == "-" || n == 0 { // 空指针直接返回
			n--
			return nil
		}
		val, _ := strconv.Atoi(list[n])
		n--
		root := &TreeNode{Val: val}
		root.Right = build() // 先解析右子树,右子树遇到空指针时表示右子树解析结束,然后开始解析左子树
		root.Left = build()
		return root
	}
	return build()
}
