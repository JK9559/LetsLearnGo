## ServeMux 自定义
1. 锁
2. 路由规则 是map 类型为 map[string]muxEntry
3. 是否在任意规则中带有host信息

### muxEntry 结构
1. 是否为精确匹配
2. 本路由表达式对应了哪个handler
3. 匹配字符串

## Handler定义
1. 是一个接口类型
2. 实现了ServeHTTP方法
3. 3_2的代码sayhelloName并未实现该接口，其中该函数作为变量传入的函数http.HandleFunc，强制转为HandlerFunc类型
4. 进而实现了ServeHTTP方法

### 针对http包 主要流程如下
* 首先调用http.HandleFunc("/", sayHelloName)，期间按顺序做了以下几件事
	1. 调用DefaultServeMux.HandleFunc(pattern, handler)
	2. HandleFunc _registers_ the handler function for the given pattern.
* 调用http.ListenAndServe(":9090", nil)
	1. 实例化Server
	2. 调用Server的ListenAndServe()
	3. 调用net.Listen("tcp", addr)监听端口
	4. 启动一个for循环 在循环体里接收请求
	5. 对于每个请求实例化一个Conn，并且开启一个goroutine为这个请求进行服务go c.serve()
	6. 读取每个请求的内容w, err := c.readRequest(ctx)
	7. 判断handler是否为空 没有则设置为DefaultServerMux
	8. 调用handler中的ServerHttp
	9. 根据req选择handler，进入到ServerHTTP
	10. 选择路由，判断是否有路由满足，遍历ServerMux的muxEntry
	11. 有满足的调 没满足的NotFoundHandler的ServeHttp