package funcs

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) (string, error) {
	runnable := true
	var body []byte
	var err error

	var response *http.Response
	response, err = http.Get(url)
	if err != nil {
		runnable = false
	}
	defer response.Body.Close()

	if runnable {
		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			runnable = false
		}
	}

	return string(body), err
}

func HttpGetHttpCode(url string) (int, error) {
	runnable := true
	var code int
	var err error

	var response *http.Response
	response, err = http.Get(url)
	if err != nil {
		runnable = false
	}
	defer response.Body.Close()

	if runnable {
		code = response.StatusCode
	}

	return code, err
}
