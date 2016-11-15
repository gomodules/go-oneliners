package oneliners

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func DumpHttpRequest(req *http.Request) {
	b, err := httputil.DumpRequest(req, true)
	if err == nil {
		fmt.Println(string(b))
	}
}

func DumpHttpRequestOut(req *http.Request) {
	b, err := httputil.DumpRequestOut(req, true)
	if err == nil {
		fmt.Println(string(b))
	}
}

func DumpHttpResponse(resp *http.Response) {
	b, err := httputil.DumpResponse(resp, true)
	if err == nil {
		fmt.Println(string(b))
	}
}
