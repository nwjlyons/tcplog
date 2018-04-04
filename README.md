# tcplog

Download binarys from releases page https://github.com/nwjlyons/tcplog/releases

Server

```
tcplog -port 8080
```

Client

```
curl localhost:8080
```

Example output:

```
2018/04/04 11:28:43 listening to tcp connections on localhost:8100
GET / HTTP/1.1
Host: localhost:8100
User-Agent: curl/7.54.0
Accept: */*
```

Another example:

```
http -v -f POST localhost:8101 X-API-Token:123 foo=bar
```

```
Host: localhost:8101
User-Agent: HTTPie/0.9.9
Accept-Encoding: gzip, deflate
Accept: */*
Connection: keep-alive
Content-Type: application/x-www-form-urlencoded; charset=utf-8
X-API-Token: 123
Content-Length: 7

foo=bar
```
