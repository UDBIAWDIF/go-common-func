package httpclient

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const HTTP_TIMEOUT_DEFAULT time.Duration = time.Second * 10
const HTTP_CONTENT_TYPE_JSON string = "application/json"

type HttpClient struct {
	Proxy          string // 代理地址, 如果有配置, 使用此地址做代理
	NicIp          string // 网卡IP, 如果有配置, 便宜此网卡做出口
	Timeout        time.Duration
	Method         string         // 访问方式, 默认GET
	RequestHandler RequestHandler // 请求处理器, 可以对请求进行一些初始化操作, 比如设置头部信息等
}

type RequestHandler func(*http.Request)

func NewHttpClient() *HttpClient {
	return &HttpClient{
		Method:  http.MethodGet,
		Timeout: HTTP_TIMEOUT_DEFAULT,
	}
}

func (_t HttpClient) Request(
	targetUrl string,
	contentType string,
	body []byte) (responseBody string, httpCode int, err error) {

	runnable := true
	var content []byte
	client := http.Client{
		Timeout:   HTTP_TIMEOUT_DEFAULT,
		Transport: _t.GetTransport()}

	var response *http.Response
	var req *http.Request

	req, err = http.NewRequest(_t.Method, targetUrl, strings.NewReader(string(body)))
	if err != nil {
		runnable = false
	}

	if runnable {
		if _t.RequestHandler != nil {
			_t.RequestHandler(req)
		}

		if contentType != "" {
			req.Header.Set("Content-Type", contentType)
		}

		response, err = client.Do(req)
		if err != nil {
			runnable = false
		}
	}

	if runnable {
		httpCode = response.StatusCode
		content, err = io.ReadAll(response.Body)
		if err != nil {
			runnable = false
		}
	}

	if response != nil {
		response.Body.Close()
	}

	responseBody = string(content)
	return
}

func (_t *HttpClient) Get(targetUrl string) (responseBody string, err error) {
	_t.Method = http.MethodGet
	responseBody, _, err = _t.Request(targetUrl, "", nil)
	return
}

func (_t *HttpClient) GetWithParams(targetUrl string, params url.Values) (string, error) {
	urlRequest, _ := url.ParseRequestURI(targetUrl)
	urlRequest.RawQuery = params.Encode()
	return _t.Get(urlRequest.String())
}

func (_t *HttpClient) Post(targetUrl string, contentType string, body []byte) (responseBody string, err error) {
	_t.Method = http.MethodPost
	responseBody, _, err = _t.Request(targetUrl, "", nil)
	return
}

func (_t *HttpClient) GetHttpCode(targetUrl string) (httpCode int, err error) {
	_t.Method = http.MethodGet
	_, httpCode, err = _t.Request(targetUrl, "", nil)
	return
}

// 设置代理地址
func (_t *HttpClient) SetProxy(proxy string) *HttpClient {
	_t.Proxy = proxy
	return _t
}

// 设置请求超时
func (_t *HttpClient) SetTimeout(timeout time.Duration) *HttpClient {
	_t.Timeout = timeout
	return _t
}

// 设置出口网卡IP
func (_t *HttpClient) SetNicIp(nicIp string) *HttpClient {
	_t.NicIp = nicIp
	return _t
}

// 设置请求处理器, 可以对请求进行一些初始化操作, 比如设置头部信息等
func (_t *HttpClient) SetRequestHandler(handler RequestHandler) *HttpClient {
	_t.RequestHandler = handler
	return _t
}

func (_t HttpClient) GetTransport() *http.Transport {
	/* var localAddr net.Addr = nil
	if _t.NicIp != "" {
		localAddr = &net.TCPAddr{IP: net.ParseIP(_t.NicIp)}
		// localAddr, _ = net.ResolveTCPAddr("tcp", _t.NicIp+":0")
	} */

	transport := &http.Transport{}
	if _t.Proxy != "" {
		transport = &http.Transport{
			/* Dial: (&net.Dialer{
				Timeout:   _t.Timeout,
				KeepAlive: _t.Timeout,
				LocalAddr: localAddr,
			}).Dial, */
			Dial: func(netType, addr string) (net.Conn, error) {
				// 网卡IP
				localAddr, err := net.ResolveTCPAddr("tcp", _t.NicIp+":0") // ":0" 表示端口自动选择
				if err != nil {
					return nil, err
				}

				remoteAddr, err := net.ResolveTCPAddr(netType, addr)
				if err != nil {
					return nil, err
				}

				conn, err := net.DialTCP(netType, localAddr, remoteAddr)
				if err != nil {
					return nil, err
				}

				return conn, nil
			},
			TLSHandshakeTimeout:   _t.Timeout,
			ResponseHeaderTimeout: _t.Timeout,
			ExpectContinueTimeout: _t.Timeout,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse(_t.Proxy)
			},
		}
	}

	return transport
}
