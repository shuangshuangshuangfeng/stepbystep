# DNS解析过程

![image-20210915145301459](DNS解析过程.assets/image-20210915145301459.png)





**流程是：**

访问递归DNS服务器，如果没有，然后访问一级域名服务器，如果没有，继续访问二级域名服务器，如果没有，则访问三级域名服务器，最后返回主机IP。

