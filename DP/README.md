## 动态规划

动态规划是一个运筹学上的一个最优化算法，所以**动态规划问题一把都是求最值**，比如让你求最长递增子序列，最小编辑距离等。

这类求最值的问题，他的核心是什么？**其核心一定是穷举**。因为要求最值，所以肯定要把所有的答案穷举出来，在其中找出最值。

虽然穷举听起来很简单，但是也不能小看他。动态规划问题一般都比较复杂，需要你熟练掌握递归。动态规划的三要素 **「重叠子问题」**，**「状态转移方程」**，**「最优子结构」**，其中 **「状态转移方程」**最为关键，只有写出正确的**「状态转移方程」**，才可以正确的穷举。**「重叠子问题」**与**「最优子结构」**可以算是特性，我发现这个问题存在**「重叠子问题」**且具备**「最优子结构」**，那么这就是一个动态规划问题。动态规划存在 **「重叠子问题」**，在穷举时效率会很低，我们需要使用「备忘录」或者「DP table」来优化穷举过程，避免不必要的计算。

这里提供一个思维框架，辅助你思考状态转移方程：

**明确 base case -> 明确「状态」-> 明确「选择」 -> 定义 `dp` 数组/函数的含义**。

- [322. 零钱兑换](https://github.com/gooohlan/leetcode/blob/master/DP/322.go)
- [509. 菲波那切数列](https://github.com/gooohlan/leetcode/blob/master/DP/509.go)

### 使用动态规划解决股票问题

[Stock](https://github.com/gooohlan/leetcode/tree/master/DP/Stock)
- [121. 买卖股票的最佳时机](https://github.com/gooohlan/leetcode/blob/master/DP/Stock/121.go)
- [122. 买卖股票的最佳时机 II](https://github.com/gooohlan/leetcode/blob/master/DP/Stock/122.go)
- [309. 最佳买卖股票时机含冷冻期](https://github.com/gooohlan/leetcode/blob/master/DP/Stock/309.go)
- [714. 买卖股票的最佳时机含手续费](https://github.com/gooohlan/leetcode/blob/master/DP/Stock/714.go)
- [123. 买卖股票的最佳时机 III](https://github.com/gooohlan/leetcode/blob/master/DP/Stock/123.go)