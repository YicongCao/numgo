#-*- coding:utf-8 -*-

import pytest
import requests
import json

# 基于pytest requests测试 接口
class TestRequestDemo:
    # 初始化
    url = "http://localhost:8080"
    session = requests.session()    

    # 通过 get 请求测试计算逻辑
    @pytest.mark.parametrize("a, b, expected",
    [(1,"echo","1"), (9,"fib","55"), (16,"sqrt","4")])  
    def test_get_num(self, a, b, expected):
        r = self.session.get(self.url + "/test?a={0}&b={1}".format(a, b))        
        # 断言状态码
        assert r.status_code == 200
        
        # 断言响应头信息
        assert r.headers["Content-Type"] == "application/json"

        # 断言计算结果
        data = r.json()        
        assert data["reference"] == expected
    
    # 通过 post 请求测试计算逻辑
    @pytest.mark.parametrize("a, b, expected",
    [(1,"echo","1"), (5,"fib","8"), (9,"sqrt","3")])  
    def test_post_num(self, a, b, expected):
        mimetype = 'application/json'
        headers = {
            'Content-Type': mimetype,
            'Accept': mimetype
        }
        data = {
            "a": a,
            "b": b,
        }
        r = self.session.post(self.url + "/test", data=json.dumps(data), headers=headers)        
        
        # 断言状态码
        assert r.status_code == 200
        
        # 断言响应头信息
        assert r.headers["Content-Type"] == "application/json"

        # 断言计算结果
        data = r.json()        
        assert data["reference"] == expected