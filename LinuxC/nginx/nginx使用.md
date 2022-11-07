（1）重启命令：【热部署】

> nginx -s reload

![](../../assets/2022-11-02-16-13-27-image.png)

所谓热部署，就是配置文件nginx.conf修改后，不需要stop Nginx，不需要中断请求，就能让配置文件生效！（nginx -s reload 重新加载/nginx -t检查配置/nginx -s stop）


