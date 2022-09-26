
在``Object``类中，``equals()``方法代码如下:

```
    public boolean equals(Object obj) {
        return (this == obj);
    }

```
所以说 `Object`类中`equal()`比较的是**引用** 是否一致。

所有的类都继承自``Object``类，是可以重写``equals()``方法的，比如``String``类中，``equals()``方法的代码如下:
```
    public boolean equals(Object anObject) {
        if (this == anObject) {
            return true;
        }
        if (anObject instanceof String) {
            String anotherString = (String)anObject;
            int n = value.length;
            if (n == anotherString.value.length) {
                char v1[] = value;
                char v2[] = anotherString.value;
                int i = 0;
                while (n-- != 0) { 
                    if (v1[i] != v2[i])
                        return false;
                    i++;
                }
                return true;
            }
        }
        return false;
    }
```
**注意：** 

while (n-- != 0)``是先判断后``--``;<br>
`` while (--n != 0)``是先``--``后判断

记忆规则：减在前先做减



