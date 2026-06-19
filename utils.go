package thesportsgo

import (
	"io"
	"net/http"
	"net/url"
)

func handleQuery(query ...url.Values) string {
	result := make(url.Values)
	for _, values := range query {
		for k, v := range values {
			data := result[k]
			result[k] = append(data, v...)
		}
	}
	return result.Encode()
}

func newRequest(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}

// 是否為絕對路徑
func pathIsAbs(path string) bool {
	u, _ := url.Parse(path)
	return u != nil && u.IsAbs()
}
