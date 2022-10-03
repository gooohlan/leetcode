## 滑动窗口

**滑动窗口算法的思路是这样:**

1、我们在字符串 S 中使用双指针中的左右指针技巧，初始化 left = right = 0，把索引左闭右开区间 [left, right) 称为一个「窗口」。

> PS：理论上你可以设计两端都开或者两端都闭的区间，但设计为左闭右开区间是最方便处理的。因为这样初始化 left = right = 0 时区间 [0, 0) 中没有元素，但只要让 right 向右移动（扩大）一位，区间 [0, 1) 就包含一个元素 0 了。如果你设置为两端都开的区间，那么让 right 向右移动一位后开区间 (0, 1) 仍然没有元素；如果你设置为两端都闭的区间，那么初始区间 [0, 0] 就包含了一个元素。这两种情况都会给边界处理带来不必要的麻烦。

2、我们先不断地增加 right 指针扩大窗口 [left, right)，直到窗口中的字符串符合要求（包含了 T 中的所有字符）。

3、此时，我们停止增加 right，转而不断增加 left 指针缩小窗口 [left, right)，直到窗口中的字符串不再符合要求（不包含 T 中的所有字符了）。同时，每次增加 left，我们都要更新一轮结果。

4、重复第 2 和第 3 步，直到 right 到达字符串 S 的尽头。

这个思路其实也不难，第 2 步相当于在寻找一个「可行解」，然后第 3 步在优化这个「可行解」，最终找到最优解，也就是最短的覆盖子串。左右指针轮流前进，窗口大小增增减减，窗口不断向右滑动，这就是「滑动窗口」这个名字的来历。


- [76. 最小覆盖子串](https://github.com/gooohlan/leetcode/blob/master/DoublePointer/SlidingWindow/76.go)
- [567. 字符串的排列](https://github.com/gooohlan/leetcode/blob/master/DoublePointer/SlidingWindow/567.go)
- [438. 找到字符串中所有字母异位词](https://github.com/gooohlan/leetcode/blob/master/DoublePointer/SlidingWindow/438.go)
- [3. 无重复字符的最长子串](https://github.com/gooohlan/leetcode/blob/master/DoublePointer/SlidingWindow/3.go)