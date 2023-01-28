## 最小栈

``leetcode``地址:[https://leetcode-cn.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/](https://leetcode-cn.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/)

### 题目描述

给你一个整数数组 arr 和两个整数 k 和 threshold 。

请你返回长度为 k 且平均值大于等于 threshold 的子数组数目。

**示例:**

``` bash
输入：arr = [11,13,17,23,29,31,7,5,2,3], k = 3, threshold = 5
输出：6
解释：前 6 个长度为 3 的子数组平均值都大于 5 。注意平均值不是整数。
```

### 解答

#### 1. 暴力法

```java
public int numOfSubarrays_1(int[] arr, int k, int threshold) {
    int count = 0;
    for(int i=k-1; i<arr.length; i++){
        // int sum = arr[i]+arr[i-1]+arr[i-2];
        int sum = 0;
        int m = 0;
        while(m<k){
            sum+=arr[i-m];
            m++;
        }
        if(sum/k >= threshold) count ++;
    }
    return count;
}

```

- 时间复杂度为O(N^2)
- 空间复杂度为O(1)

#### 2. 滑动窗口

```java

public int numOfSubarrays(int[] arr, int k, int threshold) {
    if(arr.length < k) return 0;

    int count = 0;
    int sum = 0;
    for(int i=0; i<k; i++){
        sum += arr[i];
    }
    if(sum >= threshold*k) count++;

    for(int i=k; i<arr.length; i++){
        // 滑动窗口
        sum += arr[i];
        sum -= arr[i-k];
        if(sum >= threshold*k) count++;
    }
    return count;
}

```

- 时间复杂度O(N)
- 空间复杂度O(1)
