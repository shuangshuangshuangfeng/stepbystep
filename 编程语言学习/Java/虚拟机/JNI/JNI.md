# JNI「Java Native Interface」

JNI是使用Java本地接口，可以方便在不同的平台移植。即允许Java代码调用本地的C与C++接口。

## 自己写一个JNI Demo

**官方API：**

[Java Native Interface Specification: 4 - JNI Functions](https://docs.oracle.com/en/java/javase/11/docs/specs/jni/functions.html)

（1）首先在建一个JNIDemo.java文件，内容如下：

```java
public class JNIDemo {
    static {
        System.loadLibrary("JNIDemo");
    }

    public native String helloWorld();

    public static void main(String[] args){
        JNIDemo demo = new JNIDemo();
        System.out.println(demo.helloWorld());
    }
}
```

（2）进行编译和生成头文件：

```shell
➜  test git:(main) ✗ javac JNIDemo.java
➜  test git:(main) ✗ javah JNIDemo
➜  test git:(main) ✗ ll
总用量 12K
-rw-rw-r-- 1 home home 561  2月  3 15:54 JNIDemo.class
-rw-rw-r-- 1 home home 392  2月  3 15:54 JNIDemo.h
-rw-rw-r-- 1 home home 259  2月  3 15:53 JNIDemo.java
```

（3）生成一个JNIDemo.c文件， 并填写以下内容：

```c
#include <jni.h>     // 系统目录
#include "JNIDemo.h" // 当前目录


JNIEXPORT jstring JNICALL Java_JNIDemo_helloWorld (JNIEnv *env, jobject){
        return (*env)->NewStringUTF(env, "hello world!");

}
```

（4）执行以下命令，生成.so文件

> gcc -fPIC -I /home/home/Applications/jdks/jdk-8u202-linux-x64/jdk1.8.0_202/include -I /home/home/Applications/jdks/jdk-8u202-linux-x64/jdk1.8.0_202/include/linux -shared -o libJNIDemo.so JNIDemo.c

```shell
➜  test git:(main) ✗ gcc -fPIC -I /home/home/Applications/jdks/jdk-8u202-linux-x64/jdk1.8.0_202/include -I /home/home/Applications/jdks/jdk-8u202-linux-x64/jdk1.8.0_202/include/linux -shared -o libJNIDemo.so JNIDemo.c
➜  test git:(main) ✗ ll
总用量 32K
-rw-rw-r-- 1 home home 170  2月  3 16:01 JNIDemo.c
-rw-rw-r-- 1 home home 561  2月  3 15:54 JNIDemo.class
-rw-rw-r-- 1 home home 392  2月  3 15:54 JNIDemo.h
-rw-rw-r-- 1 home home 259  2月  3 15:53 JNIDemo.java
-rwxrwxr-x 1 home home 15K  2月  3 16:01 libJNIDemo.so
```

（5）执行Java主类 `java -Djava.library.path=. JNIDemo`

```shell
➜  test git:(main) ✗ java JNIDemo 
Exception in thread "main" java.lang.UnsatisfiedLinkError: no JNIDemo in java.library.path
        at java.lang.ClassLoader.loadLibrary(ClassLoader.java:1867)
        at java.lang.Runtime.loadLibrary0(Runtime.java:870)
        at java.lang.System.loadLibrary(System.java:1122)
        at JNIDemo.<clinit>(JNIDemo.java:3)
错误: 找不到或无法加载主类 .
➜  test git:(main) ✗ java -Djava.library.path=. JNIDemo 
hello world!
```


