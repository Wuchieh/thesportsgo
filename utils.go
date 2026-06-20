package thesportsgo

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
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

func newRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, url, body)
}

// 是否為絕對路徑
func pathIsAbs(path string) bool {
	u, _ := url.Parse(path)
	return u != nil && u.IsAbs()
}

func handleContext(ctx context.Context) context.Context {
	if ctx != nil {
		return ctx
	}
	return context.Background()
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func toQuery(a any) url.Values {
	marshal, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	var resultMap map[string]any
	err = json.Unmarshal(marshal, &resultMap)
	if err != nil {
		panic(err)
	}

	result := make(url.Values)
	for k, v := range resultMap {
		vValue := reflect.ValueOf(v)
		if vValue.Kind() == reflect.Slice {
			vLen := vValue.Len()
			vSlice := make([]string, vLen)
			for i := 0; i < vLen; i++ {
				vSlice[i] = fmt.Sprint(vValue.Index(i))
			}
			v = strings.Join(vSlice, ",")
		}
		val := fmt.Sprint(v)
		if val == "" {
			continue
		}
		result.Add(k, val)
	}

	return result
}

func toPtr[T any](t T) *T {
	return &t
}

func secretGet[T Response[E], E any](ctx context.Context, c *Client, path string, query ...url.Values) (*T, error) {
	resp, err := c.SecretGet(ctx, path, query...)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result MetaResponse[E]
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.hasError() {
		return nil, result
	}

	return toPtr(T(result.toResponse())), nil
}
