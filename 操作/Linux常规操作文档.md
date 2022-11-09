（1）进程相关操作

`gnome-system-monitor` 界面版资源管理器

（2）给文件夹及其文件修改权限

> chmod -R 777 path

![](../assets/2022-11-03-16-39-35-image.png)

其中： r=4,w=2,x=1

（3）添加可执行路径到`PATH`

查看当前`PATH`内容：

> echo $PATH

![](../assets/2022-11-07-11-11-39-image.png)

在`.bash_profile`文件中添加环境变量

![](../assets/2022-11-07-11-14-30-image.png)

（4）修改本地host文件

> vim /etc/hosts

> 127.0.0.1 localhost

> sudo  systemctl restart networkManager  # 重启网络服务

![](../assets/2022-11-07-18-07-22-image.png)

(5)回到上一目录

> cd -(减号)

![](../assets/2022-11-07-18-04-43-image.png)
