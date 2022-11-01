## 二叉树

> 二叉树解题的思维模式分两类：
> 
> 1、是否可以通过遍历一遍二叉树得到答案？如果可以，用一个 traverse 函数配合外部变量来实现，这叫「遍历」的思维模式。
> 
> 2、是否可以定义一个递归函数，通过子问题（子树）的答案推导出原问题的答案？如果可以，写出这个递归函数的定义，并充分利用这个函数的返回值，这叫「分解问题」的思维模式。
> 
> 无论使用哪种思维模式，你都需要思考：
>
> 如果单独抽出一个二叉树节点，它需要做什么事情？需要在什么时候（前/中/后序位置）做？其他的节点不用你操心，递归函数会帮你在所有节点上执行相同的操作。
>
> 引用自[东哥带你刷二叉树（纲领篇）](https://labuladong.github.io/algo/1/6/)


- [104. 二叉树的最大深度](https://github.com/gooohlan/leetcode/blob/master/BinaryTree/104.go)
- [116. 填充每个节点的下一个右侧节点指针](https://github.com/gooohlan/leetcode/blob/master/BinaryTree/116.go)
- [144. 二叉树的前序遍历](https://github.com/gooohlan/leetcode/blob/master/BinaryTree/144.go)
- [145. 二叉树的后序遍历](https://github.com/gooohlan/leetcode/blob/master/BinaryTree/145.go)
- [226. 翻转二叉树](https://github.com/gooohlan/leetcode/blob/master/BinaryTree/226.go)
- [543. 二叉树的直径](https://github.com/gooohlan/leetcode/blob/master/BinaryTree/543.go)
- [114. 二叉树展开为链表](https://github.com/gooohlan/leetcode/blob/master/BinaryTree/114.go)
- [654. 最大二叉树](https://github.com/gooohlan/leetcode/blob/master/BinaryTree/654.go)
- [105. 从前序与中序遍历序列构造二叉树](https://github.com/gooohlan/leetcode/blob/master/BinaryTree/105.go)
- [106. 从中序与后序遍历序列构造二叉树](https://github.com/gooohlan/leetcode/blob/master/BinaryTree/106.go)
- [889. 根据前序和后序遍历构造二叉树](https://github.com/gooohlan/leetcode/blob/master/BinaryTree/889.go)