# Ubuntu 下解决端口占用

1. 查看端口占用情况
   
   ```
   netstat -npl | grep 8081
   （并非所有进程都能被检测到，所有非本用户的进程信息将不会显示，如果想看到所有信息，则必须切换到 root 用户）
   tcp        5      0 127.0.0.1:8081          0.0.0.0:*               LISTEN      401218/node         
   ```

2. 查看端口占用的进程有哪些
   
   ```
   sudo lsof -i:8081
   COMMAND    PID USER   FD   TYPE  DEVICE SIZE/OFF NODE NAME
   chrome  394554 home   28u  IPv4 2437381      0t0  TCP localhost:58632->localhost:tproxy (ESTABLISHED)
   chrome  394554 home   45u  IPv4 2438249      0t0  TCP localhost:50288->localhost:tproxy (ESTABLISHED)
   node    401218 home   22u  IPv4 2162176      0t0  TCP localhost:tproxy (LISTEN)
   node    401218 home   29u  IPv4 2437319      0t0  TCP localhost:tproxy->localhost:50288 (ESTABLISHED)
   ```

3. 杀死进程

```
sudo kill -9 401218
```
