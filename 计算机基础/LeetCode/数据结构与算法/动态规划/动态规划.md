
## 动态规划（Dynamic Programming, DP）

**动态规划:**动态规划回保存以前的运算结果，并根据以前的结果对当前进行选择，有回退功能。


动态规划的一个特点是：把一个大问题拆成一个子问题，一点点的解决，最终把原问题解决。**如何分解子问题是最重要的。**

当我们拿到一个动态规划题目时，必须要做三件事：
- 1.问题目标
- 2.状态定义
- 3.状态转移方程

-------------------------------------------------------------
## 题目

### 1. 最大子序和
``LeetCode``链接: [https://leetcode-cn.com/problems/maximum-subarray/](https://leetcode-cn.com/problems/maximum-subarray/ "123")

**题目描述:**<br/>
给定一个整数数组 A ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

**示例:**<br/>
```
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
```

**解题:**<br/>

1. **目标**: 给定一个数组，求里面子串中的最大和
2. **状态**: 定义一个数组``M``，数组上的第``j``个元素用来表示截止到第``j``个元素，``A[0-j]``元素的子串的最大值为``M[j]``
3. **状态转移方程**: ``M(j) = max{M[j-1]+A[j], A[j]}``

```java
public int func(int[] A){
    int[] M = new int[A.length];
    M[0] = A[0];
    for(int i=1; i<A.length; i++){
        M[i] = Math.max(M[i-1]+A[i], A[i]);
    }
    return M[A.length-1];
}

```

### 2. 打家劫舍
``LeetCode``链接: [https://leetcode-cn.com/problems/house-robber/](https://leetcode-cn.com/problems/house-robber/ 
 "123")


**题目描述:**<br/>
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。

**示例:**<br/>
```
输入：[1,2,3,1]
输出：4
解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
```

**解题:**<br/>

1. **目标**: 给定一个数组A，求数组中不连续元素的最大和
2. **状态**: 定义一个数组``M``，数组上的第``j``个元素用来表示截止到第``j``个元素，``A[0-j]``元素的不连续元素和的最大值为``M[j]``
3. **状态转移方程**: ``M(j) = max{M[j-2]+A[j], M[j-1]}``

```java
public int func(int[] A){
    int[] M = new int[A.length];
    M[0] = A[0];
    M[1] = Math.max(A[0], A[1]);
    for(int i=2; i<A.length; i++){
        M[i] = Math.max(M[i-2]+A[i], M[i-1]);
    }
    return M[A.length-1];
}
```

### 3.最长连续递增序列
``LeetCode``链接: [https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/](https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/  "123")


**题目描述:**<br/>
给定一个未经排序的整数数组，找到最长且连续的的递增序列，并返回该序列的长度。 

**示例:**<br/>
```
输入: [1,3,5,4,7]
输出: 3
解释: 最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为5和7在原数组里被4隔开。 
```

```
输入: [2,2,2,2,2]
输出: 1
解释: 最长连续递增序列是 [2], 长度为1。
```

**解题:**<br/>

1. **目标**: 给定一个数组A，求数组中连续严格递增的数列长度。
2. **状态**: 定义一个数组``M``，数组上的第``j``个元素用来表示截止到第``j``个元素，``A[0-j]``元素的严格递增的数列长度的最大值为``M[j]``
3. **状态转移方程**: 定义当前递增起始点为``pre_index``,  ``if(A[j]>A[j-1]): M(j) = max{M[j-1], j-pre_index}``

```java
public int func(int[] A){
    int[] M = new int[A.length];
    M[0] = 1;
    int pre_index = 0;
    for(int i=1; i<A.length; i++){
        if(A[i] > A[i-1]){
            M[i] = Math.max(M[i-1], i-pre_index);
        }else{
            pre_index = i;
            M[i] = Math.max(M[i-1], 1);
        }
    }
    return M[A.length-1];
}
```

### 4. 爬楼梯
``LeetCode``链接: [https://leetcode-cn.com/problems/climbing-stairs/](https://leetcode-cn.com/problems/climbing-stairs/  "123")


**题目描述:**<br/>
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

**示例:**<br/>
```
输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶 
```

```
输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶
```

**解题:**<br/>

1. **目标**: 最多有多少种爬梯方式。
2. **状态**: 定义一个数组``M``，数组上的第``j``个元素用来表示截止到第``j``个元素，``n``阶楼梯的爬梯方式的最大值为``M[j]``
3. **状态转移方程**: 当爬大匹第``n``阶楼梯时，有可能是上一阶楼梯爬一阶，也有可能是上上阶楼梯爬两阶，那么``M[j] = M[j-1]+M[j-2]``.

```java
public int func(int n){
    int[] M = new int[A.length];
    M[0] = 1;
    M[1] = 2;
    for(int i=2; i<n; i++){
        M[i] = M[i-1]+M[i-2];
    }
    return M[n-1];
}
```





















