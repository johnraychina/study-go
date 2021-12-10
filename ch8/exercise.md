练习 8.2： 实现一个并发FTP服务器。
服务器应该解析客户端发来的一些命令，比如cd命令来切换目录，
ls来列出目录内文件，get和send来传输文件，close来关闭连接。
你可以用标准的ftp命令来作为客户端，或者也可以自己实现一个。

练习 8.10： HTTP请求可能会因http.Request结构体中Cancel channel的关闭而取消。
修改8.6节中的web crawler来支持取消http请求。
（提示：http.Get并没有提供方便地定制一个请求的方法。
你可以用http.NewRequest来取而代之，设置它的Cancel字段，然后用http.DefaultClient.Do(req)来进行这个http请求。）

练习 8.11： 紧接着8.4.4中的mirroredQuery流程，实现一个并发请求url的fetch的变种。
当第一个请求返回时，直接取消其它的请求。