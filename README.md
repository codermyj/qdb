# qdb
### qdb-1.0.0 接口说明：
- 1. 接口程序位于src/bin目录下：
    - 1. qdb.go：本地命令行访问
    - 2. server.go： 运行go run server.go ip:port，启动一个服务监听ip的port端口
    - 3. client.go： 运行go run client.go ip:port，启动一个客户端命令行，访问server所在的ip地址的port
- 2. 命令行接口说明
    - 1. set key value，设置key对应的value，并存储到磁盘
    - 2. get key，从存储的值中获取key对应的value值
    - 3. remove key，删除key








