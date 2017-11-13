# hello --- Go Version

## 工具
### negroni+mux

## 理由
### 因为功能简单，所以选择轻量组件 mux+negroni

## 测试结果
```
curl -v hhtp://localhost:9090/hello/nick
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 9090 (#0)
> GET /hello/nick HTTP/1.1
> Host: localhost:9090
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Date: Mon, 13 Nov 2017 14:15:26 GMT
< Content-Length: 23
< Content-Type: text/plain; charset=utf-8
< 
{"Test":"Hello nick"}
* Connection #0 to host localhost left intact
```

## log记录
```
[negroni] 2017-11-13T22:15:26+08:00 | 200 |      220.183µs | localhost:9090 | GET /hello/nick
```

## 压力测试

### 测试命令
```
ab -n 1000 -c 100 http://localhost:9090/hello/nick
```
* -n 请求数量
* -c 并发数量

### 测试结果
```
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            9090

Document Path:          /hello/nick
Document Length:        23 bytes

Concurrency Level:      100
Time taken for tests:   0.147 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      140000 bytes
HTML transferred:       23000 bytes
Requests per second:    6783.43 [#/sec] (mean)
Time per request:       14.742 [ms] (mean)
Time per request:       0.147 [ms] (mean, across all concurrent requests)
Transfer rate:          927.42 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   1.2      0       5
Processing:     0   13   7.7     11      38
Waiting:        0   12   7.9     11      38
Total:          0   13   7.5     12      40

Percentage of the requests served within a certain time (ms)
  50%     12
  66%     16
  75%     17
  80%     18
  90%     23
  95%     29
  98%     38
  99%     38
 100%     40 (longest request)
```
* **Requests per second** --- 吞吐量指的是某个并发用户数下单位时间内处理的请求数。
* **Time per request** --- 用户平均请求等待时间
* **Time per request:across all concurrent requests** --- 服务器平均请求等待时间