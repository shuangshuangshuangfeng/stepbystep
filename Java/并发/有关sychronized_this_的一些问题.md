# synchronized(this)



1. **`synchronized`锁住的对象：**

- **对于同步方法**，锁住当前对象 （见下文）

- **对于静态同步代码块**，锁住的是当前类的`Class对`象
- **对于同步代码块**，锁住的是`synchronized`括号中的对象, ？？？



2. 具体 （代码块）

   - 当两个并发线程，**访问同一个对象**的`object`中的**`synchronized(this) `代码块**是，一个时间只能有一个线程执行，另一个线程阻塞，直到上一个线程执行完毕。（也就是**<font color='red'>第一个线程获得了object的对象锁，其他线程需要阻塞</font>**）

   - 当一个线程访问`object`的一个`synchronized(this)`同步代码块时，另一个线程仍然可以访问该`object`的非`synchronized(this)`同步代码块。

   - 当**一个线程访问`object`的一个`synchronized(this)`同步代码块**时，其他线程对**`object`所有其他的`synchronized(this)`同步代码块**的访问阻塞。



3. 例子

   ```Java
   class ObjectA{
       public int cnt = 0;
       public String thName ;
       public ObjectA(){
       }
   
       public synchronized void add(String thr_name){
           thName = thr_name;
           for(int i=0; i<100; i++){
               try {
                   Thread.sleep(100);
               } catch (InterruptedException e) {
                   e.printStackTrace();
               }
               cnt++;
               System.out.println(thName+" cnt:"+cnt);
           }
   
       }
   
       public synchronized void show(String thr_name){
           thName = thr_name;
           for(int i=0; i<10; i++){
               System.out.println(thName + " cnt show :"+cnt);
               try {
                   Thread.sleep(1000);
               } catch (InterruptedException e) {
                   e.printStackTrace();
               }
           }
       }
   }
   
   
   public class Test extends Thread{
       public ObjectA obA;
       public Integer num;
       public String Name;
       public Test(ObjectA o, int n, String name){
           obA = o;
           num = n;
           Name = name;
       }
   
       @Override
       public void run() {
           if(num == 0){
               obA.add(Name);
           }else{
               obA.show(Name);
           }
       }
   
       public static void main(String[] args){
           ObjectA oa = new ObjectA();
           Test t1 = new Test(oa, 0, "thread1");
           Test t2 = new Test(oa, 1, "thread2");
           t1.start();
           t2.start();
       }
   
   }
   
   
   ```

运行结果如下：**说明锁住的是object的对象**

![image-20210908111528066](有关sychronized(this)的一些问题.assets/image-20210908111528066.png)



4. Java中创建线程主要有三种方式：
   - 继承Thread类
   - 现Runnable接口
   - 使用Callable和Future

