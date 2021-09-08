悲观锁：synchronized 关键字实现的是悲观锁，每次访问共享资源时都会上锁
非公平锁：synchronized 关键字实现的是非公平锁，即线程获取锁的顺序并不一定是按照线程阻塞的顺序
可重入锁：synchronized 关键字实现的是可重入锁，即已经获取锁的线程可以再次获取锁
独占锁或者排他锁：synchronized 关键字实现的是独占锁，即该锁只能被一个线程所持有，其他线程均被阻塞





重要：

https://blog.csdn.net/liuwg1226/article/details/119905478