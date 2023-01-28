## 斐波那契数列

斐波那契数列为``1, 1, 2, 3, 5, 8, 13, 21, ...``<br>
**问题:** 怎样求出序列中第N个数？？ (上帝密码)
```
def fib(n):
    # base case
    if n<3: 
        return 1
    
    return fib(n-2)+fib(n-1)

```

**时间复杂度:** 为 ``2^n``， （根据一棵树）<br>
![gongshi](https://github.com/shuangshuangshuangfeng/daguaishengji/blob/master/nlp/passage1/note/gongshi4.png?raw=true) <br>


因此需要知道这棵树到底有多深？？
```
n = 8 --> h=6
n = ... --> h=n-2
```
所以时间复杂度依赖于``2^n``，为``O(2^n)``。


### 空间复杂度
主要是为了介绍**递归方式下的空间复杂度** 。<br>
当一个函数调用另一个函数的适合，要进行上下文的切换，每次上下文切换都需要有内存空间的使用，我们称使用一个单位内存空间。
![gongshi](https://github.com/shuangshuangshuangfeng/daguaishengji/blob/master/nlp/passage1/note/gongshi5.png?raw=true) <br>
``fib(8)``最多占用8个单位内存空间，因此它的空间复杂度为``O(1)``。

### 使用循环表示数列
使用数组来存实例,这种思路叫``dynamic programming``动态规划，就是能复用的尽量去复用。
```
def fib(n):
    temp = np.zeros(n)
    temp[0] = 1
    temp[1] = 1
    for i in range(2, n):
        temp[i] = temp[i-2]+temp[i-1]

    return temp[n-1]
```
时间复杂度为``O(N)``，但是算法中，第``n``个数中仅仅需要前两个数字，所以额外的占用了一些其他内存空间，仅仅维护``n``前面两个数字就可以了。


---------------------------
**参考:**<br>
1. https://blog.csdn.net/liu_zhen_kai/article/details/82630060



