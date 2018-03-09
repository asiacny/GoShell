# GoShell

> A command-line tool for a deity | 懒人命令行必备神器

## 联系方式

- E-mail: lauixData@gmail.com
- QQ群：216827114

## 支持平台

- Mac (darwin)
- Linux
- Windows (有BUG)

## 使用场景

- 经常记不住 Linux 命令 ？
- Linux 命令太多了，不想手打 ？
- 想节省写命令的时间 ？
- 。。。
- (那么 GoShell 适合你)

## 使用介绍

正常遇到一个服务启动了很多进程，需要使用(比如Python)

**`ps -ef | grep python | awk '{print "kill -9 "$2}' | sh`**

身为懒人感觉好麻烦。

使用 GoShell  **`gs kall python`** 即可，省了打很多命令

## 安装软件

下载 GoShell

复制 gs 到 /usr/local/bin or ln -s 设置软链接

(别忘记给权限)

**初始化**

**`gs init`**

**帮助**

**`gs help`**

**命令帮助**

**`gs kall --help`**

## 编辑命令

支持自定义命令，编写配置文件 默认路径： /etc/goshell/shell.conf


```
# gs 参数名
[kall] 
# 需要执行的命令 {{name}} 是需要传的参数
command = ps -ef | grep {{name}} | awk '{print "kill -9 "$2}' | sh 
# 命令说明
introduce = Kill 全部相关进程，自定义参数 key 为服务名 
```

## 使用演示

![](http://ww1.sinaimg.cn/large/005Bpb8ily1fp6f5bqaw7g30h40anncg.gif)
