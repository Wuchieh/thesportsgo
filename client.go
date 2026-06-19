package thesportsgo

import (
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	httpClient *http.Client
	user       string
	secret     string
	baseURL    string
}

// GetClient 取得 http.Client
func (c *Client) GetClient() *http.Client {
	client := c.httpClient
	if client == nil {
		client = http.DefaultClient
	}
	return client
}

func (c *Client) Get(path string, query ...url.Values) (*http.Response, error) {
	q := handleQuery(query...)
	path = c.handlePath(path)
	request, err := newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	request.URL.RawQuery = q
	return c.do(request)
}

func (c *Client) handlePath(path string) string {
	if pathIsAbs(path) {
		return path
	}
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	if strings.HasSuffix(c.baseURL, "/") {
		return c.baseURL + path
	}
	return c.baseURL + "/" + path
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	client := c.GetClient()
	return client.Do(req)
}

type ClientOption func(*Client)

// WithBaseURL 使用自定的 baseURL
// 輸入的 url 需包含 Scheme 以及 Host 不可包含 Path
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithHTTPClient 使用自定的 http.Client
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = client
	}
}

// NewClient 創建客戶端
func NewClient(user, secret string, opts ...ClientOption) *Client {
	client := &Client{
		user:    user,
		secret:  secret,
		baseURL: "https://api.thesports.com/v1",
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
