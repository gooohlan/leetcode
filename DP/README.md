## 动态规划

动态规划是一个运筹学上的一个最优化算法，所以**动态规划问题一把都是求最值**，比如让你求最长递增子序列，最小编辑距离等。

这类求最值的问题，他的核心是什么？**其核心一定是穷举**。因为要求最值，所以肯定要把所有的答案穷举出来，在其中找出最值。

虽然穷举听起来很简单，但是也不能小看他。动态规划问题一般都比较复杂，需要你熟练掌握递归。动态规划的三要素 **「重叠子问题」**，*
*「状态转移方程」**，**「最优子结构」**，其中 **「状态转移方程」**最为关键，只有写出正确的**「状态转移方程」**，才可以正确的穷举。*
*「重叠子问题」**与**「最优子结构」**可以算是特性，我发现这个问题存在**「重叠子问题」**且具备**「最优子结构」**
，那么这就是一个动态规划问题。动态规划存在 **「重叠子问题」**，在穷举时效率会很低，我们需要使用「备忘录」或者「DP
table」来优化穷举过程，避免不必要的计算。

这里提供一个思维框架，辅助你思考状态转移方程：

**明确 base case -> 明确「状态」-> 明确「选择」 -> 定义 `dfs` 数组/函数的含义**。

- [322. 零钱兑换](https://github.com/gooohlan/leetcode/blob/master/DP/322.go)
- [509. 菲波那切数列](https://github.com/gooohlan/leetcode/blob/master/DP/509.go)
- [192. 打家劫舍](https://github.com/gooohlan/leetcode/blob/master/DP/198.go)
- [213. 打家劫舍 II](https://github.com/gooohlan/leetcode/blob/master/DP/213.go)
- [337. 打家劫舍 III](https://github.com/gooohlan/leetcode/blob/master/DP/337.go)
- [410. 分割数组的最大值](https://github.com/gooohlan/leetcode/blob/master/DP/410.go)
- [300. 最长递增子序列](https://github.com/gooohlan/leetcode/blob/master/DP/300.go)
- [354. 俄罗斯套娃信封问题](https://github.com/gooohlan/leetcode/blob/master/DP/354.go)
- [72. 编辑距离](https://github.com/gooohlan/leetcode/blob/master/DP/72.go)
- [931. 下降路径最小和](https://github.com/gooohlan/leetcode/blob/master/DP/931.go)
- [53. 最大子数组和](https://github.com/gooohlan/leetcode/blob/master/DP/53.go)
- [1143. 最长公共子序列](https://github.com/gooohlan/leetcode/blob/master/DP/1143.go)
- [583. 两个字符串的删除操作](https://github.com/gooohlan/leetcode/blob/master/DP/583.go)
- [712. 两个字符串的最小ASCII删除和](https://github.com/gooohlan/leetcode/blob/master/DP/712.go)
- [516. 最长回文子序列](https://github.com/gooohlan/leetcode/blob/master/DP/516.go)
- [1312. 让字符串成为回文串的最少插入次数](https://github.com/gooohlan/leetcode/blob/master/DP/1312.go)
- [416. 分割等和子集](https://github.com/gooohlan/leetcode/blob/master/DP/416.go)
- [518. 零钱兑换 II](https://github.com/gooohlan/leetcode/blob/master/DP/518.go)
- [494. 目标和](https://github.com/gooohlan/leetcode/blob/master/DP/494.go)
- [64. 最小路径和](https://github.com/gooohlan/leetcode/blob/master/DP/64.go)
- [174. 地下城游戏](https://github.com/gooohlan/leetcode/blob/master/DP/174.go)
- [514. 自由之路](https://github.com/gooohlan/leetcode/blob/master/DP/514.go)
- [787. K 站中转内最便宜的航班](https://github.com/gooohlan/leetcode/blob/master/DP/787.go)
- [10. 正则表达式匹配](https://github.com/gooohlan/leetcode/blob/master/DP/10.go)
- [887. 鸡蛋掉落](https://github.com/gooohlan/leetcode/blob/master/DP/887.go)
- [312. 戳气球](https://github.com/gooohlan/leetcode/blob/master/DP/312.go)
- [486. 预测赢家](https://github.com/gooohlan/leetcode/blob/master/DP/486.go)
- [877. 石子游戏](https://github.com/gooohlan/leetcode/blob/master/DP/887.go)
- [651. 4键键盘](https://github.com/gooohlan/leetcode/blob/master/DP/651.go)
- [28. 找出字符串中第一个匹配项的下标](https://github.com/gooohlan/leetcode/blob/master/DP/28.go)
- [435. 无重叠区间](https://github.com/gooohlan/leetcode/blob/master/DP/435.go)
- [452. 用最少数量的箭引爆气球](https://github.com/gooohlan/leetcode/blob/master/DP/452.go)
- [253. 会议室 II](https://github.com/gooohlan/leetcode/blob/master/DP/253.go)
- [1024. 视频拼接](https://github.com/gooohlan/leetcode/blob/master/DP/1024.go)
- [55. 跳跃游戏](https://github.com/gooohlan/leetcode/blob/master/DP/55.go)
- [45. 跳跃游戏 II](https://github.com/gooohlan/leetcode/blob/master/DP/45.go)

### [1. 使用动态规划解决股票问题](https://github.com/gooohlan/leetcode/tree/master/DP/Stock)

- [121. 买卖股票的最佳时机](https://github.com/gooohlan/leetcode/blob/master/DP/Stock/121.go)
- [122. 买卖股票的最佳时机 II](https://github.com/gooohlan/leetcode/blob/master/DP/Stock/122.go)
- [309. 最佳买卖股票时机含冷冻期](https://github.com/gooohlan/leetcode/blob/master/DP/Stock/309.go)
- [714. 买卖股票的最佳时机含手续费](https://github.com/gooohlan/leetcode/blob/master/DP/Stock/714.go)
- [123. 买卖股票的最佳时机 III](https://github.com/gooohlan/leetcode/blob/master/DP/Stock/123.go)
- [188. 买卖股票的最佳时机 IV](https://github.com/gooohlan/leetcode/blob/master/DP/Stock/188.go)