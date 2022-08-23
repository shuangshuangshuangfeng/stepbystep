# HashMap

## `get()`jdk1.8以后

```java
public V get(Object key) {
    Node<K,V> e;
    // 计算hash值,与key一起找key的值
    return (e = getNode(hash(key), key)) == null ? null : e.value;
}

/**
 * Implements Map.get and related methods
 *
 * @param hash hash for key
 * @param key the key
 * @return the node, or null if none
 */
final Node<K,V> getNode(int hash, Object key) {
    // 设置一些局部变量
    Node<K,V>[] tab; Node<K,V> first, e; int n; K k;
    // 首先获取hashmap中的数组和长度,并判断是否为空,如果为空,返回null
    if ((tab = table) != null && (n = tab.length) > 0 &&
        (first = tab[(n - 1) & hash]) != null) {
        // 获取key对应的下标对应的链表对象, 并比较第一个是否满足条件
        if (first.hash == hash && // always check first node
            ((k = first.key) == key || (key != null && key.equals(k))))
            // 第一个如果满足条件,则直接返回
            return first;
        // 判断当前对象是否是最后一个,如果是,说明没有找到对应的key的值
        if ((e = first.next) != null) {
            // 如果不为空,判断是否是红黑树,如果是红黑树,使用红黑树获取对应key的值
            if (first instanceof TreeNode)
                return ((TreeNode<K,V>)first).getTreeNode(hash, key);
            // 如果不是红黑树, 遍历链表,找到对应hash和key的node对象
            do {
                if (e.hash == hash &&
                    ((k = e.key) == key || (key != null && key.equals(k))))
                    return e;
            } while ((e = e.next) != null);
        }
    }
    return null;
}
```





## `put` JDK1.8以后



自己实现的思路：



```java
public V put(K key, V value) {
    // 获取key的value值,调用用putval方法
    return putVal(hash(key), key, value, false, true);
}


/**
* Implements Map.put and related methods
*
* @param hash hash for key
* @param key the key
* @param value the value to put
* @param onlyIfAbsent if true, don't change existing value
* @param evict if false, the table is in creation mode.
* @return previous value, or null if none
*/
final V putVal(int hash, K key, V value, boolean onlyIfAbsent, boolean evict) {
    Node<K,V>[] tab; Node<K,V> p; int n, i;
    // 首先创建局部变量并进行赋值,获取当前数组的长度
    if ((tab = table) == null || (n = tab.length) == 0)
        n = (tab = resize()).length;
    // 判断key对应下标的数组位置是否为空,如果为空,在这个下标位置创建一个
    if ((p = tab[i = (n - 1) & hash]) == null)
        tab[i] = newNode(hash, key, value, null);
    else {
        // 如果key对应的参数不为空
        Node<K,V> e; K k;
        // 判断新添加的对象与旧对象中第一个对象的key的hash值和key是否相等
        if (p.hash == hash &&
            ((k = p.key) == key || (key != null && key.equals(k))))
            // 如果相等, node值存进临时值
            e = p;
        else if (p instanceof TreeNode)
            // 如果 已经存在的node对象是 treennode(红黑树)的实现类, 调用treenode的putval方法存
            e = ((TreeNode<K,V>)p).putTreeVal(this, tab, hash, key, value);
        else {
            // 如果新添加的key不遇第一个node中的key相同,并且不是treenode的实现类,遍历node链表
            for (int binCount = 0; ; ++binCount) {
                // 判断当前node是否是最后一个node,如果是,把新添加的node添加进这个链表中
                if ((e = p.next) == null) {
                    p.next = newNode(hash, key, value, null);
                    // TREEIFY_THRESHOLD 默认是8, 即判断当前链表的长度是否大于7,如果是,则把链表转换成红黑树
                    if (binCount >= TREEIFY_THRESHOLD - 1) // -1 for 1st
                        treeifyBin(tab, hash);
                    // 执行到最后一个后,跳出循环
                    break;
                }
                // 如果当前node不是最后一个,比较key和hash值,如果存在相同,跳出循环(此时已经找到链表中key对应的node)
                if (e.hash == hash &&
                    ((k = e.key) == key || (key != null && key.equals(k))))
                    break;
                // 下一个node返回
                p = e;
            }
        }
        // 判断找到的node是否为空,如果不为空
        if (e != null) { // existing mapping for key
            // 获取旧值
            V oldValue = e.value;
            // 如果设置是否可以改变值,如果设置不能改,但是旧值为空,则可以改,否则,需要设置可以改值的情况下才可以把新值赋给旧值
            if (!onlyIfAbsent || oldValue == null)
                e.value = value;
            afterNodeAccess(e);
            // 返回旧值
            return oldValue;
        }
    }
    ++modCount;
    // 判断当前数量是否超过扩容阈值,如果超过,进行扩容, 与1.7先扩容再添加不同,1.8是先扩容再添加
    if (++size > threshold)
        resize();
    afterNodeInsertion(evict);
    return null;
}

```

