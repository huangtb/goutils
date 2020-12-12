package web

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	// method
	MethodGet    = "GET"
	MethodPost   = "POST"
	MethodPut    = "PUT"
	MethodDelete = "DELETE"
	// send type
	SendTypeFrom = "from"
	SendTypeJson = "json"
)


type HttpSend struct {
	Link     string
	SendType string
	Header   map[string]string
	Body     map[string]interface{}
	Timeout  time.Duration
	sync.RWMutex
}

func NewHttpSend(link string) *HttpSend {
	return &HttpSend{
		Link: link,
	}
}

func (h *HttpSend) SetTimeOut(seconds time.Duration) {
	h.Lock()
	defer h.Unlock()
	h.Timeout = time.Second * seconds
}

func (h *HttpSend) SetBody(body map[string]interface{}) {
	h.Lock()
	defer h.Unlock()
	h.Body = body
}

func (h *HttpSend) SetHeader(header map[string]string) {
	h.Lock()
	defer h.Unlock()
	h.Header = header
}

func (h *HttpSend) SetSendType(sendType string) {
	h.Lock()
	defer h.Unlock()
	h.SendType = sendType
}

func (h *HttpSend) Get() ([]byte, error) {
	return h.send(MethodGet)
}

func (h *HttpSend) Post() ([]byte, error) {
	return h.send(MethodPost)
}

func (h *HttpSend) Put() ([]byte, error) {
	return h.send(MethodPut)
}

func (h *HttpSend) Delete() ([]byte, error) {
	return h.send(MethodDelete)
}


//reqParams：请求结构体
func (h *HttpSend) HttpPost(reqParams interface{}) ([]byte ,error) {
	var (
		req      *http.Request
		resp     *http.Response
		client   http.Client
		err      error
	)

	b, err := json.Marshal(reqParams)
	if err != nil {
		fmt.Errorf("request param marshal error:%v", err.Error())
		return nil,err
	}
	if h.Timeout <= 0 {
		h.Timeout = 60 * time.Second
	}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client.Timeout = h.Timeout
	req, err = http.NewRequest(http.MethodPost, h.Link, strings.NewReader(string(b)))
	if err != nil {
		fmt.Errorf("new request error:%s", err.Error())
		return nil,err
	}
	req.Header.Set("Content-Type", "application/json; encoding=utf-8")
	if len(h.Header) != 0 {
		for key, value := range h.Header {
			req.Header.Add(key, value)
		}
	}
	//执行请求
	//client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Errorf("do http post request error:%s", err.Error())
		return nil,err
	}
	defer resp.Body.Close()
	//读取响应
	body, err := ioutil.ReadAll(resp.Body) //此处可增加输入过滤
	if err != nil {
		fmt.Errorf("read resp body error:%v", err.Error())
		return nil,err
	}
	fmt.Printf(" Http success ,response: %s", string(body))

	return body,nil
}

