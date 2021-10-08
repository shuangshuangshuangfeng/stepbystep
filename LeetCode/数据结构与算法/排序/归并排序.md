## 归并排序(merge sort)

归并排序属于Divide and Conquer(分而治之)的一种算法。

**分治法:**<br>
分治法是一种算法思想，它将问题的实例划分为若干个较小的实例(最好拥有相同的规模)，对这些小的实例进行递归求解，然后合并这些解，得到原问题的解。

例如一个数字序列``A=[3,4,1,6,7,2,5,9]``， 我们的**目标是：**``Sort(A)``。

![gongshi](https://github.com/shuangshuangshuangfeng/daguaishengji/blob/master/nlp/passage1/note/gongshi1.png?raw=true) <br>

**注意：** 上图中的单词conquer写成merge,懒得改了。<br>

- 首先将``A``序列``devide``为两个子序列``A1=[3, 4, 1, 6]``和``A2=[7, 2, 5, 9]``
- 然后将子问题求解得到``Sort(A1)=[1, 3, 4, 6]``和``Sort(A2)=[2, 5, 7, 9]``。<br>
- 最后，将子问题的结果merge(合并):1和2比较，1小于2，A1指针移到3， 然后3和2比较，3大于2，A2指针移到5...，最后两个子串结果合并为``Sort(A)=[1, 2, 3, 4, 5, 6, 7, 9]``。

**时间复杂度:**<br>
如果数列中有``n``个数，那么原始问题的时间为两个子问题的时间加上两个子串合并的时间(合并时间是做了n次比较)
，因此``T(n) = T(n/2)+T(n/2)+n``。<br>

该算法的时间复杂度实际是是O(NlogN),如何得到的该复杂度，需要到 <a href="https://github.com/shuangshuangshuangfeng/daguaishengji/blob/master/nlp/passage1/note/master_theorem.md">主定理方法</a>中了解。

**空间复杂度:**<br>
该算法的空间复杂度为O(1)。空间复杂度也就是**递归的空间复杂度**，在 <a href="/nlp/passage1/note/master_theorem.md">斐波那契数列</a>中有讲。



----------------------
**参考：**
1. https://blog.csdn.net/huanghanqian/article/details/78828788




