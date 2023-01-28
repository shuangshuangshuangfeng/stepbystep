1. [175. 组合两个表 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/combine-two-tables/)

```sql
select p.FirstName, p.LastName, a.City, a.State from Person as p 
left join Address as a
on p.PersonId = a.PersonId;
```

2. [176. 第二高的薪水 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/second-highest-salary/)

```sql
select 
    ifnull(
        (select DISTINCT Salary 
        from Employee 
        order by Salary DESC
        limit 1 offset 1), null
    ) as SecondHighestSalary;
```

3. [181. 超过经理收入的员工 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/employees-earning-more-than-their-managers/)

```sql
SELECT A.Name as 'Employee'
FROM Employee AS A, Employee AS B  # A是员工， B是经理
WHERE A.ManagerId = B.Id
and A.Salary > B.Salary;
```

4. [182. 查找重复的电子邮箱 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/duplicate-emails/)

```sql
select p.Email
from Person as p
group by p.Email 
having count(p.Email)>1;
```

5. [183. 从不订购的客户 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/customers-who-never-order/)

   ```sql
   select c.name as 'Customers'
   from Customers as c 
   where c.id not in (
       select CustomerId from Orders
   );
   ```

6. [196. 删除重复的电子邮箱 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/delete-duplicate-emails/)

```sql
delete p1 
from person as p1 , person as p2
where p1.Email=p2.Email and p1.Id > p2.Id;
```

7. [197. 上升的温度 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/rising-temperature/)

```sql
select w1.id
from weather as w1 join weather as w2
on (w1.recordData = adddate(w2.recordData, INTERVAL 1 DAY ))
where w1.temperature> w2.temperature;

```

8. [511. 游戏玩法分析 I - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/game-play-analysis-i/submissions/)

```sql
select a.player_id, min(event_date) as 'first_login'
from Activity as a
group by player_id;
```

9. [512. 游戏玩法分析 II - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/game-play-analysis-ii/)

```sql
select b.player_id, b.device_id
from Activity as b
where (b.player_id, b.event_date) in(
    select a.player_id, min(event_date)
    from Activity as a 
    group by player_id
)

```

10. [577. 员工奖金 - 力扣（LeetCode） (leetcode-cn.com)](https://leetcode-cn.com/problems/employee-bonus/)

```sql
select name, bonus
from Employee left join Bonus
on Employee.EmpId = Bonus.EmpId
where Bonus.bonus<1000 or Bonus.bonus is null;

```

