## 一步一步编译JDK

### 1. 文档

首先，官方文档如下：

```bash
https://github.com/openjdk/jdk/blob/master/doc/building.md
```

1. 安装开发人员包

```bash
sudo apt-get install build-essential
```

2. 进行配置

```bash
sudo bash configure
```

------

## 2. 本地编译与调试JDK环境初步搭建

1. 本地准备一个JDK（**本地JDK版本是需要编译的JDK前一个版本(例如：编译JDK18，本地准备JDK17)**）

2. 下载openjdk源码（注意： **下载openjdk中打过tag的源码**）

3. 进入JDK源码目录
   
   ```shell
   bash configure --with-debug-level=slowdebug --with-jvm-variants=server  --disable-warnings-as-errors
   ```
   
   ![](../../assets/2023-01-03-14-06-55-image.png)
   
   在这个过程中，如果有些依赖没有被安装，会不断的中断，例如下图：
   
   ![](../../assets/2023-02-01-14-46-24-image.png)
   
   然后需要根据提示，安装不同的依赖。
   
   会有如下显示，则表示配置成功。
   
   ![](../../assets/2023-02-01-14-48-39-image.png)
   
   然后进行编译，并等待...
   
   ```shell
   sudo make all
   ```
   
   ![](../../assets/2023-02-01-14-51-06-image.png)

4. `make`完成后，在JDK根目录下会有一个build文件夹，其中则是编译好的JDK。

![](../../assets/2023-02-01-14-53-14-image.png)

5. 使用vs code 打开源码目录(换行后再按下换行就对齐了)

![](../../assets/2023-01-04-09-31-57-image.png)

6. 在根目录下创建.vscode文件夹，然后创建launch.json文件

在launch.json文件中填写：（鼠标放上配置后有描述）

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "jdk19s-src",
            "type": "cppdbg",
            "request": "launch",
            "program": "/home/tina/workspace/xuexi/jdktest/jdk-jdk-19-36/build/linux-x86_64-server-slowdebug/jdk/bin/java", //要运行的目标程序
            "args": [
                "HelloWord" //要运行的.class文件的包名
            ],
            "stopAtEntry": true,
            "cwd": "${workspaceFolder}/java_src", // 鼠标放上去有惊喜,目标的工作目录
            "environment": [ //配置环境变量
                // {
                //     "name": "CLASSPATH",
                //     "value": "/home/tina/workspace/stepbystep/LeetCode/HelloWorld/out/production"
                // },
            ],
            "setupCommands": [
                {
                    "text": "handle SIGSEGV pass noprint nostop", 
                    "description": "ignore SIGSEGV", 
                    "ignoreFailures": true 
                }
            ],
            "externalConsole": false,
            "MIMode": "gdb",
            // "preLaunchTask": "jdk11s-build"  //运行编译任务
        }
    ]
}
```

7. 调试一下
