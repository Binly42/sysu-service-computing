cloudgo-io
============

NOTE
------------
这个 repo 本应在 `sysu-service-computing`_ 里面

.. _`sysu-service-computing`: https://github.com/Binly42/sysu-service-computing/tree/master/cloudgo-io


Resource
------------
`BLABLA`_

.. _`BLABLA`: http://blog.csdn.net/pmlpml/article/details/78539261


Overview
------------
是依照老师的 cloudgo-XXX_ 来写的, 没有用框架或是第三方包, 因为功能实在非常简单... 所以只是简单包装了下 net/http 包里的东西即可 ;

.. _cloudgo-XXX: https://github.com/pmlpml/golang-learning/tree/master/web

.. _基本要求:

参照基本要求:

    #. 支持静态文件服务

    #. 支持简单 js 访问

    #. 提交表单，并输出一个表格

    #. 对 /unknown 给出开发中的提示，返回码 5xx



Usage
------------
在该 cloudgo/ 目录下运行:

`go build`

`./cloudgo-io`

即可在 (默认的) 8080端口 上运行 ;


说明
------------

参照 基本要求_:

    #. 间接调用 net/http 包里自带的 FileServer 及相关接口, 直接直接整个 *asset/* 目录 ; (默认可用 url: `localhost:8080/static/` 访问)

    #. 会以跟老师的类似的方式响应 `api/test` 请求 ; (默认可用 url: `localhost:8080/api/test` 访问)

    #. 提供了一个 post-phone-info.html 的静态页面 (现作为首页, 可直接通过 `/` 访问), 里面允许提交表单(form), 然后服务器端会随便解析处理一下, 最后返回一个含表格的 (借助 html/template 包里的 模板接口 生成的) 静态页面 ; (默认可用 url: `localhost:8080/` 访问)

    #. 直接向响应头中写入 http.StatusNotImplemented ; (默认可用 url: `localhost:8080/unknown` 访问)

