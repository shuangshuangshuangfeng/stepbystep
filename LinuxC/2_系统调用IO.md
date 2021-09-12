

## 02 系统调用IO介绍

`I/O`：`input & output`， 是一切实现的基础，主要分为

- `stdio` 标准IO （所有的标准都是为了和稀泥整合不同）

- `sysio` 系统调用IO

**注意**：系统调用IO又称为文件IO， 文件描述符（`fd`, `File discriptor` 文件描述符）是文件IO中贯穿始终的类型。



当前章需要讨论的问题：

- 文件描述符的概念（整型数，数组下标，文件描述符优先使用当前可用范围最小的一个）

- 文件IO操作：`open` , `close`,`read`, `write`,`lseek`

- 文件IO与标准IO的区别

- IO效率问题

- 文件共享

- 原子操作

- 程序中的重定向：`dup`,`dup2`

- 同步：`sync`, `fsync`, `fdatasync`

- `fcntl `管家级函数

- `ioctl` 管家级函数

- 虚目录 `/dev/fd/`



系统调用IO是支持标准IO的，也就是说标准IO的实现是依赖于系统调用IO来实现的。比如说 `fopen`依赖于`open`， `fclose`依赖于`close`。





## 文件描述符

在标准IO当中，有一个类型贯穿始终，是`FILE`，它其实是一个结构体（我们叫他`FILE`结构体），我们知道，它里面一定有一个文件位指针`position`。

在系统调用IO中，一个文件的唯一标识，叫做`inode`， 每个文件有一个唯一的`inode`号。当我们打开一个文件的时候，它会产生一个结构体，这个结构体存的是这个文件的几乎所有属性信息，也就是我们当前要用的所有信息。当我们拿着一个指向这个结构体的指针的时候，就相当于拿着一个类似于指向FILE结构体的指针，但是系统并没有给这样做，而是将数据结构隐藏，不让用户知道定义的结构体的类型是什么。

于是定义了一个数组，这个数组中存的是指向结构体的指针，然后交给用户的是当前指向数组指针的下表`index`，也就是一个整型数。

**注意**：当用户想访问一个文件，FILE结构体中一定会有一个整型的文件描述符，然后用户拿着这个整型数过来，可以找到指向结构体的指针，然后根据指向结构体的指针找到结构体，然后根据结构体中的`position`访问文件内容。



### `fopen`和`fclose`时做了什么操作？？

当调用标准IO`fopen`函数的时候，会产生一个FILE类型结构体，然后系统调用IO调用`open`函数创建一个指向文件的结构体，将这个结构体的指针放在数组中，然后将数组下标给FILE结构体。

当调用标准IO`fclose`函数的时候，会先调用系统调用IO的`close`函数，来free指向文件的结构体，然后再freeFile结构体。



### 存放结构体指针的数组多大？？

上一章中，写过一个小例子，用来测试一个文件中最多可以打开的文件流的数量。当时测出的是`1024（1021+3）`个。

**注意**：如果通过`ulimit -n`命令来更改文件中能打开的文件的个数的话，其实就是在更改这个数组的大小。**默认情况下** **数组中的下表0,1,2关联的是stdin, stdout, stderr三个设备** 。实际上是可以改的，因为看**0,1,2是从父亲进程那里继承过来的** ，如果他的父亲没有0,1,2，或者说父亲在创建当前进程之前把0，1，2关掉了或者重定向了，当前进程的0，1，2关联的也不会是标准设备。

存放结构体指针的数组存在于进程空间中的，每一个进程都会有一个这样的数组。

当一个文件在一个进程中被打开多次的时候，每次打开都会产生一个指向文件的结构体，如果按照一定的协议，可以对一个文件进行协同操作。

当数组中是可以多个下表关联同一个指向文件的结构体的，当其中一个指针调用`close`的时候，是不会`free`这个结构体的，所以这个指向文件的结构体中一定会有一个计数器，用来表示当前结构体被几个指针引用着。当这个计数器减为0的时候，这个结构体才会被free掉。





## 文件IO操作相关函数

### `open() `函数

若获得一些文件描述符的话，先执行`open`

![](https://tcs.teambition.net/storage/31246f79f6d5ae4cbf91d7e273abb825590e?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQ2Zjc5ZjZkNWFlNGNiZjkxZDdlMjczYWJiODI1NTkwZSJ9.CQ-6rCVoBk-OADmtGYxQRT3W55eqZ0aJMBgMKxN13w4&download=image.png "")

如果成功，则返回一个文件描述符，如果失败则返回-1 。



**注意**：`buffer`需要理解成写的缓冲区，`catch`可以理解成读的缓冲区。



后面讲时间专题的时候，会讲到三个时间：

- `ATIME`， 最后读的时间

- `MTIME`， 最后写的时间

- `CTIME` ,   最后亚数据修改的时间



### **使用**`fopen`**函数的时候，有几种权限， 放到系统调用IO中权限应该怎么写？？？**：

- `r `: `O_RDONLY`， 要求文件存在

- `r+`：`O_RDWR`， 首先是读写，但是要求文件存在

- `w`：`O_WRONLY`(只写) |` O_CREATE`(文件不存在的话要创建) |` O_TRUNK`（文件存在的话要截断）

- `w+`：`O_RDWR`（读写）| `O_TRUNK`（有则清空）| `O_CREATE`（无则创建）



在上一章中，使用`fopen`函数去创建文件的时候，是遵循一定的规则去创建文件权限的，在系统调用IO中创建文件，它的权限是？？

![](https://tcs.teambition.net/storage/3124649d5e2f4740c678815f262e7c990671?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQ2NDlkNWUyZjQ3NDBjNjc4ODE1ZjI2MmU3Yzk5MDY3MSJ9.PAqNYI20y4M_1xaj_YMiCWAxJpjuGV-nA3PjnSI6mxA&download=image.png "")

如果`flag`当中有`CREATE`， 那么一定要用三参的形式，如果`flag`当中没有`CREATE`， 那么一定要有两参的形式。

open的第三个参数是给的权限，当然不是你给多少就有多少，而是用你给的值按位与`umask`取反， 也就是`mode & ~umask`，和上一章的公式一样。(`umask`: 默认权限)



### 定参和变参

上面`man open`中，可以看到`open`函数有两种实现方式，一种有两个参数，一种有三个参数，有什么区别？？

> 如果函数名相同，函数的参数不相同的时候，这种现象叫做**重载**。但是C语言中没有重载，因为重载是定参的 ，那么这是如何实现的？？**和**`printf()`**的实现相同，`printf`叫做变参函数，所以open是用变参来实现的** 。

**注意**：`gcc`有一个选项叫`gcc -Wall` 是用来显示当前所有的警告，我们可以把`-Wall`加到`makefile` 文件中去，因为` gcc`有本身它的提醒会比较圆滑，有时间它告诉你这是一个`warning`， 但是实际上已经是一个`error`了。所以一定要把程序调到没有警告为止，除非你可以解释这个警告。



### `read() `， `write() `和` lseek()`函数

`man open` 的结果如下：

![](https://tcs.teambition.net/storage/3124a650884055d61256dcf7a84d27311022?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjRhNjUwODg0MDU1ZDYxMjU2ZGNmN2E4NGQyNzMxMTAyMiJ9.gKpTJ1Oek9E4xyl_yxXeF9AFigiZ-zWywh_jCyYEvus&download=image.png "")

**注意**：`fd`是一个文件描述符，根据一切皆文件的描述，这个文件描述符可以是指向文件的， 也可以是指向设备的。

`buf`是要读取的缓冲区



`man write`如下：

![](https://tcs.teambition.net/storage/3124a7da2d93d5f51e41b34ee3958860767f?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjRhN2RhMmQ5M2Q1ZjUxZTQxYjM0ZWUzOTU4ODYwNzY3ZiJ9.fY8ml1Tb2qEak_rin0fGQJcyX1tMWCkMjNIAD7iBAbI&download=image.png "")

要注意这里的`buf`前的类型是`const`，说明这块内存的地址是不能修改的，而读的话是可以在任意地址读的。

`wirite`的返回值：

![](https://tcs.teambition.net/storage/3124f844b0031d89445d2aed38e627a1516f?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjRmODQ0YjAwMzFkODk0NDVkMmFlZDM4ZTYyN2ExNTE2ZiJ9.JxdJz9cb_iqBOreKAvtf46sOs0iRn0t1uz2QPPo3WSI&download=image.png "")

如果成功的话，返回字节的个数（注意这个个数可以是0， 也可以是大于0的数），如果不成功，则会返回-1。如果返回值为0的话，表示什么也没写进去，它并不表述出错，大概是由于这些errors中有真错有假错，如果发生了假错，实际上是因为某些阻塞或者非阻塞的事件打断了。



`man lseek `如下：

![](https://tcs.teambition.net/storage/31245a140f9621133e98808f520d2c246db7?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQ1YTE0MGY5NjIxMTMzZTk4ODA4ZjUyMGQyYzI0NmRiNyJ9.E7FtGGG0mEwSPjLXnRzKoSNdq6n8p8PmVLX2hx6NqX4&download=image.png "")

可以发现，和`stdio`中的`fseek`函数是类似的。实际上，`fseek`是通过调用`lseek`实现的。

- `fd `文件描述符

- `offset` 偏移量

- `whence `相对偏移位置，（ 标红的）



### 例子

仍然是`mycopy`的例子

```c
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>


#define BUFSIZE 1024

int main(int argc, char **argv)
{
	int source_fd;
	int target_fd;
	char buf[BUFSIZE];

	if(argc < 3)
    {
		fprintf(stderr, "Usage ... \n");
        exit(1);
	}

    source_fd = open(argv[1], O_RDONLY);
    if(source_fd < 0)
    {
		perror("open()");
        exit(1);
    }

	target_fd = open(argv[2], O_WRONLY|O_CREAT, O_TRUNC, 0600);
    if(target_fd < 0)
    {
		close(source_fd);
		perror("open()");
		exit(1);
    }

    int len, ret, pos;
    while(1)
    {
		len = read(source_fd, buf, BUFSIZE);
        if(len < 0)
        {
			perror("read");
            break;
        }
        if(len == 0)
        {
            break; // 表述已经读完了
        }
		pos = 0;
		// 坚持读多少一定要写多少
		while(len > 0) // 考虑到有没有全部写完的情况
		{
			ret = write(target_fd, buf, len);
	        if(ret < 0)
	        {
	            perror("write");
	            exit(1);
	        }
			pos += ret;
			len -n ret;
		}	
   
	}


    close(target_fd);
    close(source_fd);
    exit(0);

}


```

向一个文件写内容，可以使用`echo '123424' > temp`， 是将123424写入到temp文件中。





## 标准IO和系统调用IO的区别

1. 第一贯穿始终的类型

    - 标准IO中是FILE类型

    - 系统调用IO中是int类型数组下标

1. 响应速度&吞吐量 ， 因为 缓冲机制（例如，输出操作）

    - 标准IO具有缓冲的机制，看上去已经操作了，实际上是放到了输出的缓冲区当中，比如调用`fflush`函数时才会写入， 所以其吞吐量大。

    - 系统调用IO每调用一次就会实打实的从user态切换到kernel态去执行一次，实时性很高，会立即输出，响应速度快，吞吐量小。



**面试题** ：如何使一个程序变快？？？ 可以问一下吞吐量上还是响应速度上变快？

如果响应速度快，那就多用系统调用IO， 如果吞吐量大，那就多用标准IO。其实从用户角度来讲，用户体验觉得程序变快了，其实不是响应速度，而是吞吐量。

**注意**：在系统调用IO和标准IO两者都能解决问题的情况下，推荐尽量使用标准IO。

**提醒**： 标准IO与系统调用IO不可混用。



### 为什么标准IO和系统调用IO不能混用呢？？

比如FILE类型中有一个`position`， 然后文件描述符的结构体中也有一个`position`， 那么这两个是一样的吗？？答案是**往往一定不一样** 。

因为标准IO有缓冲机制。

```c
FILE *pf;
fputc(fp);  -> position++  // 所以FILE中的positon是向后移动的
fputc(fp);  -> position++
```

这里可以看到FILE类型中的`position`是进行了一个+2的操作，那么文件描述符指向的结构体中的`position`是否也`+2`？？实际上不是的，**因为当**`fputc()`的时候，实际上并没有写入到磁盘上，而是写到了缓冲区当中 **，所以文件描述符指向的结构体中的**`position`并没有**`+2`**，只有当各种刷新的时候，文件描述符指向的结构体中的**`position`**才会+2 。



### **例子**

```c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main()
{
	putchar('a');
	write(1, "b", 1);

	putchar('a');
	write(1, "b", 1);

	putchar('a');
	write(1, "b", 1);

	exit(0);
}
```

结果是：

![](https://tcs.teambition.net/storage/31248283cbb31be5896ed4645e95fd632c53?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQ4MjgzY2JiMzFiZTU4OTZlZDQ2NDVlOTVmZDYzMmM1MyJ9.rJhFsQPiIrI99pSK6A8WiirJ6IFTOaY6d1pILPC9aLk&download=image.png "")



**注意**：`strace ./print `命令可以帮助我们看一个可执行文件(print)的系统调用是如何发生的。

![](https://tcs.teambition.net/storage/312478c66b6588564116c1cd480e40cff086?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQ3OGM2NmI2NTg4NTY0MTE2YzFjZDQ4MGU0MGNmZjA4NiJ9.nXlVPnYwDNngxL39DoYH2MznZUiIjuvVPpW7P73T5ro&download=image.png "")

这里可以看到，先写入了三次`b`字符。后面一次写入了`aaa`字符串，系统调用IO中的`write`支撑起了标准IO中的所有写入函数，所以三次`putchar`相当于一次`write`来实现的。所以看到的效果是"bbbaaa"。





## IO效率问题

**一个课下习题**：自己写的`mycopy`修改`BUFSIZE`的大小，从`128B `二倍增长（按理说性能是一直增的，但是增到一个拐点性能肯定会下降），增长到16M， 复制的文件大约3G-5G大小之间。然后出一个表，当`BUFSIZE=128`时，`real time` , `sys time`, `user time`的时间， 然后注意性能的最高拐点出现在`BUFSIZE`多大的时候，以及程序何时会出问题。

一个命令`time ./mycopy /etc/services /tmp/out  `，用来测试后面这个操作执行了多长时间。

![](https://tcs.teambition.net/storage/3124afcdb64319e92015337abc1908ca695b?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjRhZmNkYjY0MzE5ZTkyMDE1MzM3YWJjMTkwOGNhNjk1YiJ9.Qo7kMU4WnyV2hNdz4tK4o72ruVV-s850OdwCmhH8Gpk&download=image.png "")

-  `real time`: 理论上的值是`user time + sys time + 一点点时间`， 这个一点点时间是调度等待的时间

- `user time`: 当前操作在user层面消耗的时间

- `sys time`: 当前操作在系统调用层面，或者说在`kernel`层面所消耗的时间。

| BUFZISE | real time | user time | sys time   | info |
| ------- | --------- | --------- | ---------- | ---- |
| 128     | 3m51.615s | 0m2.448s  | 0m39.721s  |      |
| 256     | 3m53.279s | 0m2.307s  | 0m40.413s  |      |
| 512     | 7m5.793s  | 0m3.870s  | 1m16.263s  |      |
| 1024    | 4m2.589s  | 0m2.852s  | 0m40.289s  |      |
| 2048    | 1m42.300s | 0m0.891s  | 0m.17.968s |      |
| 4096    | 0m50.092s | 0m0.318s  | 0m8.298s   |      |
| 8192    | 0m24.896s | 0m0.156s  | 0m4.299s   |      |
| 16384   | 0m12.498s | 0m0.050s  | 0m2.185s   |      |
| 32768   | 0m7.038s  | 0m0.000s  | 0m1.216s   |      |
| 65536   | 0m7.028s  | 0m0.057s  | 0m1.142s   |      |
| 131072  | 0m5.323s  | 0m0.040s  | 0m0.804s   |      |

(文件大约1G, 电脑比较慢。。使用得是Windows10的子系统`Linux`.) **这个结果和老师讲得不一样，差很多，然后并 没找得到拐点。。。先留这个问题。**

理论上来说，出现的这个拐点正好是扇区数（`block_size`）的整倍数



## 文件共享：多个任务共同操作一个文件或者协同完成任务

面试题：写程序删除一个文件的第10行





## 原子操作

原子：不可分割的最小单位

原子操作：不可分割的操作

原子操作的作用：解决竞争和冲突

如：`tmpnam`函数，是不适合用来创建临时文件的，因为它的操作不原子，容易出现冲突。



## 程序中的重定向: dup, dup2

![](https://tcs.teambition.net/storage/31249f6aa8e54874ca23a4850a8e5804110c?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQ5ZjZhYThlNTQ4NzRjYTIzYTQ4NTBhOGU1ODA0MTEwYyJ9.76E0IpjTI3pZl4V_lLDLk8DwZTvLUp2_DzBI9k54FYA&download=image.png "")

### dup

`dup` 函数，将一个文件描述符放在当前可用范围内最小的数组index中。

例子如下：

```c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

#define FNAME "temp"

// 没有简单的程序，只有头脑简单的程序员!
int main()
{

        // 当程序s执行的时候，hello没有在终端中，而是在一个文件中
        // 以下是操作

        // 文件描述符优先使用当前可用范围内最小的内容
        int fd;
        
        fd = open(FNAME, O_CREAT|O_WRONLY|O_TRUNC, 0600);
        if(fd < 0)
        {
                perror("打开失败");
                exit(1);
        }
		
		close(1); // 关闭终端
        dup(fd); // 将当前的文件描述符放在1上，然后相当于将输出重定向
		close(fd);

        /**************************************/
        puts("hello!");
        printf("heheheheh.....\n");

        exit(0);
}

```

以上代码会有以下问题：

1. 如果本身当前程序中文件描述符没有1， 那么，会导致fd本身就会等于1， 然后后面再dup再关闭，会出现问题

2. 如果旁边还有一个程序在跑，如果两个程序公用一个文件描述符表，一旦执行了close(1)， 还没来得及dup(fd), 另一个程序开了一个文件，那么会占用文件描述符为1的位置。

**之所以会出现上面的问题2，是因为下面的两个操作并没有原子**

```c
close(1); // 关闭终端
dup(fd); // 将当前的文件描述符放在1上，然后相当于将输出重定向
```



### dup2

所以有了函数`dup2`。它的输入有两个文件标识符，一个新的，一个旧的。将旧的fd放到新的fd上去。

![](https://tcs.teambition.net/storage/3124c7baa68333e5c47560fd580733b86868?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjRjN2JhYTY4MzMzZTVjNDc1NjBmZDU4MDczM2I4Njg2OCJ9.CXk8_44ilI9YckoyZ4EaetUdVTrvfA3aGf4Ot3d6Pr4&download=image.png "")

```c
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

#define FNAME "temp"

// 没有简单的程序，只有头脑简单的程序员!
int main()
{
        int fd;
        fd = open(FNAME, O_CREAT|O_WRONLY|O_TRUNC, 0600);
        if(fd < 0)
        {
                perror("打开失败");
                exit(1);
        }
		
        dup(fd, 1); // 将当前的文件描述符放在1上，然后相当于将输出重定向
		if(fd == 1): // TODO: 如果fd本身就是1, dup2什么也不会做，所以也不需要关闭
			close(fd);

        /**************************************/
        puts("hello!");
        printf("heheheheh.....\n");

        exit(0);
}
```

在你完成一个模块的时候，最起码保持程序进入模块前是什么状态，程序结束以后还应该是什么状态，如果把输出写入了一个文件当中，就应该在模块结束以后将标准输出再恢复回去。

**所以在编程时需要注意几点：**

- 不要有内存泄漏

- 不要有写越界的现象

- 不要当作自己在写`main`函数，永远要当作自己在写一个小模块





## 同步：`sync`, `fsync`, `fdatasync`

`sync`函数指的是：同步`buffer` 和` cache`, 指的是同步内核层面的`buffer`和`cache`。通常什么时候会做这样的操作？？在解除设备挂载的时候。需要把当前正在`buffer`或者`cache`当中还没有进行同步的数据刷新以下，会用到sync， 然后下一步解除设备挂载。

![](https://tcs.teambition.net/storage/3124180d13d9fa9554fca11da154271d3593?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQxODBkMTNkOWZhOTU1NGZjYTExZGExNTQyNzFkMzU5MyJ9.XU9z8dlx8iPdnPFaS0FRWu7TBFbnT_AEJBfoCTFti-8&download=image.png "")



`man fsync `结果：

![](https://tcs.teambition.net/storage/31248750c706cc49b3599e1e7a5964afbe53?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQ4NzUwYzcwNmNjNDliMzU5OWUxZTdhNTk2NGFmYmU1MyJ9.x88k8QSFn9e2ROgqiF0wFkyd3M7Bicx81NCsTjyUrh4&download=image.png "")

- `fsync` 指的是同步一个文件的buffer或cache

- `fdatasync`指的是只刷数据，不刷亚数据。



**什么叫只刷数据不刷亚数据？？**

数据指的是一个文件当中的有效内容。亚数据指的是文件最后的修改时间、文件的属性等的一些数据。





## `fcntl` 和 `iocntl`

### `fcntl`: 文件描述符所变得魔术几乎都来源于该函数

![](https://tcs.teambition.net/storage/3124c8c926480fde6665e0f315adef239135?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjRjOGM5MjY0ODBmZGU2NjY1ZTBmMzE1YWRlZjIzOTEzNSJ9.eLehsIkG6OQuiOAnf6c8NwFhXkFsHWDkcnl5zwNR_xw&download=image.png "")

它的参数有:

- `cmd:` 命令

- `arg`: 参数

什么是文件描述符相关的魔术？

由于这个函数的`cmd`不同，所以它的参数`arg`不同，它的返回值也不同。因为这个函数的功能比较杂，所以称它为管家级别的函数。



### `ioctl`: 设备相关的内容都归它管

![](https://tcs.teambition.net/storage/3124271d538c6555d30ffcda8f8bd45daf5a?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQyNzFkNTM4YzY1NTVkMzBmZmNkYThmOGJkNDVkYWY1YSJ9.YscZxeED-HaaSobfSXMlXV1MMchSMtLhhvUhlrLVkjc&download=image.png "")

用于控制设备，是一切皆文件的光环下的垃圾堆，乱七八糟全弄里面了。。





## `/dev/fd/`目录： 虚目录,显示的是当前进程的文件描述符信息

底下是一些`link`， 指向当前的标准的设备

![](https://tcs.teambition.net/storage/31245db7151bc9a3c9c22ed2b75bba8a4b21?Signature=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBJRCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9hcHBJZCI6IjU5Mzc3MGZmODM5NjMyMDAyZTAzNThmMSIsIl9vcmdhbml6YXRpb25JZCI6IiIsImV4cCI6MTYzMTUzNTM3NSwiaWF0IjoxNjMwOTMwNTc1LCJyZXNvdXJjZSI6Ii9zdG9yYWdlLzMxMjQ1ZGI3MTUxYmM5YTNjOWMyMmVkMmI3NWJiYThhNGIyMSJ9.nIgbYBjdRGSbieyDn-wo-lxM8TYt5wgbkT7yNUR-9KQ&download=image.png "")







宏观编程能力不行，微观编程能力还行，俗称**码驴...**













