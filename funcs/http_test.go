package funcs

import (
	"net/http"
	"testing"
)

func HttpPostAddHeader(req *http.Request) {
	req.Header.Set("Content-Type", HTTP_CONTENT_TYPE_JSON)
	req.Header.Set("Power-By", "郑洪耕")
}

// go test -v -run="TestHttpPost"
func TestHttpPost(t *testing.T) {
	response, _ := HttpPostUseHandle("http://127.0.0.1:8891/api/human/zoobll", []byte("UID is the best!"), HttpPostAddHeader)
	t.Log(response)
}
