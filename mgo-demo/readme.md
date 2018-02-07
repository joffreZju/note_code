#### 实验1
- 数据：一个collection，130万条记录，将其中的18万数据分成24*60组，分别将每组数据查询取出并修改然后update。
- 串行版本：27m56.9226705s
- 并行版本：1m49.959916s
- 存在问题：不限制并发数的话，由于mongo connection太多服务器会自动关闭，因此需要限制并发数，目前为100,同时也要限制mongo连接数（二者至少要限制一个）。
- 报错信息
```
服务器
2017-07-21T14:25:13.893+0000 I SHARDING [conn35960] Exception thrown while processing query op for hzds.$cmd :: caused by :: 9001 socket exception [SEND_ERROR] server [10.214.224.126:57696] 
2017-07-21T14:25:13.893+0000 I NETWORK  [conn35960] SocketException handling request, closing client connection: 9001 socket exception [SEND_ERROR] server [10.214.224.126:57696] 

客户端
2017/07/22 12:59:33 Find : ERROR : read tcp 10.214.224.126:50985->10.214.224.142:20000: wsarecv: An existing connection was forcibly closed by the remote host.
2017/07/22 12:59:34 Find : ERROR : End of file

```

#### 实验2
- 其他同上，但是只取24*60中的一组数据，查询并更新。
- 串行：56.113806s
- 并行：2.91993s

#### 实验3
- 并发测试,完全不能处理
```
./ab.exe -c 100 -n 1000 http://localhost:9002/fast

```



