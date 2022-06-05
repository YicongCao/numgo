# numgo

简易 http 小程序，用来跑自动化测试框架。

支持 GET/POST 方法，传参如下：

|field|type|desc|
|-|-|-|
|a|int|任意正整数|
|b|string|echo/fib/sqrt|

可以计算斐波那契第a位，或者向下取整的a的开方。

给出响应如下：

|field|type|desc|
|-|-|-|
|error_code|int|错误码|
|error_message|string|错误信息|
|reference|string|计算结果|

测试：

```bash
go run main.go # 启动服务
curl "http://localhost:8080/test?a=9&b=fibx"
```