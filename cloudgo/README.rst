cloudgo
============

NOTE
------------
这个 repo 本应在 `sysu-service-computing`_ 里面

.. _`sysu-service-computing`: https://github.com/Binly42/sysu-service-computing/tree/master/cloudgo


Resource
------------
`BLABLA`_

.. _`BLABLA`: http://blog.csdn.net/pmlpml/article/details/78404838


Overview
------------
是依照老师的 cloudgo_ 来写的, 没有用框架或是第三方包, 因为功能实在非常简单... 所以只是简单包装了下 net/http 包里的东西即可 ;

.. _cloudgo: https://github.com/pmlpml/golang-learning/blob/master/web/cloudgo


Usage
------------
在该 cloudgo/ 目录下运行:

`go build`

`./cloudgo`

即可在 (默认的) 8080端口 上运行 ;


Sample
------------

    * 运行 `curl -v localhost:8080/hello/me` 输出为:

        ::

            * timeout on name lookup is not supported
            *   Trying ::1...
            * TCP_NODELAY set
            % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                            Dload  Upload   Total   Spent    Left  Speed
            0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0* Connected to localhost (::1) port 8080 (#0)
            > GET /hello/me HTTP/1.1
            > Host: localhost:8080
            > User-Agent: curl/7.53.0
            > Accept: */*
            >
            < HTTP/1.1 200 OK
            < Date: Tue, 14 Nov 2017 14:18:18 GMT
            < Content-Length: 10
            < Content-Type: text/plain; charset=utf-8
            <
            { [10 bytes data]
            100    10  100    10    0     0    312      0 --:--:-- --:--:-- --:--:--   625Hello me!

            * Connection #0 to host localhost left intact

    * 运行 `./ab.exe -n 1000 -c 100 localhost:8080/hello/me` (Apache 的 ab测试) 输出为:

        ::

            This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
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


            Server Software:
            Server Hostname:        localhost
            Server Port:            8080

            Document Path:          /hello/me
            Document Length:        10 bytes

            Concurrency Level:      100
            Time taken for tests:   0.943 seconds
            Complete requests:      1000
            Failed requests:        0
            Total transferred:      127000 bytes
            HTML transferred:       10000 bytes
            Requests per second:    1060.31 [#/sec] (mean)
            Time per request:       94.312 [ms] (mean)
            Time per request:       0.943 [ms] (mean, across all concurrent requests)
            Transfer rate:          131.50 [Kbytes/sec] received

            Connection Times (ms)
                            min  mean[+/-sd] median   max
            Connect:        0    0   0.3      1       4
            Processing:     4   89  52.2     75     266
            Waiting:        2   85  54.0     73     243
            Total:          5   90  52.2     75     267
            ERROR: The median and mean for the initial connection time are more than twice the standard
                    deviation apart. These results are NOT reliable.

            Percentage of the requests served within a certain time (ms)
                50%     75
                66%     88
                75%    106
                80%    108
                90%    184
                95%    239
                98%    242
                99%    243
                100%    267 (longest request)
            Finished 1000 requests

        其中, 参数 `-n 1000` 表示一共测试 1000个请求, 而 `-c 100` 表示并发的请求数目 ;
