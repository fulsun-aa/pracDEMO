这里想要弄清楚http

如果要创建Web服务器端，则需要
    1.http.HandleFunc(路径，处理函数)，HandleFunc用于为指定的URL注册一个处理器，它会调用defaultServeMux的handdlefanc()，也就是说
        HandleFunc会把处理器注册到多路复用当中。
    2.http.ListenAndServe(url,多路复用器)，当多路复用器为nil的时候，会默认使用DefaultServeMux
当接收到客户端的请求，多路复用器（serveMus）根据url映射到一个相应的处理器。回去找到一个能处理请求的处理器，根据url，找到一个对应的处理器的serveHTTP()来处理请求
handler是一个接口，你可以写一个结构体，让它只要实现serveHTTP(ResponseWriter,*Request)即可

