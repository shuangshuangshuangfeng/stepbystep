

## 1. 线程池的所有类的关系

Java线程池中，**最基础的一个接口：Executor**，ExecutorService继承Executor接口，AbstractExecutorService接口继承ExecutorService：

```text
Executor->ExecutorService->AbstractExecutorService->各种线程池的类
```



**最重要的一个类：Executors**,  这是创建所有种类线程池的类



## 2. 线程池中的几个重要参数

下边是类ThreadPoolExecutor的构造函数

```java

/**
	  核心线程： threads to keep in the pool ， 会一直保持在线程池中的线程
     *  默认的线程工厂的初始化参数
     * @param corePoolSize : 要保留在线程池的核心线程数量（包括空闲状态）， 核心线程在空闲状态下是不会被回收的
     * @param maximumPoolSize : 池中允许的最大线程数
     * @param keepAliveTime : 非核心空闲线程的存活时间
     * @param unit : keepAliveTime 存活时间的时间单位
     * @param workQueue： 任务队列
     */

public ThreadPoolExecutor(int corePoolSize,
                              int maximumPoolSize,
                              long keepAliveTime,
                              TimeUnit unit,
                              BlockingQueue<Runnable> workQueue) {
        this(corePoolSize, maximumPoolSize, keepAliveTime, unit, workQueue,
             Executors.defaultThreadFactory(), defaultHandler);
}
```

**重要的参数也就是注释这些：**

```java
/**
	  核心线程： threads to keep in the pool ， 会一直保持在线程池中的线程
     *  默认的线程工厂的初始化参数
     * @param corePoolSize : 要保留在线程池的核心线程数量（包括空闲状态）， 核心线程在空闲状态下是不会被回收的
     * @param maximumPoolSize : 池中允许的最大线程数
     * @param keepAliveTime : 非核心空闲线程的存活时间
     * @param unit : keepAliveTime 存活时间的时间单位
     * @param workQueue： 任务队列
     */
```



## 3. 线程池的用法



### （1） 直接写代码块

```java
ExecutorService executor = Executors.newFixedThreadPool(5); // 创建线程池???
executor.submit(() ->{  /* 创建一个任务，交给线程池去管理运行 */
    System.out.println("可以这样写任务代码段...");
});

```



### （2）执行继承Runnable接口的类中的run方法

```java
ExecutorService executor = Executors.newFixedThreadPool(5); // 创建线程池???
Run run = new Run("thr_A");
executor.submit(run);  // 创建一个任务，交给线程池去管理运行
```



```java
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

class Run implements Runnable{
    /*
    * 自定义任务
    * */
    private int cnt = 0;
    private String thr_name;

    public Run(String thr_name){
        this.thr_name = thr_name;
    }

    @Override
    public void run() {
        while (true){
            System.out.println(this.thr_name+"_cnt : "+this.cnt);
            this.cnt++;
            try {
                Thread.sleep(1000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}

public class TestThrPool {
    public void func(){
        System.out.println("这是调用方法func");
    }

    public static void main(String[] args){
        //  TODO: ThreadPoolExecutor
        ExecutorService executor = Executors.newFixedThreadPool(5); // 创建线程池???

        executor.submit(() ->{  /* 创建一个任务，交给线程池去管理运行 */
            System.out.println("可以这样写任务代码段...");
        });

        Run run = new Run("thr_A");
        executor.submit(run);  // 创建一个任务，交给线程池去管理运行


    }
}


```

