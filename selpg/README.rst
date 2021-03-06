selpg
============


resource
------------
`开发 Linux 命令行实用程序`_

.. _`开发 Linux 命令行实用程序`: https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html


设计说明
------------
直接按相关网页(IBM.com)上的意思, 基本上是按部就班地实现:

    * 代码没有作什么结构上的调整, 所以基本就靠几个函数的分工组织起来 ;

    * 几个主要函数的功能就如其名, 直接放在顶层, 而次要的便用 `var` 作定义 ;

    * 直接用大量的全局变量等来通讯及记录状态, 包括跟命令行参数有关的变量 ;

    * 用 flag 处理命令含参数解析; 为了尽量靠近网页中对参数的要求, 便在 `flag.Parse()` 之后再作一些检验工作 ; 加之, 能借助 shell 自身的功能就让 shell 接管, 所以对于 pipe 和 重定向操作 所需适配不多 ;

    * 由读取输入的函数返回 `channel` 按行传输 string 类型的输入内容, 因为全程不需要对内容本身作改动, 且 Go 的 `for range` 语句在遍历 string 时会自动按 unicode 形式读取每个单位字符, 恰好可免去相关的转换操作; 而至于读取方式, 则采用 bufio 中的相关接口, 写入也是一样;

    * 根据命令行参数而具体决定要进行 "打印" 的页数范围(我以 1 作为起始页号) 以及分页方式 ; 这里没顾得上将选择 "页" 进行打印的过程抽象为 filter ;

    * 最后, `-d` 参数的具体实现效果不清楚, 因为没有打印机设备供调试 ... 至于实现方式, 则是也是类似着网页文档里说的, 利用 os/exec 包启动 `lp` 命令 ;


使用
------------
使用方式与 selpg 相同;

例:
    
    `go build main.go`
    
    `./main.exe -s 1 -e 2 -l 3 input_flie`


测试结果
------------
不考虑 `-d` 参数的话, 各测试样例均合乎预期;
