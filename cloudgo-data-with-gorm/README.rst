cloudgo-data-with-gorm
============

NOTE
------------
这个 repo 本应在 `sysu-service-computing`_ 里面

.. _`sysu-service-computing`: https://github.com/Binly42/sysu-service-computing/tree/master/cloudgo-data-with-gorm


Resource
------------
`老师的博客`_

.. _`老师的博客`: http://blog.csdn.net/pmlpml/article/details/78602290


Overview
------------
是依照老师的 cloudgo-data_, 基于 ORM (这里用的是 *gorm*) (而非 手写 DAO) 来修改的 ;

.. _cloudgo-data: https://github.com/pmlpml/golang-learning/tree/master/web/cloudgo-data


说明
------------

从 编程效率、程序结构、服务性能 等角度对比 database/sql 与 orm 实现的异同:

    * for 编程效率, 我认为, (起码在 规模/复杂度 相对较小 的情景下) 基于 ORM 的开发效率会更高, 毕竟抽象不是白做的 ..... ;

    * for 程序结构, 不一定 ; 不过, ORM 的模式下的程序结构并不一定会比 (写的好的) DAO 好, 毕竟一般不是专门定制的 ... ;

    * for 服务性能, 显然, ORM 的实现模式往往会带来性能方面的问题 ;

    #. orm 是否就是实现了 dao 的自动化？

        #. 不只是 ;

    #. 使用 ab 测试性能:

        Using command `ab -n 1000 -c 100 localhost:8080/service/userinfo?userid=`

        ::

            binly@binly-virtual-machine:~$
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
            Server Port:            8080

            Document Path:          /service/userinfo?userid=
            Document Length:        890 bytes

            Concurrency Level:      100
            Time taken for tests:   1.848 seconds
            Complete requests:      1000
            Failed requests:        0
            Total transferred:      1014000 bytes
            HTML transferred:       890000 bytes
            Requests per second:    541.21 [#/sec] (mean)
            Time per request:       184.771 [ms] (mean)
            Time per request:       1.848 [ms] (mean, across all concurrent requests)
            Transfer rate:          535.92 [Kbytes/sec] received

            Connection Times (ms)
                        min  mean[+/-sd] median   max
            Connect:        0    3   7.3      0      39
            Processing:     2  176 252.6     95     966
            Waiting:        2  174 252.9     91     944
            Total:          2  180 252.1     99     968

            Percentage of the requests served within a certain time (ms)
            50%     99
            66%    115
            75%    138
            80%    156
            90%    695
            95%    939
            98%    943
            99%    950
            100%    968 (longest request)
