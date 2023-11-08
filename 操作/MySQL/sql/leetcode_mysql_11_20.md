11. [584. 寻找用户推荐人 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/find-customer-referee/)

```sql
SELECT name
FROM customer 
WHERE referee_id != 2 OR referee_id IS NULL; # 注意为空不是=null， 而是is NULL
```

12. [586. 订单最多的客户 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/customer-placing-the-largest-number-of-orders/)

```sql
select customer_number
from orders
group by customer_number  # 按照customer_number分组
order by count(*) desc # 分组后， 从大到小排， count(*)统计数量
limit 1; # 限制只取一个
```

13. [595. 大的国家 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/big-countries/)

```sql
select name, population, area
from World
where area>3000000 or population >25000000;
```

14.[596. 超过5名学生的课 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/classes-more-than-5-students/)

```sql
select class
from courses
group by class
having count(DISTINCT student)>=5; # distinct student， student这一列不可重复
```

15.[597. 好友申请 I：总体通过率 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/friend-requests-i-overall-acceptance-rate/)

```sql
select round(ifnull(
    (select count(distinct requester_id, accepter_id) from RequestAccepted ) /
    (select count(distinct sender_id, send_to_id) from FriendRequest)
    ,0),2) accept_rate;
```
