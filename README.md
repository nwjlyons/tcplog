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
