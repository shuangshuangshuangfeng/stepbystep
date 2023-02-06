# Shell

## 1. 介绍

Shell既是一种脚本编程语言，又是一个连接内核和用户的软件。

## 2. 使用

（1）查看系统当前Shell ：

> echo $SHELL

![](../../../../assets/2023-02-06-09-59-23-image.png)

（2）当前系统有效的登录shell

> cat /etc/shells

![](../../../../assets/2023-02-06-10-02-20-image.png)

（3）修改当前默认shell

> chsh -s /bin/sh


