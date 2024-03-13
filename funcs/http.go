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
	defer response.Body.Close()

	if runnable {
		content, err = io.ReadAll(response.Body)
		if err != nil {
			runnable = false
		}
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
	defer response.Body.Close()

	if runnable {
		code = response.StatusCode
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
	defer response.Body.Close()

	if runnable {
		content, err = io.ReadAll(response.Body)
		if err != nil {
			runnable = false
		}
	}

	return string(content), err
}
