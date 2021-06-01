# 微云本地播放代理工具

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [微云本地播放代理工具](#微云本地播放代理工具)
  - [介绍](#介绍)
  - [使用方法](#使用方法)
    - [下载源码](#下载源码)
    - [登陆微云获取相关信息](#登陆微云获取相关信息)
  - [TODO](#todo)

<!-- /code_chunk_output -->

## 介绍

百度云的在线播放我就不说了吧，感觉很难用，而且完全不可能原画质播放的。
微云的在线播放可以顶满带宽，而且会员只要10块即可。

但是微云这方面并没有完善，主要有两个缺点：

1. 没办法加载字幕，很多视频没字幕没法看的= =。
2. `mkv`原画质播放会没有声音。

第一个问题比较好解决，我写了一个油猴的脚本[http://onns.xyz/js/weiyun-sub.user.js](http://onns.xyz/js/weiyun-sub.user.js)。

第二个问题是由很多历史原因组成的，`chrome`内基本无解。要么使用微云提供的客户端来播放（依然无法加载字幕），要么只能通过第三方工具来打开视频。但是`VLC`之类的播放器本身不是不支持`cookie`的，所以不行，解决办法只有加一层代理，将需要`cookie`校验的视频下载链接通过代理来屏蔽。

感谢[https://github.com/hezhizheng/go-reverse-proxy/](https://github.com/hezhizheng/go-reverse-proxy/)。

## 使用方法

### 下载源码

可以从[https://github.com/onns/weiyun-proxy](https://github.com/onns/weiyun-proxy)处下载自行编译或者是直接下载我编译好的[源码](https://github.com/onns/weiyun-proxy/releases/)。

### 登陆微云获取相关信息

1. 打开微云后，`cmd+alt+i`打开`开发者模式`，选择`Network`标签。
2. 点击图示里的`清除`按钮，清空当前的记录信息（方便找到我们需要的那条记录）。

![Screen Shot 2021-06-01 at 16 16 42](https://user-images.githubusercontent.com/16622934/120291427-e0d89b00-c2f5-11eb-9505-f30cbee5d6bc.png)

3. 选择需要本地播放的那个视频，点击下载按钮。

![Screen Shot 2021-06-01 at 16 16 56](https://user-images.githubusercontent.com/16622934/120291450-e504b880-c2f5-11eb-990e-25cb09a72630.png)

4. 左下角会有一个正在下载的视频，然后记录里也有这个正在下载的信息记录（`序号2`），将`Request URL`和`Cookie`分别复制到**config.json**里的`url`和`cookie`中。

```json
{
  "url": "********************************",
  "port": ":1996",
  "cookie": "********************************"
}
```

![Screen Shot 2021-06-01 at 16 17 09](https://user-images.githubusercontent.com/16622934/120291800-49c01300-c2f6-11eb-92e3-57042b7e0e8f.png)
![Screen Shot 2021-06-01 at 16 19 05](https://user-images.githubusercontent.com/16622934/120291933-68260e80-c2f6-11eb-81d4-fa3aa689375f.png)

5. 打开命令行，到相应的目录下，运行`weiyun-video-proxy`（不同系统的命令不同）。

![Screen Shot 2021-06-01 at 16 18 46](https://user-images.githubusercontent.com/16622934/120292006-796f1b00-c2f6-11eb-8a5e-5da13e02fab5.png)

6. 打开`VLC Media Player`，`File` -> `Open Network`。

![Screen Shot 2021-06-01 at 16 29 26](https://user-images.githubusercontent.com/16622934/120292269-c18e3d80-c2f6-11eb-8579-537d35930b64.png)

7. 填入`http://127.0.0.1:1996/`（`端口`是在**config.json**里指定的）。

![Screen Shot 2021-06-01 at 16 29 33](https://user-images.githubusercontent.com/16622934/120292283-c3f09780-c2f6-11eb-8fdc-9fac4fed6e9f.png)

8. 然后就像正常本地视频一样播放即可！

![Screen Shot 2021-06-01 at 16 30 20](https://user-images.githubusercontent.com/16622934/120292325-cce16900-c2f6-11eb-89ed-370c007024aa.png)
![Screen Shot 2021-06-01 at 16 30 33](https://user-images.githubusercontent.com/16622934/120292339-cfdc5980-c2f6-11eb-8164-cf1380580128.png)

## TODO

- [ ] 自动获取下载链接和cookie。