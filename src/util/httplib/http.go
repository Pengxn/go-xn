package httplib

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

// Client http request client provides more useful methods
// for requesting one url than http.Request.
type Client struct {
	config      *Config
	httpClient  *http.Client // go http client
	cookie      []*http.Cookie
	header      map[string]string
	params      map[string][]string
	host        string
	contentType ContentType
	keepCase    bool
}

// New returns http Client with default settings.
func New() *Client {
	return NewClient(NewHttpConfig())
}

// NewTimeoutClient returns http Client with custom timeout.
func NewTimeoutClient(second int64) *Client {
	config := NewHttpConfig()
	config.Timeout = time.Duration(second) * time.Second
	return NewClient(config)
}

// NewClient returns http Client with custom settings.
func NewClient(config *Config) *Client {
	transport := &http.Transport{
		DisableKeepAlives: config.DisableKeepAlives,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: config.InsecureSkipVerify,
			Certificates:       config.Certificates,
		},
	}

	return &Client{
		header: make(map[string]string),
		params: make(map[string][]string),
		config: config,
		httpClient: &http.Client{
			Timeout:   config.Timeout,
			Transport: transport,
		},
	}
}

// SetCookie add cookie into request.
func (c *Client) SetCookie(cookie *http.Cookie) *Client {
	c.cookie = append(c.cookie, cookie)
	return c
}

// SetHeader sets http header.
// Keep the first letter of param is capitalized.
func (c *Client) SetHeader(h map[string]string) *Client {
	c.header = h
	return c
}

// SetHeaderFace sets http header.
// Keep param is original words, don't change the case of letters.
// https://golang.org/pkg/net/http/#CanonicalHeaderKey
func (c *Client) SetHeaderFace(h map[string]string) *Client {
	c.header = h
	c.keepCase = true
	return c
}

// SetParams adds http query param in to request.
func (c *Client) SetParams(p map[string][]string) *Client {
	c.params = p
	return c
}

// SetHost sets host.
func (c *Client) SetHost(host string) *Client {
	c.host = host
	return c
}

// SetContentType sets "Content-Type" header field.
func (c *Client) SetContentType(contentType ContentType) *Client {
	c.contentType = contentType
	return c
}

// GET sends request with GET method.
func (c *Client) GET(url string) (*http.Response, error) {
	values := c.httpParamsDeal()
	values.Encode()
	if len(c.params) > 0 {
		if !strings.Contains(url, "?") {
			url += "?"
		} else {
			url += "&"
		}
	}
	url += c.httpValuesDeal(values)
	return c.do("GET", url, nil)
}

// POST sends request with POST method.
func (c *Client) POST(url string) (*http.Response, error) {
	data := c.httpValuesDeal(c.httpParamsDeal())
	c.header["Content-Type"] = "application/x-www-form-urlencoded;charset=" + c.config.Charset
	return c.do("POST", url, strings.NewReader(data))
}

// PUT sends request with PUT method.
func (c *Client) PUT(url string) (*http.Response, error) {
	data := c.httpValuesDeal(c.httpParamsDeal())
	return c.do("PUT", url, strings.NewReader(data))
}

// DELETE sends request with DELETE method.
func (c *Client) DELETE(url string) (*http.Response, error) {
	data := c.httpValuesDeal(c.httpParamsDeal())
	return c.do("DELETE", url, strings.NewReader(data))
}

// PATCH sends request with PATCH method.
func (c *Client) PATCH(url string) (*http.Response, error) {
	data := c.httpValuesDeal(c.httpParamsDeal())
	return c.do("PATCH", url, strings.NewReader(data))
}

// do sends request with custom header and specified method.
func (c *Client) do(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return &http.Response{}, err
	}
	if len(c.header) > 0 {
		c.createHeader(req)
	}
	for _, v := range c.cookie {
		req.AddCookie(v)
	}
	if c.host != "" {
		req.Host = c.host
	}

	return c.httpClient.Do(req)
}

// PostRaw sends POST request with raw content.
func (c *Client) PostRaw(url string, data []byte) (*http.Response, error) {
	contentType := Raw
	if c.contentType != "" {
		contentType = c.contentType
	}
	c.header["Content-Type"] = fmt.Sprintf("%s; charset=%s", contentType, c.config.Charset)
	return c.do("POST", url, bytes.NewReader(data))
}

// PostFile sends POST request with file content.
func (c *Client) PostFile(url, fileName, field string, fileContent []byte) (*http.Response, error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	c.postFileValuesDeal(bodyWriter, c.params)
	fileWriter, err := bodyWriter.CreateFormFile(field, fileName)
	if err != nil {
		return &http.Response{}, err
	}
	if _, err = fileWriter.Write(fileContent); err != nil {
		return &http.Response{}, err
	}
	bodyWriter.Close()
	c.header["Content-Type"] = bodyWriter.FormDataContentType()
	return c.do("POST", url, bodyBuf)
}

func (c *Client) createHeader(req *http.Request) {
	if c.keepCase {
		h := make(map[string][]string)
		for key, value := range c.header {
			h[key] = []string{value}
		}
		req.Header = h
	} else {
		for key, value := range c.header {
			req.Header.Set(key, value)
		}
	}
}

func (c *Client) postFileValuesDeal(w *multipart.Writer, v map[string][]string) {
	isEnCode := c.config.Encode
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if len(vs) > 1 {
			k = k + "[]"
		}
		for _, v := range vs {
			if isEnCode {
				v = url.QueryEscape(v)
			}
			if w.WriteField(k, v) != nil {
				break
			}
		}
	}
}

func (c *Client) httpValuesDeal(v url.Values) string {
	isEnCode := c.config.Encode
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if !c.config.IsEncodeForGo {
			if len(vs) > 1 {
				k = k + "[]"
			}
		}
		prefix := url.QueryEscape(k) + "="
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			if isEnCode {
				buf.WriteString(url.QueryEscape(v))
			} else {
				buf.WriteString(v)
			}
		}
	}
	return buf.String()
}

// httpParamsDeal to handle URL params.
func (c *Client) httpParamsDeal() url.Values {
	values := url.Values{}
	for k, v := range c.params {
		for _, val := range v {
			values.Add(k, val)
		}
	}
	return values
}
