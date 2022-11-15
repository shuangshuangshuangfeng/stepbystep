## 1. docker 的基本使用流程

（1）新建一个新的容器并运行

> docker run -it ubuntu:22.04 /bin/bash

（2）查看所有容器

> docker ps -a

![](../../assets/2022-10-25-11-04-56-image.png)

（3）启动一个容器

> docker container start id

> docker exec -it id /bin/bash

![](../../assets/2022-10-25-11-08-47-image.png)

（4）查看正在运行的容器

> docker container ls

![](../../assets/2022-10-25-11-07-29-image.png)

> docker ps

![](../../assets/2022-10-25-11-52-29-image.png)

（5）进入正在运行的容器：

> docker attach containerID

![](../../assets/2022-11-01-13-59-42-image.png)

（6）删除容器

> docker rm -f  id

（7）关闭容器

> docker container stop id

![](../../assets/2022-10-25-11-49-53-image.png)

## 2. 外部访问容器

:star2:  docker 容器如果在启动的时候，如果不指定端口映射的参数，容器外部不能通过网络来访问容器内部。

（1）将容器80端口映射到宿主机的8000端口

> docker run -it -p 8000:80  id /bin/bash

（2）查看端口映射配置

> docker port id

![](../../assets/2022-10-25-12-03-30-image.png)

## 3. 挂载外部文件或目录

（1）单目录挂载

> docker run -it -v /宿主机目录 : /容器目录

将当前目录挂载到容器的`/usr/local/src`目录上：

![](../../assets/2022-10-31-09-45-02-image.png)

![](../../assets/2022-10-31-09-46-15-image.png)

需要注意的是： 当宿主机对文件进行修改，那么容器中的文件也会被修改，被称为：**双向数据同步**。

## 4. 从镜像仓库拉取指定镜像

> docker pull  [OPTIONS] NAME[:TAG|@DIGEST]

![](../../assets/2022-10-31-09-55-14-image.png)

## 5. docker 免sudo登录

**文件权限：`drwxrwxrwx`**

- 第一: 文件类型;

- 前三位：属主权限；

- 中三位：属组权限；

- 后三位：其他人权限;

（1）搜索docker执行文件：

> sudo find / -name 'docker' 2>/dev/null

![](../../assets/2022-10-31-10-32-10-image.png)

（2）查看docker执行文件权限

![](../../assets/2022-10-31-10-31-35-image.png)

（3）docker其属组为docker

给docker属组添加当前用户：

> docker gpasswd -a ${USER} docker

系统重启生效。

## 6. dockerfile 执行

> docker build -t <image_name> -f <dockerfile_name> <dockerfile_path>

![](../../assets/2022-10-31-17-48-29-image.png)

写一个Dockerfile

```shell
FROM ubuntu:22.04

RUN echo '这里是nginx'

RUN cd ~ 

RUN touch aaa.txt
```

然后执行dockerfile (记得加上最后的那个点)

> docker build -t shuang .   

![](../../assets/2022-11-14-14-20-52-image.png)

`--->`后面是镜像的ID

进入镜像去看一眼：

> docker run -it id /bin/bash

![](../../assets/2022-11-14-14-23-17-image.png)

可以看到已经创建了aaa.txt这个文件
