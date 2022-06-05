package numgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	retSuccess     = 0
	retSystemError = 1
	retParamError  = 2
)

type numreq struct {
	A int    `json:"a"`
	B string `json:"b"`
}

type numrsp struct {
	ErrCode int    `json:"error_code"`
	ErrMsg  string `json:"error_message"`
	Ref     string `json:"reference"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	req, rsp := &numreq{}, &numrsp{}
	// 封装返回
	defer func() {
		rspbytes, _ := json.MarshalIndent(rsp, "", "  ")
		w.Write(rspbytes)
	}()
	// 解析请求
	switch r.Method {
	case "GET":
		u, err := url.Parse(r.URL.Path)
		if err != nil {
			rsp.ErrCode = retSystemError
			rsp.ErrMsg = fmt.Sprintf("parse req failed, err: %v", err)
		}
		m, _ := url.ParseQuery(u.RawQuery)
		if ma, ok := m["a"]; ok {
			req.A, _ = strconv.Atoi(ma[0])
		}
		if mb, ok := m["b"]; ok {
			req.B = mb[0]
		}
	case "POST":
		d := json.NewDecoder(r.Body)
		if err := d.Decode(req); err != nil {
			rsp.ErrCode = retSystemError
			rsp.ErrMsg = fmt.Sprintf("parse req failed, err: %v", err)
		}
	default:
		rsp.ErrCode = retParamError
		rsp.ErrMsg = "method not allowed"
	}
	// 计算逻辑
	switch req.B {
	case "echo":
		rsp.Ref = fmt.Sprintf("%d", Echo(req.A))
	case "fib":
		rsp.Ref = fmt.Sprintf("%d", Fib(req.A))
	case "sqrt":
		rsp.Ref = fmt.Sprintf("%d", Sqrt(req.A))
	default:
		rsp.ErrCode = retParamError
		rsp.ErrMsg = fmt.Sprintf("unknown cmd: %s", req.B)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
