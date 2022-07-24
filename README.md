本周作业
编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

1. 接收客户端 request，并将 request 中带的 header 写入 response header
1. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
1. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
1. 当访问 `localhost/healthz` 时，应返回 200