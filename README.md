# numgo

简易 http 小程序，用来跑自动化测试框架。

支持 GET/POST 方法，传参如下：

|field|type|desc|
|-|-|-|
|a|int|任意正整数|
|b|string|echo/fib/sqrt|

测试：

```bash
curl "http://localhost:8080/test?a=9&b=fibx"
```