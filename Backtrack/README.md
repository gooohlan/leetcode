## 回溯

回溯算法

### 排列/组合/子集问题
无论是排列、组合还是子集问题，简单说无非就是让你从序列 arr 中以给定规则取若干元素，主要有以下几种变体

- 元素无重不可复选，即`arr`中的元素都是唯一的，每个元素最多只能被使用一次，这也是最基本的形式
  以组合为例，如果输入`arr = [2,3,6,7]`，和为`7`的组合应该只有`[7]`
- 元素可重不可复选，即`arr`中的元素可以存在重复，每个元素最多只能被使用一次。
  以组合为例，如果输入`arr = [2,5,2,1,2]`，和为`7`的组合应该有两种`[2,2,2,1]`和`[5,2]`
- 元素无重可复选，即 arr 中的元素都是唯一的，每个元素可以被使用若干次
  以组合为例，如果输入`arr = [2,3,6,7]`，和为`7`的组合应该有两种`[2,2,3]`和`[7]`

其实还有第四种形式，即元素可重可复选.但既然元素可复选，那又何必存在重复元素呢？元素去重之后就等同于元素无重可复选，所以这种情况不用考虑

上面举例使用的事组合，除去组合还有排列和子集问题，他们也有上述三种基本形式，所以共9种变化
除了这些基本变化以为，题目中还会添加各种限制条件，比如限制和为`7`但是元素个数为`3`的组合，这样一来又可以衍生出一堆变体

但万变不离其宗，其本质依旧是穷举所有解，而这些解呈现树形结构，所以合理利用回溯算法框架，稍微改下代码就可以把这些问题一网打尽

回溯算法框架:
```go
var res []int
func bacjtrack(路径, 选择列表):
    if 满足条件 {
        res = append(res, 路径)
        return
    }
    for _, 选择 := range 选择列表 {
        if xxx {
            剪枝			
        }
        做选择
        bacjtrack(路径, 选择列表)
        撤销选择	
    }
}
```
整个框架的核心就是`for`循环里面的递归，在递归前「做选择」，在递归完之后「撤销选择」，所有的变化都可以套用这个框架，上述的三种变化形式，无非是在做一些剪枝操作。

开始之前，我们先记住两棵树
![](https://labuladong.github.io/algo/images/%e6%8e%92%e5%88%97%e7%bb%84%e5%90%88/1.jpeg)
![](https://labuladong.github.io/algo/images/%e6%8e%92%e5%88%97%e7%bb%84%e5%90%88/2.jpeg)
组合问题和子集问题其实是等价的，这两棵树对应上面的算法框架，上述的三种变化形式和其他的限制条件，都是在树上进行剪枝

- 子集（元素无重不可复选）
  - [78. 子集](https://github.com/gooohlan/leetcode/blob/master/Backtrack/78.go)
- 组合（元素无重不可复选）

  **大小为`k`的组合就是大小为`k`的子集**
  - [77. 组合](https://github.com/gooohlan/leetcode/blob/master/Backtrack/77.go)
  - [216.组合总和 III](https://github.com/gooohlan/leetcode/blob/master/Backtrack/216.go)
  - [698. 划分为k个相等的子集](https://github.com/gooohlan/leetcode/blob/master/Backtrack/698.go)
- 排列（元素无重不可复选）
  - [46. 全排列](https://github.com/gooohlan/leetcode/blob/master/Backtrack/46.go)
  - [51. N 皇后（特殊变种）](https://github.com/gooohlan/leetcode/blob/master/Backtrack/51.go)
  - [52. N 皇后 II（特殊变种）](https://github.com/gooohlan/leetcode/blob/master/Backtrack/52.go)
- 子集/组合（元素可重不可复选）
  - [90. 子集 II](https://github.com/gooohlan/leetcode/blob/master/Backtrack/90.go)
  - [40. 组合总和 II](https://github.com/gooohlan/leetcode/blob/master/Backtrack/40.go)
- 排列（元素可重不可复选）
  - [47. 全排列 II](https://github.com/gooohlan/leetcode/blob/master/Backtrack/47.go)
- 子集/组合（元素无重可复选）
  - [39. 组合总和](https://github.com/gooohlan/leetcode/blob/master/Backtrack/39.go)
- 排列（元素无重可复选）

  力扣上没有相关题目，所以我们根据情况假设一下，`arr`数组中的元素无重复客复选会互的情况,会有那些排列

  比如输入`arr = [1, 2]`，那么这种情况下的全排列共有`2^2 = 4`种:
  ```
  [
    [1,1], [1,2],
    [2,1], [2,2],
  ]
  ```
  46题的全排列使用`used`进行剪枝，避免使用重复元素，允许使用的话，去掉所有`used`不进行剪枝即可，代码就不做赘述了。