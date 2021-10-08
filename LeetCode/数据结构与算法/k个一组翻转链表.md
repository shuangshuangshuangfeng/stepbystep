## K个一组翻转链表

``leetcode``地址:[https://leetcode-cn.com/problems/reverse-nodes-in-k-group/](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/ )

### 题目描述

给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

**注意**：<br/>

- 你的算法只能使用常数的额外空间。
- 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

**示例:**

```
给你这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5
```

### 解答

#### 1. 使用k个额外空间

```java
public ListNode reverseKGroup(ListNode head, int k) {
    // 使用k个单位的数组作为缓存
    ListNode[] temp = new ListNode[k];

    int cnt = 0; //计数器
    // 反转后链表的头指针
    ListNode result = new ListNode();
    // 反转后链表的尾指针
    ListNode tail = result;

    while(head != null){
        // 先判断不等于K, 然后再判断等于k， 很巧妙
        // 这样当第k个数填完以后可以直接做操作
        // 从而省去了while结束后判断是否cnt=k的操作了
        if(cnt != k){
            temp[cnt] = head;
            head = head.next;
            cnt ++;
        }

        if(cnt == k){
            for(int i=cnt-1; i>=0; i--){
                tail.next = temp[i];
                tail = temp[i];
                // 防止结果出现循环链表
                tail.next = null;
            }
            head = head.next;
            cnt = 0;
        }
    }

    // 剩余部分
    if(cnt < k && cnt != 0){ 
        // 如果小于k，则不反转
        for(int i=0; i<cnt; i++){
            tail.next = temp[i];
            tail = temp[i];
            tail.next = null;
        }
    }

    return result.next;
}
```

- 时间复杂度为O(N)
- 空间复杂度为O(N)

#### 2. 尾插法

使用双指针，将链表分为三个部分：已翻转、待反转、未翻转。
**待翻转区间为前开后闭区间**

```java

public ListNode reverseKGroup(ListNode head, int k) {
        // 反转后链表的头指针
        ListNode result = new ListNode();
        result.next = head;
        // 指向待翻转部分的头指针
        ListNode reverse_head = result;
        // 指向待翻转部分的尾指针
        ListNode reverse_tail = result;
        
        while(true){
            // 计数器
            int cnt = k;
            // 找到待翻转尾部
            while(cnt>0 && reverse_tail!=null){
                cnt--;
                reverse_tail = reverse_tail.next;
            }

            if(reverse_tail == null) break;
            //暂存新head（注意这里要的新head是翻转以后的head）
            ListNode new_head = reverse_head.next;
            
            // 尾插法翻转
            while(reverse_head.next != reverse_tail){
                ListNode temp = reverse_head.next;
                reverse_head.next = temp.next;
                temp.next = reverse_tail.next;
                reverse_tail.next = temp;
            }
            // 重置新的待翻转起点
            reverse_head = new_head;
            reverse_tail = new_head;
            // System.out.println("new head:"+new_head.val);

            // 打印
            // ListNode r = result;
            // while(r != null){
            //     System.out.print(r.val+"-");
            //     r = r.next;
            // }
            // System.out.println();
        }
        return result.next;
    }
```

- 时间复杂度为O(N)
- 空间复杂度为O(1)

#### 3. 递归

```java
public ListNode reverseKGroup(ListNode head, int k) {
    // 递归结束条件
    if(head == null || head.next == null) return null;
    // 找尾指针位置
    ListNode tail = head;
    for(int i=0; i<k; i++){
        tail = tail.next;
        if(tail == null) return null;
    }
    //翻转前k个元素
    ListNode new_head = _reverse(head, tail);

    // 将已经翻转的接上以后翻转的
    // 这里的head 已经变成了翻转后的尾，所以是拼接尾巴
    head.next = reverseKGroup(tail, k);
    return new_head;
}

// 翻转链表前几个元素
// 前闭后开
public ListNode _reverse(ListNode head, ListNode tail){
    ListNode result = new ListNode();
    ListNode temp = head;
    while(head != tail){
        // 头插法
        ListNode p = head;
        head = head.next;p
        p.next = result.next;
        result.next = p;
    }
    temp.next = tail;

    return result.next;
}

```

- 时间复杂度为O(N)
- 空间复杂度为O(N)
