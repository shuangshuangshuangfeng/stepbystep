代码块分四种：

- 普通代码块

- 构造代码块

- 静态代码块

- 同步代码块(synchronized修饰的)



## 1. 普通代码块

在函数中括起来的代码块叫普通代码块，函数括起来的代码块也叫普通代码块，只有在被调用函数的时候才能被执行。

```java
public class test {

    public void function(){
        {
            System.out.println("普通代码块");
        }
    }

    public static void main(String[] args){
        test test1 = new test();
    }
}


```



## 2. 静态代码块

在类中定义的static的代码块(注意不是在方法里定义的)，静态代码块用于初始化类，为类的属性初始化，每个静态代码块只会执行一次，在JVM加载类的时候执行。

```java
public class test {

    static {
        /*
        * 静态代码块用于初始化类，为类的属性初始化，每个静态代码块只会执行一次。
        * JVM加载类时会执行静态代码块的
        * */
        System.out.println("静态代码块");
    }

    public static void main(String[] args){
        test test1 = new test();
        test test2 = new test();
        test test3 = new test();

    }
}

```



## 3. 构造代码块

在类中的，但是不用static修饰的叫构造代码块，每次创建对象都会调用一次

```text
public class test {

    {   /*
        * 构造代码块在创建对象（new）的时候被调用
        * */
        System.out.println("构造块1");
    }
    {
        System.out.println("构造块2");
    }

    public static void main(String[] args){
        test test1 = new test();
        test test2 = new test();
        test test3 = new test();

    }
}


```



## 4. 同步代码块

使用synchronized 修饰的代码块叫同步块，当多个线程访问的时候，只有一个线程可以执行

```java
class Count{
    private int count = 100;

    /*
    * 这里是同步块，当多个线程访问的时候，只有一个线程可以执行
    * */
    public  int getCount(String thr_name) {
        System.out.println(thr_name+"减减之前_未在同步块中_获取Count值:"+this.count);
        synchronized(this){
            this.count --;
            System.out.println(thr_name+"_获取Count值:"+this.count);
        }
        this.count --;
        System.out.println(thr_name+"_获取Count值:"+this.count);
        return count;
    }

    public void setCount(int count) {
        this.count = count;
        System.out.println("设置Count值:"+this.count);
    }
}


public class test1 extends Thread{
    private String thr_name;
    private Count count;
    public test1(String thr_name, Count count){
        this.thr_name = thr_name;
        this.count = count;
    }

    @Override
    public void run() {
        super.run();
        while (true){
            count.getCount(thr_name);
            try {
                Thread.sleep(1000);
            }catch (Exception e){
                e.printStackTrace();
            }

        }
    }

    public static void main(String[] args){
        Count count = new Count();
        count.setCount(100);
        test1 test1 = new test1("A", count);
        test1 test2 = new test1("B", count);
        test1.start();
        test2.start();
    }
}
```



运行结果：可以看出在同步块中的打印结果是正确的，同步块以外打印结果并不是按顺序

![](https://tcs.teambition.net/storage/31284daf7a72d427b34b8c2741c447e02493?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTU0NTk5OCwiaWF0IjoxNjMwOTQxMTk4LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjg0ZGFmN2E3MmQ0MjdiMzRiOGMyNzQxYzQ0N2UwMjQ5MyJ9.ix10wTCQldIw5KZnNfTYmmr1NSkQK4clBYKkFBIkx88&download=image.png "")

