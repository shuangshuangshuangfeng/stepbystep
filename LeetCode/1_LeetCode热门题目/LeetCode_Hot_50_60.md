[142. 环形链表 II - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/linked-list-cycle-ii/)

```java
public class Solution {
    public ListNode detectCycle(ListNode head) {
        if(head == null) return null;
        ListNode fast = head;
        ListNode slow = head;
        while(fast!=null && fast.next != null){
            fast = fast.next.next;
            slow = slow.next;
            if(fast == slow) break;
        }
        if(fast == null || fast.next==null) return null;
       
        ListNode tmp = head;
        while(tmp != slow){
            slow = slow.next;
            tmp = tmp.next;
        }
        return tmp;
    }
}
```



[146. LRU 缓存机制 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/lru-cache/)

最近最少使用

```java



```

![img](https://tcs.teambition.net/storage/31270fe622082069ad27d5becc9067014b45?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzNDIxNDg5NSwiaWF0IjoxNjMzNjEwMDk1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjcwZmU2MjIwODIwNjlhZDI3ZDViZWNjOTA2NzAxNGI0NSJ9.9IB8qYaGywCgSKRuv7XmPMSAA81DgQv7FBECBFmyns8)
