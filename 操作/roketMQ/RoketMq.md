**启动namesrv**（nameserver 理解成zookeeper的效果）

后台执行，不要结束进程

> nohup sh mqnamesrv

**创建broker**(broker可以理解成RocketMQ本身)

后台执行，不要结束进程

> nohup sh mqbroker



自动创建Topic

> nohup sh mqbroker -n localhost:9876 autoCreateTopicEnable=true



**报错：**

```
org.apache.rocketmq.client.exception.MQClientException: No route info of this topic: TopicTest
See http://rocketmq.apache.org/docs/faq/ for further details.
```

可能是nameserver中没有成功注册broker， 检查broker日志

> tail -f ~/logs/rocketmqlogs/broker.log

![](../../assets/2023-09-19-11-27-34-image.png)

表明注册成功， 再试一下，即可以向mq中发送信息了~

![](../../assets/2023-09-19-11-28-22-image.png)



参考：

[RocketMQ 解决 No route info of this topic 异常步骤_IT_农厂的博客-CSDN博客](https://blog.csdn.net/chenaima1314/article/details/79403113)



[RocketMQ入门到入土（一）新手也能看懂的原理和实战！ - 开发者头条](https://toutiao.io/posts/jluhew1/preview)


