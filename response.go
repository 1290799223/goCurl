package goCurl

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

// Response response object
type Response struct {
	resp       *http.Response
	req        *http.Request
	cookiesJar *cookiejar.Jar
	err        error
}

// GetCookies, 获取服务端生成的全部cookies
func (r *Response) GetCookies() []*http.Cookie {
	return r.cookiesJar.Cookies(r.req.URL)
}

// GetCookie, 通过键获取相关的cookie值
func (r *Response) GetCookie(cookieName string) *http.Cookie {
	cookies := r.cookiesJar.Cookies(r.req.URL)
	if len(cookies) > 0 {
		for i := 0; i < len(cookies); i++ {
			if cookies[i].Name == cookieName {
				return cookies[i]
			}
		}
	}
	return nil
}

// GetRequest get request object
func (r *Response) GetRequest() *http.Request {
	return r.req
}

// GetRequest get request object
func (r *Response) GetResponse() *http.Response {
	return r.resp
}

// GetBody parse response body
func (r *Response) GetContents() (string, error) {
	defer func() {
		_ = r.resp.Body.Close()
	}()
	temp := strings.ReplaceAll(fmt.Sprintf("%v", r.resp.Header["Content-Type"]), " ", "")
	var bodystr string
	body, err := ioutil.ReadAll(r.resp.Body)
	if err != nil {
		return "", err
	}
	if strings.Contains(strings.ToLower(temp), "charset=utf") {
		bodystr = string(body)
	} else {
		bodystr = simpleChinese2Utf8(body)
	}

	return bodystr, nil
}

// Get Response ContentLength
func (r *Response) GetContentLength() int64 {
	return r.resp.ContentLength
}

// GetBody parse response body
func (r *Response) GetBody() io.ReadCloser {
	//defer r.resp.Body.Close()

	return r.resp.Body
}

// GetStatusCode get response status code
func (r *Response) GetStatusCode() int {
	return r.resp.StatusCode
}

// GetReasonPhrase get response reason phrase
func (r *Response) GetReasonPhrase() string {
	status := r.resp.Status
	arr := strings.Split(status, " ")

	return arr[1]
}

// IsTimeout get if request is timeout
func (r *Response) IsTimeout() bool {
	if r.err == nil {
		return false
	}
	netErr, ok := r.err.(net.Error)
	if !ok {
		return false
	}
	if netErr.Timeout() {
		return true
	}

	return false
}

// GetHeaders get response headers
func (r *Response) GetHeaders() map[string][]string {
	return r.resp.Header
}

// HasHeader get if header exsits in response headers
func (r *Response) HasHeader(name string) bool {
	headers := r.GetHeaders()
	for k := range headers {
		if strings.ToLower(name) == strings.ToLower(k) {
			return true
		}
	}

	return false
}

// GetHeader get response header
func (r *Response) GetHeader(name string) []string {
	headers := r.GetHeaders()
	for k, v := range headers {
		if strings.ToLower(name) == strings.ToLower(k) {
			return v
		}
	}

	return nil
}

// GetHeaderLine get a single response header
func (r *Response) GetHeaderLine(name string) string {
	header := r.GetHeader(name)
	if len(header) > 0 {
		return header[0]
	}

	return ""
}
