

# 01 标准IO介绍

`I/O`：`input & output`， 是一切实现的基础，主要分为

- `stdio` 标准IO （所有的标准都是为了和稀泥整合不同）

- `sysio` 系统调用IO

**注意**：标准库函数都在`man`手册的第三章,如果`man`手册不了解的话，可以先`man man`， 将`man`手册的使用搞清楚。





## **`stdio`中的函数** （`FILE`类型贯穿始终）

- 文件的操作：

    - `fopen();`

    - `close();`

- 字符的操作：

    - `fgetc();`

    - `fputc();`

- 字符串的操作：

    - `fgets();`

    - `fputs();`

- 二进制数据块操作:

    - `fread();`

    - `fwrite();`

- `printf();`一族的函数

- `scanf();`一族的函数

- `fseek();`

- `ftell();`

- `rewind();`

- `fflush();`



### 思维导图介绍：





## 新建节点`man`的用法

1. `man`手册第三章是标准的库函数, 例如`man fopen`, 会自动跳到第三章

![](https://tcs.teambition.net/storage/31240fc43cb8ae8fa564f1df63bd6ab37c0b?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTI5MiwiaWF0IjoxNjMwOTMwNDkyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQwZmM0M2NiOGFlOGZhNTY0ZjFkZjYzYmQ2YWIzN2MwYiJ9.R9ZVMeXiBQnVgf26mX3KHzjkFo-SAU3Kr5S728_Tb6U&download=image.png "")

1. 第二章是系统调用

1. 第一章是基本的命令, 例如 `man 1 ls`

1. 对于开发者来说，最重要的是第七章，第七章讲的是机制，比如不明白什么是tcp， 可以用`man 7 tcp` 来解释什么是TCP。`man epoll`机制等





## `fopen()`函数

### **`man`文档描述：**

![](https://tcs.teambition.net/storage/31242c02ed18e9dedcb1df3db7b1f9b64242?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTI5MiwiaWF0IjoxNjMwOTMwNDkyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQyYzAyZWQxOGU5ZGVkY2IxZGYzZGI3YjFmOWI2NDI0MiJ9.j_6lPzDuWwdaXRNExZN8ynFGfSNJ2uRrq1xc2imLxHg&download=image.png "")

**`mode`**参数指向一个字符串，但是只取字符串的第一个字符，所以可以写成`mode="write"`，也可以写成`mode="w"`， 他们的效果是一样的。

### **例子** :

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>

int main()
{
    FILE *fp;
    fp = fopen("tmp", "r");
    if(fp == NULL)
    {
        fprintf(stderr, "fopen() failed! errno=%d\n", errno);
        perror("fopen() 函数"); //推荐使用, 可以自动关联全局变量errno
        fprintf("fopen 失败了， 失败原因是:%s", strerror(errno));
        exit(1);
    }
    puts("OK!");
    exit(0);
}
```





## `FILE`指针指向的是哪里的内存？

**栈 、 静态区、堆？？** 

```c
FILE *fopen(const char *pathname, const char *mode)
{

    return ...; //指针
}
```

1. 如果在栈上，当程序执行到`fopen()`函数时，创建一个栈上的变量，但是函数结束后，该空间被释放，所以在栈上是不可能的。

1. 在栈上保存是不现实的，那么可以保存在静态空间上，但是，函数中的变量放在静态区的时候，如果函数被重复调用，变量在静态区中的内存只有一块，第二次调用的时候第二个文件的结构体会把第一次调用的文件的数据覆盖，所以不适合放在静态区。

1. 如果放在堆上， 需要`malloc()`空间。相对应的`free()`函数是放在`fclose()`函数中的。这是成对出现的。

**注意** ：如果一个函数的返回值是指针，同时会有一个逆操作，那么这个函数的指针是放在堆上的。





## 生成的文件权限

**注意：**

- 谁打开谁关闭

- 谁申请谁释放

- 是资源就有上限

当前进程中可以打开文件，但是我打开的文件一定是有上限的。下边程序是看进程中最多可以打开多少文件。

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>

int main()
{
    FILE *fp=NULL;
    int count = 0;
    
    while(1)
    {
        fp = fopen("tmp", "w");
        if(fp == NULL)
        {
            perror("fopen() 函数"); 
            break;
        }
        count++;
    }
    printf("count = %d", count);
    exit(0);
}
```

执行结果：

![](https://tcs.teambition.net/storage/3124c5e1bc5691e26ed2043fe6b3577f6dd7?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTI5MiwiaWF0IjoxNjMwOTMwNDkyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjRjNWUxYmM1NjkxZTI2ZWQyMDQzZmU2YjM1NzdmNmRkNyJ9.LexKBHmv774PWZlzjcYa9J_LFBCzsJaKhg2hRUHxU8Q&download=image.png "")

**注意：** 在不更改当前任何环境的情况下，一个进程打开默认打开了三个流(`stream`), 分别是: `stdin`标准输入、 `stdout`标准输出、 `stderr`标准出错。所以最多可以打开的文件流是`1021+3=1024`个。



### `ulimit -a` 命令

有一个命令，叫做`ulimit -a`

![](https://tcs.teambition.net/storage/31241e755fa6bcfbd959d561969fa356ab86?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTI5MiwiaWF0IjoxNjMwOTMwNDkyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQxZTc1NWZhNmJjZmJkOTU5ZDU2MTk2OWZhMzU2YWI4NiJ9.p1KNXIq5pKJTknNSx5uG89sbINhFQCarqd6rhX4YKY8&download=image.png "")

这里显示了`open file`的最大个数是1024.

如果写`ulimit -n 1000` 这个命令的意思是修改`open file`的最大数值。



### 生成文件的权限

如果文件不存在的话，`mode=w`权限会创建文件.

![](https://tcs.teambition.net/storage/3124e22218376dd6e471b39c09728cecfa9d?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTI5MiwiaWF0IjoxNjMwOTMwNDkyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjRlMjIyMTgzNzZkZDZlNDcxYjM5YzA5NzI4Y2VjZmE5ZCJ9.QN0ZVVghlPuxyna-gbVDefrb3DPPx-bQdHTHEJs5mNY&download=image.png "")

生成的文件权限是<font color='red'>664</font>， 但是程序中，并没有指定文件的权限。这个文件的权限是怎么来的？

**生成的文件的权限是遵循一个公式：**`0666 & ~umask`umask的值是`0002`（0开头的是8进制数）：转换为二进制数是`000 000 010`

![](https://tcs.teambition.net/storage/31240d5da4d9756275fa7030fbd55fa87371?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTI5MiwiaWF0IjoxNjMwOTMwNDkyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQwZDVkYTRkOTc1NjI3NWZhNzAzMGZiZDU1ZmE4NzM3MSJ9.kPDklDFq9b6JBa3-OMDRCKOyhlBLgheNs9AywLOxoMA&download=image.png "")

`0666 & ~0002`：-->

`110 110 110` & ~`000 000 010` -->

`110 110 110` & `111 111 101` -->

`110 110 010` 也就是生成的文件的权限664了。

**注意** ：`umask`这种机制的存在，就是为了防止产生权限过松的文件。 可以看出的是，`umask`的值越大，被降的值的权限越低，消掉的权限越多。





## `fgetc()` 和 `fputc()`函数

`man`一下`getchar`如下：

![](https://tcs.teambition.net/storage/3124165c76743b7fe72f59fc439bad551fe5?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTI5MiwiaWF0IjoxNjMwOTMwNDkyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQxNjVjNzY3NDNiN2ZlNzJmNTlmYzQzOWJhZDU1MWZlNSJ9.WWe5ZxW5BT0QadYPHtC9bHbVhVazyntpFiOewjd-Dy0&download=image.png "")

`getchar()`相当于`getc()`， 默认的输入在`stdin`中来的。

`getc()`相当于`fgetc()`函数。



### `getc()`和`fgetc()`的区别是？？

从原始定义上来讲，这<font color='red'>两个函数一个被定义为宏， 一个被定义为函数</font>。

- `getc()`被定义为宏来使用

- `fgetc()`被定义为函数来使用



### 函数和宏的区别：

**注意** ： 内核中链表的实现，通篇都是用宏来实现的，没有任何的函数，为什么要这样做？ 因为内核的实现是在帮你节约一点一滴的时间，**宏不占用你的调用时间，只占用编译时间；函数不占用编译时间，只占用调用时间。**



### 例子

1. 文件`copy`功能

```c
#include <stdio.h>
#include <stdlib.h>
#include <errno.h>

int main(int argc, char **argv)
{
    FILE *source , *target;
    int c;
    // 如果用到了命令行参数，需要先判断命令行参数个数
    if(argc < 3)
    {
        fprintf(stderr, "Usage ... \n");
        exit(1);
    }
    source = fopen(argv[1], "r");
    if(source == NULL)
    {
        perror("source file fopen()");
        exit(1);
    }
    target = fopen(argv[2], "w+");
    if(target == NULL)
    {
        perror("target file fopen()");
        fclose(source); // 如果不关闭source， 会造成内存泄漏
        eixt(1);
    }
    while(1)
    {
        c = fgetc(source); // 返回值是整形，因为有-1的情况存在
        if(c == EOF) break;
        fputc(c, target);
    }

    fclose(source); // 先关闭依赖别人的对象
    fclose(target);
    
    exit(0);
}
```

可以使用`diff`命令查看两个文件是否相同：`diff temp temp1`

如果命令行没有输出，则说明两个文件内容相同

![](https://tcs.teambition.net/storage/3124a239ff8f68de62184c7381272ece2538?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTI5MiwiaWF0IjoxNjMwOTMwNDkyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjRhMjM5ZmY4ZjY4ZGU2MjE4NGM3MzgxMjcyZWNlMjUzOCJ9.gmz5nrk4j8-FzcsLA1cTR5ll8mSfuLoeg5sgzCJxY0g&download=image.png "")





## `fgets()` 和` fputs()`函数

`man`一下`gets()`函数如下：

![](https://tcs.teambition.net/storage/3124be9752edecb48c5e8dd42f6a82fcaa59?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTI5MiwiaWF0IjoxNjMwOTMwNDkyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjRiZTk3NTJlZGVjYjQ4YzVlOGRkNDJmNmE4MmZjYWE1OSJ9.iApXNtbUYWFYtJ4nxhu4jn62xu9MRCsM2kXeWYcz1jM&download=image.png "")

**建议不要使用`gets()`函数，因为它有`Bug`, 它不检查缓冲区的溢出**，我们可以使用`fgets()`代替使用`gets()`。



### 为什么`gets()`函数危险？？

![](https://tcs.teambition.net/storage/31244e872168be8470c0a579a1bbcd998677?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTI5MiwiaWF0IjoxNjMwOTMwNDkyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQ0ZTg3MjE2OGJlODQ3MGMwYTU3OWExYmJjZDk5ODY3NyJ9.GK3R1hmFXJZKbu6GubKlx4c_TyLyPRv8y0iFXGq6IC4&download=image.png "")

因为它只约定了一个地址，从终端上接收来的内容（敲一串字符以后回车， 这行字符串在输入的时候并没有放在指定的地址里面，而是放在了当前的输入缓冲区中，当回车的时候，这串内容才被放到指定的地方去。）

**`gets()`函数不检查缓冲区溢出，这里的缓冲区说的是输入缓冲区，而不是程序的缓冲区**



### `fgets()`

`fgets()`函数的定义：

`char *fgets(char *s, int size, FILE *stream);`

![](https://tcs.teambition.net/storage/31244bbd594e2758745115003909c10f6034?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTI5MiwiaWF0IjoxNjMwOTMwNDkyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQ0YmJkNTk0ZTI3NTg3NDUxMTUwMDM5MDljMTBmNjAzNCJ9.2VjRsWTziCo1LFKGiJsBgI1rUSNFuPDTfdTMj2ihTxc&download=image.png "")

- `size`: 指定的接收大小

- `stream`: 指定的流

- `*s`: 放入的内存地址

- 

**`fgets()`的结束有两种情况：**

1. 读到了`size-1`个字节（最后一个字节留给补全的字符串尾名`\0`： `\0`表示当前串的结束）

1. 读到了`\n`字符





## `fread() `和` fwrite()`函数

**注意：**` linux`环境下不区分二进制流和文本流



### 网络套接字是文件IO操作

**网络套接字抽象出来，是一个文件IO操作，是一个文件描述符。**

`FILE *fdopen(int fd, const char *mode);`

这个函数是使用已有的文件描述符来指定打开方式，从而把它封装成一个流的操作。换句话说，你的套接字`SOCKET`返回给你的文件描述符，但是摇身一变，这个文件描述符就能封装到一个流当中。你对一个`stream`的读写，实际上就是相当于在读写网络套接字。

`Unix`的一句话：**一切皆文件** ， 文件是所有实现的基础，包括一些设备，管道全部抽象的是IO操作





## `printf()`族函数

`man 3 printf`可以看到`printf()`一族的函数

![](https://tcs.teambition.net/storage/3124dbf693421719f520162ed19ef4388362?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTI5MiwiaWF0IjoxNjMwOTMwNDkyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjRkYmY2OTM0MjE3MTlmNTIwMTYyZWQxOWVmNDM4ODM2MiJ9.d-uVoMITM-Ez9BHl-APa2jeiLd6tmfOEMNVlmlLqxsk&download=image.png "")

`int printf(const char *format, ...);printf()`的函数的功能是，将一些函数按照一定的格式输出到stdout上。



**注意** ：建议使用`fprintf()`，因为当前的流该往哪输出就往哪走向，不建议将所有的输出都放在`stdout`上，因为有的时候当命令行的传参不正确的时候，（默认情况下，`stdin`是指向键盘的，`stdout`和`stderr`是指向显示器），我们可以根据自己的需要改变输出位置（把输入输出做重定向），例如，`stderr`指向一个打开的文件。

>  一般来说会把`stderr`输出到一个文件里，因为会发生例如：在跑内核或者大程序的时候，跑的过程当中会出现一些警告，还有一些正常的输出的内容，我们有的时候会需要把这两者分开，把输入重定向到一个文件当中，把标准出错重定向到一个文件当中。

### `sprintf()`

`int sprintf(char *str, const char *format, ...);`

这个函数的功能是把format的内容，输出到一个字符串当中去。



**`atoi()`**函数，把一个串转换成一个整形数。



```c
#include <stdio.h>
#include <stdlib.h>

int main()
{
    char str[] = "123456"; // 如果字符串中间有字符的话，就拿到字母为止。
    printf("%d \n", atoi(str));
    exit(0);
}
```

如果我需要一个将多个不同的数据类型综合成一个串， 使用`sprintf()`，例如:

```c
#include <stdio.h>
#include <stdlib.h>

int main()
{
    int year=2020, month=12, day=11;
    // 想输出2020-12-11
    char buf[1024];
    sprintf(buf, "%d-%d-%d", year, month, day);
    puts(buf);
    exit(0);
}
```





## `scanf()`一族函数

![](https://tcs.teambition.net/storage/31246fd5ad3a91a35da9f42c8ac136b97c63?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTI5MiwiaWF0IjoxNjMwOTMwNDkyLCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQ2ZmQ1YWQzYTkxYTM1ZGE5ZjQyYzhhYzEzNmI5N2M2MyJ9.Liy9SG7VrpaTTDRxvDXJeXylZ2_SFCAZ2c49qpygixE&download=image.png "")

**注意** ：在`scanf()`一族的函数中，一定要慎重使用`%s`， 因为在终端上输入的时候，你是不清楚这个字符串有多长的，所以看不到目标位置有多大，这是`scanf`的缺陷之一。





## `fseek()`,` ftell() `和 `rewind()`文件位置函数

这三个函数是用来操作文件位置指针的。

**注意** ：`fopen()`函数由于打开方式的不同，文件位指针的位置不同，有的在文件首，有的在文件尾。

文件中有个文件位指针，在打开文件进行读写的时候，文件位指针所在的位置称为当前位置。不是每次用的时候都在文件首的位置。所以三个函数是用来解决以上的问题的。



### `fseek()` 定位

`int fseek(FILE *stream, long offset, int whence);`

- `offset`: 偏移量大小

- `whence`: 偏移的相对位置（包含`SEEK_SET`文件首, `SEEK_CUR`文件当前位置, or `SEEK_END`文件尾,）



### `ftell()`

`ftell()`用来显示文件指针当前位置在哪，通常和`fseek()`一起使用。

**例子** ：计算文件有效字符的数量

```php
#include <stdio.h>

int main(int argc, char **argv)
{
    FILE *fp;
    if(argc < 2)
    {
        fprintf(stderr, "Usage...");
        exit(1);
    }
    fp = fopen(argv[1], "r");
    if(fp == NULL)
    {
        perror("fopen()");
        exit(1);
    }
    fseek(fp, 0, SEEK_END); // 将文件位指针指向文件尾
    printf("count = %ld", ftell(fp)); // 打印当前文件指针位置
    fclose(fp);
    exit(0);
}
```



### `rewind()`

`rewind()`将文件指针定位到文件首位置，相当于`fseek(fp, 0, SEEK_SET);`



### 空洞文件

`fseek()`函数经常用来帮助我们完成一个空洞文件（空洞文件种，全部或者一部分位置充斥着字符`\0`, ASCII码为0的字符）。



**空洞文件的用途**：

下载工具：如果下载一个2G的文件，建立下载任务之后，在磁盘上会马上产生一个文件，这个文件不是慢慢涨为2G的文件，而是直接产生的时候就是2G的大小。因为这样不会产生下载到一半的时候产生空间不足的问题。这个文件就是空洞文件，因为要先占上磁盘，下载文件产生之后马上调用	`fseek()`, 将文件指针指向文件尾。这2G空间中全部为`\0`的字符。

下载工具会把空洞文件切成片，用多线程来进行每一小块的数据下载。





## `fflush()`函数

例子

```c
#include <stdio.h>

int main()
{
    printf("Before while()");
    while(1);
    printf("After while()");
    exit(0);
}
```

我们认为当这个程序执行的时候，`"Before while()"`会被打印，但是`"After while()"`不会被打印。但是实际执行的时候`"Before while()"`也不会被打印，为什么？？？

**注意：** 因为`printf()`是典型的行缓冲模式，只有遇到换行的时候，才会刷新缓冲区，或者一行满了来刷新缓冲区。所以不会打印。

> 所以强调如果没有特殊要求，要在`printf()`末尾加上`\n`。还有一种方法是使用`fflush()`强制刷新缓冲区。如果有多个流同时打开调用`fflush(NULL)`时，会将多个流都可以	强制刷新。

### 缓冲区

1. **缓冲区的作用**  :缓冲区的存在大多数情况下时好事，用来合并系统调用。

1. **主要分为**:

    - 行缓冲： 换行的时候刷新，满了的时候刷新, 强制刷新(`fflush`，标准输出是这样的，因为是终端设备)

    - 全缓冲: 满了的时候刷新，强制刷新(默认， 只要不是终端设备)

    - 无缓冲：如`stderr`，需要立即输出的内容



**技巧**： `VIM`视窗模式下， 光标停留在某个函数上，然后按`shift + k` 就会跳转到这个函数的`man`手册中去。然后按两次`q`会回到原来的文件中去。





## `getline()` 函数 获得一行

以上讲的函数并没有一个函数可以取得一行的数据。如何解决这个问题???

(**动态内存实现**)， 先要一块内存，如果这块内存不够了再继续要，一直要到满足内存大小为止。这个功能可以借助`malloc`和`realloc`一系列动态内存函数实现。

`getline()`函数实际上是将上面说的内容封装了一遍。





## 临时文件

临时数据放在临时文件中，

1. 如何创建临时文件才能不去冲突？？？

1. 忘记及时销毁？？



有两个常用函数`tmpnam` 和 `tmpfile`

1. `tmpnam`: `tmpnam`创建临时文件是比较危险的，因为它创建临时文件需要两步,而当前的机器中，并不能做到绝对的一次性执行这两步（产生文件马上创建文件）。所以没有办法创建一个非常安全的临时文件。

1. `tmpfile`: `tmpfile`以二进制读写的形式打开一个临时文件，它产生的临时文件有一个特点：对于一个程序员来说并不关心临时文件的文件名是什么，而`tmpfile`函数会产生一个匿名文件（实际上磁盘上已经创建了这个文件， 但是我们并不知道这个文件存在哪，这种文件称为**匿名文件**），然后`tmpfile`返回给一个指向这个文件的文件指针。



**注意** ：一个文件如果说没有任何的硬链接指向它，而当前的文件的打开计数又已经成为0，那么这块数据就要被释放了，所以不用考虑第二个及时销毁的问题。



**下一章 《**[**系统调用IO**](https://thoughts.teambition.com/share/60694647a86a91004a91d260#title=Linux系统编程笔记（李慧琴）2__系统调用IO)**》**

