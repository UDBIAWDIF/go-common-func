package funcs

import (
	"io"
	"net/http"
	"strings"
	"time"
)

const HTTP_TIMEOUT_DEFAULT time.Duration = time.Second * 10
const HTTP_CONTENT_TYPE_JSON string = "application/json"

func HttpGet(url string) (string, error) {
	runnable := true
	var content []byte
	var err error
	client := http.Client{Timeout: HTTP_TIMEOUT_DEFAULT}

	var response *http.Response
	response, err = client.Get(url)
	if err != nil {
		runnable = false
	}

	if runnable {
		content, err = io.ReadAll(response.Body)
		if err != nil {
			runnable = false
		}
	}

	if response != nil {
		response.Body.Close()
	}

	return string(content), err
}

func HttpGetHttpCode(url string) (int, error) {
	runnable := true
	var code int
	var err error
	client := http.Client{Timeout: HTTP_TIMEOUT_DEFAULT}

	var response *http.Response
	response, err = client.Get(url)
	if err != nil {
		runnable = false
	}

	if runnable {
		code = response.StatusCode
	}

	if response != nil {
		response.Body.Close()
	}

	return code, err
}

func HttpPost(url string, contentType string, body []byte) (string, error) {
	runnable := true
	var content []byte
	var err error
	client := http.Client{Timeout: HTTP_TIMEOUT_DEFAULT}

	var response *http.Response
	response, err = client.Post(url, contentType, strings.NewReader(string(body)))
	if err != nil {
		runnable = false
	}

	if runnable {
		content, err = io.ReadAll(response.Body)
		if err != nil {
			runnable = false
		}
	}

	if response != nil {
		response.Body.Close()
	}

	return string(content), err
}

type HttpRequestHandler func(*http.Request)

// 发启HTTP请求, 可以使用请求处理器对请求进行一些初始化操作, 比如设置头部信息等
func HttpPostUseHandle(url string, body []byte, httpRequestHandler HttpRequestHandler) (string, error) {
	runnable := true
	var content []byte
	var err error
	client := http.Client{Timeout: HTTP_TIMEOUT_DEFAULT}
	var response *http.Response

	req, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		runnable = false
	}

	if runnable {
		if httpRequestHandler != nil {
			httpRequestHandler(req)
		}

		response, err = client.Do(req)
		if err != nil {
			runnable = false
		}
	}

	if runnable {
		content, err = io.ReadAll(response.Body)
		if err != nil {
			runnable = false
		}
	}

	if response != nil {
		response.Body.Close()
	}

	return string(content), err
}
