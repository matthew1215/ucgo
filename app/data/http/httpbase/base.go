package httpbase

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type BaseRequest struct {
	Url        string
	Method     string
	Header     map[string]string
	GetParams  map[string]string
	PostParams map[string]string
	Timeout    time.Duration
}

//type TimeoutConn struct {
//	net.Conn
//	timeout time.Duration
//}

var DefaultBaseRequest = BaseRequest{
	Timeout: time.Second * 2,
	Method:  "GET",
}

//func NewTimeoutConn(conn net.Conn, timeout time.Duration) *TimeoutConn {
//	return &TimeoutConn{conn,timeout}
//}
//
//func (c *TimeoutConn) Read(b []byte) (n int, err error) {
//	c.SetReadDeadline(time.Now().Add(c.timeout))
//	return c.Conn.Read(b)
//}
//
//func (c *TimeoutConn) Write(b []byte) (n int, err error) {
//	c.SetWriteDeadline(time.Now().Add(c.timeout))
//	return c.Conn.Write(b)
//}

func (this BaseRequest) Do() (string, error) {
	body := strings.NewReader("")
	if this.Method == "" && len(this.PostParams) != 0 {
		this.Method = "POST"
	} else if this.Method == "" {
		this.Method = DefaultBaseRequest.Method
	}
	if this.Timeout == 0 {
		this.Timeout = DefaultBaseRequest.Timeout
	}
	if this.Method == "POST" {
		values := url.Values{}
		for k, v := range this.PostParams {
			values.Add(k, v)
		}
		postParams := values.Encode()
		body = strings.NewReader(postParams)
	}
	if this.Method == "JSON" {
		this.Method = "POST"
		postParams, _ := json.Marshal(this.PostParams)
		body = strings.NewReader(string(postParams))
	}
	req, _ := http.NewRequest(this.Method, this.Url, body)
	req.Header.Set("Content-type", "application/x-www-form-urlencoded;charset=utf-8")
	for k, v := range this.Header {
		req.Header.Set(k, v)
	}
	if this.Method == "JSON" {
		req.Header.Set("Content-type", "application/json;charset=utf-8")
	}
	values := req.URL.Query()
	for k, v := range this.GetParams {
		values.Add(k, v)
	}
	req.URL.RawQuery = values.Encode()
	return this.getResponse(req)
}

func (this BaseRequest) getResponse(req *http.Request) (string, error) {
	client := this.getClient()
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	//程序在使用完回复后必须关闭回复的主体。
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	resStr := string(body)
	return resStr, err
}

func (this BaseRequest) getClient() *http.Client {
	//client := &http.Client{
	//	Transport: &http.Transport{
	//		Dial: func(netw, addr string) (net.Conn, error) {
	//			conn, err := net.DialTimeout(netw, addr, time.Second*2)
	//			if err != nil {
	//				return nil, err
	//			}
	//
	//			return NewTimeoutConn(conn, time.Second*2), nil
	//		},
	//		ResponseHeaderTimeout: time.Second * 2,
	//	},
	//}
	client := &http.Client{
		Timeout: time.Second * this.Timeout,
	}
	return client
}
