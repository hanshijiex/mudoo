# mudoo
Go实现的网络通信库

## 楔子
受陈硕的muduo网络库影响，立志也要在Go语言领域里打造一个如此的利器。

## 我心目中理想的网络库的样子
* 协程安全，原生支持多核多协程.
* 不考虑可移植性，不跨平台，只支持Linux，不支持Windows.
* 主要支持x86-64，兼顾IA32.
* 不支持UDP，只支持TCP.
* 不支持IPv6，只支持IPv4.
* API简单易用，只暴露具体类和标准库里的类。
* 只满足常用需求的90%，不面面俱到，必要的时候以app来适应lib.
* 只做library，不做成framework.
* 争取全部代码在1000行以内（不含测试）.
* 在不增加复杂库的前提下可以支持FreeBSD/Darwin, 方便将来用Mac作为开发用机，但不为它做性能优化。也就是说，IO multiplexing使用poll(2)和epoll(4).

## 欢迎加入
Gopher厦门QQ群：480356472
