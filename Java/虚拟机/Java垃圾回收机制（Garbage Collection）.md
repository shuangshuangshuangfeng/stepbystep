# <center>Java垃圾回收机制</center>

**JVM添加参数：**

参数开始固定都会有`-XX`:

- `-XX:+<option> `表示开始默认关闭的选项

- `-XX:-<option>` 表示关闭默认开启的选项

- `-XX:<option>=<value>  `



## 1.1 Java运行时一个类什么时候被加载？？？



给IDEA添加虚拟机参数`-XX:+TraceClassLoading`，用于监控类的加载。

```java
public class LoadClass {

    public static final int a = 123;

    public static int b;

    public int abc;

    public void func(){
        System.out.println("hello world");
    }

    // -XX:+TraceClassLoading JVM参数: 监控类的加载
    public static void main(String[] args){
        LoadClass loadclass = new LoadClass();
        loadclass.func();

    }

}


```

可以看出首先加载的类是各种类的父类`Object`

![](https://tcs.teambition.net/storage/31283a29601b23fc86e7f3da1e976ac2c8e4?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjgzYTI5NjAxYjIzZmM4NmU3ZjNkYTFlOTc2YWMyYzhlNCJ9.tbjotOGcvR2PmPZW7rRxgO4AUGZwIBA_w2htMHc4XyE&download=image.png "")

当用到了`LoadClass`类了，才会加载它， 并执行控制台打印。

![](https://tcs.teambition.net/storage/312864e3c993ac76e0b2e050be9c9d393dab?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjg2NGUzYzk5M2FjNzZlMGIyZTA1MGJlOWM5ZDM5M2RhYiJ9.VYG1WzLmVq_URaliJWOoNv8fpgQe8hhZJA7c-vxl3oU&download=image.png "")



**总结：**

**JVM加载类按需加载，什么时候用到什么时候加载。**



## 1.2 一个类的加载过程 ？？？

一个类从加载到`JVM`内存，到从`JVM`内存写在，期间经历七个阶段，也就是这个类的生命周期经历七个阶段：(其中验证、准备、解析并称为连接阶段)

```text
加载->验证->准备->解析->初始化->使用->卸载
加载->|----连接----| ->初始化->使用->卸载
```



### （1）加载：（涉及到类加载器）

将`classpath`、`jar`包、**网络、磁盘**某个位置下的类的**`class`二进制字节流读进**来，在**内存中生成**一个代表这个类的`**Java.lang.Class`对象放进元空间**`(JDK1.8`以后定义的名字)。

class二进制文件长这样：开头 `cafe babe`

![](https://tcs.teambition.net/storage/312802a2a3eab2d04556f366f39452fecdb0?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjgwMmEyYTNlYWIyZDA0NTU2ZjM2NmYzOTQ1MmZlY2RiMCJ9.d5hOuQIe1doNwP-qpk8uxIzDMMiYHa9vkfWveNFNQ5w&download=image.png "")



### （2）验证

验证**Class文件的字节流**中包含的信息是否符合《Java虚拟机规范》的全部约束要求，以保证虚拟机的安全。

  

### （3）准备

**类变量赋初始值**，int为0， long为0L， boolean为false， 引用类型为null，**常量也赋值**。

![](https://tcs.teambition.net/storage/312894fe8b0348455d878b3a1c09ab95420f?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjg5NGZlOGIwMzQ4NDU1ZDg3OGIzYTFjMDlhYjk1NDIwZiJ9.FczBQKxhe033dUaxaML3ThdWwSLZ2sW-uvtXoKxFkCk&download=image.png "")



### （4）解析

把符号引用翻译为直接引用



### （5）初始化

当我们`new`一个类的对象，访问一个类的静态属性，修改一个类的静态属性，调用一个类的静态方法，用反射`API`对一个类进行调用，初始化当前类，当前类的父类也会被初始化，以上的这些操作都会**触发类的初始化**。

### （6）使用

使用当前类



### （7）卸载

。。。



## 1.3 一个类的初始过程？？？

类的初始化阶段， `Java`虚拟机才真正开始执行类中编写的`Java`程序代码。

**进行准备阶段是，变量已经赋过一次系统要求的初始零值，而在初始化阶段，才真正初始化类变量和其他资源。**



**总结：**

- **加载类：类常量初始化->静态类变量->静态代码块在初始化阶段执行**

- **创建对象：类变量->普通代码块->构造器执行**



## 1.4 继承时父子类的初始化顺序？？？

1. **加载类：**

> 父类静态变量
>
> 父类静态代码块
>
> > 子类静态变量
>
> > 子类静态代码块



2. **创建对象：**

>父类变量
>
>父类代码块
>
>父类构造器
>
>> 子类变量
>
>> 子类代码块
>
>> 子类构造器



**总结：从加载类到创建对象，从父类到子类。**



## 1.5 究竟什么是类加载器？？？

在类加载阶段，通过一个类的全限定名来获取**描述该类的二进制字节流**的这个动作的代码，被称为类加载器，这个动作是可以自定义实现的。读到`JVM的`内存中。

自己写类加载器的的需要继承`java.lang.ClassLoader`类





## 2.1 什么是可回收的垃圾对象？？？

![](https://tcs.teambition.net/storage/3128034dca62c83ebd30e231b7b41b81a5af?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjgwMzRkY2E2MmM4M2ViZDMwZTIzMWI3YjQxYjgxYTVhZiJ9.ky1BhQOgfgL9Jq8vNPyXyhqrjwuNZaTc6XQpdkqKr0M&download=image.png "")



箭头中所指的就是垃圾对象，那么如何找到垃圾对象的？？？有两种算法



### （1）引用计数法

![](https://tcs.teambition.net/storage/3128726fae90d372cede9e2fa7259ab6c056?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjg3MjZmYWU5MGQzNzJjZWRlOWUyZmE3MjU5YWI2YzA1NiJ9.TqA7Ws5s3DYMd3b-ljuutFBnOBQ14CFTtvVhesnQtY4&download=image.png "")

**给对象中添加一个引用计数器**，每当一个地方引用它，计数器就会加1，当引用失效，计数器就会减1.任何情况下，计数器为0的对象就是可以被回收的对象，这个方法实现简单，效率高。

但是主流的虚拟机中并没有选择这个算法来管理内存**，起主要原因是它很难解决对象之间相互循环引用的问题**。

![](https://tcs.teambition.net/storage/3128dbf77b5e6be32cbb9c2cc8d7c6f032e3?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjhkYmY3N2I1ZTZiZTMyY2JiOWMyY2M4ZDdjNmYwMzJlMyJ9.BbBJAb_KgKz_oy1BOFD4QluDXamKfMjV5HB1p8AnsyI&download=image.png "")



### （2）可达性分析法

将`GC roots` 对象作为起点，从这些节点开始向下搜索引用的对象，找到的对象标记为非垃圾对象，其余未标记的对象都是垃圾对象。

**`GC Root`根节点：线程栈的本地变量、静态变量、本地方法栈的对象等。**

![](https://tcs.teambition.net/storage/3128518ab3ba72b8b864493361d895699ba2?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjg1MThhYjNiYTcyYjhiODY0NDkzMzYxZDg5NTY5OWJhMiJ9.nSo_81QMWs-0fQospBLIYM4d7RLzQ6KxastW-ROyqr8&download=image.png "")





## 2.2 垃圾回收算法？？？

编程语言有很多种，但是其垃圾回收算法其实就那么几个。



### （1）标记清除算法

![](https://tcs.teambition.net/storage/3128907b66d58b699c36442a4c5ef5270fe1?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjg5MDdiNjZkNThiNjk5YzM2NDQyYTRjNWVmNTI3MGZlMSJ9.nY3mD3J3Fy8uL9dUf5oUzrCG_0ch_2sbCBh3fMCVU68&download=image.png "")



**缺点： 空间利用率底，位置不连续，会产生碎片**



### （2）复制算法

![](https://tcs.teambition.net/storage/3128e9c67eb33ae85463095c7620e607a463?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjhlOWM2N2ViMzNhZTg1NDYzMDk1Yzc2MjBlNjA3YTQ2MyJ9.gjFsXpeIyPTQ9GJq-9gAbQCbv-vwGcfy3bCDnI9YaFY&download=image.png "")

内存一分为二， 当其中一块清理完后，剩余的对象复制到另一块空间紧密排列。



**缺点：没有碎片，但是浪费空间，有一半空间是在运行当中没有被用到的。**



### （3）标记整理算法

![](https://tcs.teambition.net/storage/31288c62dbec2422bd75e5c01d1c8cec5e78?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjg4YzYyZGJlYzI0MjJiZDc1ZTVjMDFkMWM4Y2VjNWU3OCJ9.JcCyxKEGTIWhLcY5DgSEtgsby_cLIC_6RBJhdEBJ_60&download=image.png "")



标记清理之后，会把空间整理一下。

**缺点：可能会有点浪费时间**



## 2.3 垃圾收集器有什么？？？

到目前为止，一共有十种垃圾收集器，

![](https://tcs.teambition.net/storage/3128145517b4caca46e7946c7a9e3975cc23?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjgxNDU1MTdiNGNhY2E0NmU3OTQ2YzdhOWUzOTc1Y2MyMyJ9.mrq7W6BBKTiyDpNpZYIrJuKD5gqLuAahOshj9R_EC5o&download=image.png "")

左边的是`JDK1.8`以及之前用的垃圾收集器，而且上部分主要用在内存模型的年轻代，下边的主要用于内存模型的老年代。

> 分代模型：`JDK1.8`以及之前的版本

> 分区模型：`JDK1.9`及之后的版本



### 什么是年轻代什么是老年代？？？

![](https://tcs.teambition.net/storage/3128849b55aff1a1200a571b906c7ee99da0?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjg4NDliNTVhZmYxYTEyMDBhNTcxYjkwNmM3ZWU5OWRhMCJ9.s_w2_-bjQqprSuPXf8Jv5PdAsRX6z7qp4XlZAkh1Mjw&download=image.png "")

**类装载系统**：是将`.class`字节流二进制文件加载到`JVM`的内存中，也就是上面一节讲的类加载器

**执行引擎**：程序的执行，垃圾线程的开启，这些都是执行引擎完成的

**运行时数据区**：也就是我们所说的内存模型

黄色的是线程共享的区域，紫色的是是线程私有的，包括`JVM`调优，最主要是对堆内存进程调优，所以要把堆内存的内存结构弄明白。



**注意：堆内存的结构是和`JDK`的版本是有关系的，也就是说使用`JDK`不同的版本，它的内存里面的结构是有区别的。**

![](https://tcs.teambition.net/storage/31288caaa156f7a5ea0927c9405a03b1897c?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjg4Y2FhYTE1NmY3YTVlYTA5MjdjOTQwNWEwM2IxODk3YyJ9.Vq5vsswQgWCfyrk1u8MDYwITYG5qjbu0zYkS8mzMVms&download=image.png "")

**堆内存分为年轻代和老年代，并且年轻代和老年代的内存是有比例的，默认情况下，年轻代是占三分之一，老年代是占三分之二。**但是这个比例是可以调的，根据我们的需求灵活的去调整这个分配。

**年轻代中还会有很多分区，为什么会这么分呢？？？**伊甸园（新生的，刚创建的）

(1) 新生的刚创建的对象放在伊甸园区的，但是随着不断的创建对象，这个伊甸园区放不下了，这时候执行引擎就会开启一个垃圾收集线程，那么如何判断是不是垃圾对象？？用根可达性分析算法。非垃圾的对象就会放到幸存区，此时年龄+1（表示它经历过一次`GC`垃圾回收）。

(2) 当伊甸园区的内存又要满了，执行引擎启动一个垃圾回收线程，清理垃圾(使用可达性分析算法)，然后将上次幸存区的对象的年龄再加1，然后向后挪动，将伊甸园中幸存的对象放到年龄为2的位置中(使用复制算法)，以后一直这样进行。。。

(3) 如果有的对象一直存在，年龄一直加1，加到一定程度以后（15次），我们就认为这个对象一直被引用，那么就把这个对象放到老年代中，

![](https://tcs.teambition.net/storage/312895b3b020cf32997926733711dab5d42b?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjg5NWIzYjAyMGNmMzI5OTc5MjY3MzM3MTFkYWI1ZDQyYiJ9.Xxk5NdsEoIcgoedAQjXj-kpZD2XPAIQH4Hc-wrYM6IY&download=image.png "")



**如果老年代中的内存也满了怎么办？？？**

执行引擎会开启`GC`垃圾收集线程（`full GC`）会把老年代的垃圾对象进行回收，同时年轻代的对象也会区回收。所以可以看出一旦开启这个`full GC`, 这个线程所需要的时间一定是比较长的。并且，**它在做垃圾回收的时候会导致用户线程阻塞暂停。所以一个网站中一旦频繁的产生full GC, 一定要去好好的优化优化**。  

垃圾回收的时候具体用哪个算法，要看它使用的是哪个垃圾收集器。



### 演示内存中的模型

**`jvisualvm`**  用来看`JVM`内存中的情况，是`java`自带的。

![](https://tcs.teambition.net/storage/3128f28012f6659bc06538b9839f31515e07?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjhmMjgwMTJmNjY1OWJjMDY1MzhiOTgzOWYzMTUxNWUwNyJ9.YOzgNL5H8GPhI1ULOakMBtG5nXt_6GXGxQDuh6a_qWM&download=image.png "")

```java
import java.util.ArrayList;
import java.util.List;
import static java.lang.Thread.sleep;

public class LoadClass {
    public static void main(String[] args) throws InterruptedException {
            List<LoadClass> list = new ArrayList<>();
            while (true){
                list.add(new LoadClass());
            }
    }
}
```

用上面的代码一直去执行，然后就会把堆内存中的老年代占满，然后报内存溢出的错误。程序结束执行。

![](https://tcs.teambition.net/storage/3128008c316f60311c83d4452c2a9c833923?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjgwMDhjMzE2ZjYwMzExYzgzZDQ0NTJjMmE5YzgzMzkyMyJ9.9sM6ebyZo_cAU_yHOkfgSKfaPdzTQ1jfXPM4Ps5sceA&download=image.png "")





 一般来说，年轻代的垃圾回收器和老年代的垃圾回收器组合来用。

但是有一个例外，就是`CMS（Concurrent Mark Sweep）`并发标记清除

### CMS

`CMS`的优点就是处理较大的内存数据比较快，相比于其他的垃圾回收器要快得多，它使用的算法是标记清楚法，所以他也有一个缺点，就是有空间碎片

![](https://tcs.teambition.net/storage/3128e27b54f6ec8ac51f59def272c9de9f28?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NjI2MiwiaWF0IjoxNjMwOTQxNDYyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjhlMjdiNTRmNmVjOGFjNTFmNTlkZWYyNzJjOWRlOWYyOCJ9.CmQb9RppPrKaBHNdhyJiszyFC2VPyS7AJ5EgWSW1uSA&download=image.png "")

## 2.4 内存屏障

内存屏障（`Memory Barrier`，或有时叫做内存栅栏，`Memory Fence`）是一种CPU指令，用于控制特定条件下的重排序和内存可见性问题。Java编译器也会根据内存屏障的规则禁止重排序。

**内存屏障可以被分为以下几种类型**

1. `LoadLoad`屏障：对于这样的语句`Load1; LoadLoad; Load2`，在`Load2`及后续读取操作要读取的数据被访问前，保证`Load1`要读取的数据被读取完毕。

2. `StoreStore`屏障：对于这样的语句`Store1; StoreStore; Store2`，在`Store2`及后续写入操作执行前，保证`Store1`的写入操作对其它处理器可见。

3. `LoadStore`屏障：对于这样的语句`Load1; LoadStore; Store2`，在`Store2`及后续写入操作被刷出前，保证`Load1`要读取的数据被读取完毕。

4. `StoreLoad`屏障：对于这样的语句`Store1; StoreLoad; Load2`，在`Load2`及后续所有读取操作执行前，保证`Store1`的写入对所有处理器可见。它的开销是四种屏障中最大的。

5. 在大多数处理器的实现中，这个屏障是个万能屏障，兼具其它三种内存屏障的功能。有的处理器的重排序规则较严，无需内存屏障也能很好的工作，`Java`编译器会在这种情况下不放置内存屏障。

