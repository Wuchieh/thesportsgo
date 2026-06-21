package thesportsgo_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/wuchieh/thesportsgo"
)

var (
	PROXY  = os.Getenv("T_PROXY")
	USER   = os.Getenv("T_USER")
	SECRET = os.Getenv("T_SECRET")

	client *thesportsgo.Client
)

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func toPtr[T any](t T) *T {
	return &t
}

func getPath(path ...string) string {
	return filepath.Join(append([]string{"_test", "output"}, path...)...)
}

func debug[T any](data T, err error) func(t *testing.T, name string) {
	return func(t *testing.T, name string) {
		if err != nil {
			t.Error(fmt.Errorf("%s error: %w", name, err))
			return
		}

		m, err := json.Marshal(data)
		if err != nil {
			t.Error(err)
			return
		}

		err = os.WriteFile(getPath(fmt.Sprintf("%d_%s.json", time.Now().Unix(), name)), m, 0o666)
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func init() {
	if err := os.MkdirAll(getPath(), 0o666); err != nil {
		panic(err)
	}
	var op []thesportsgo.ClientOption
	if PROXY != "" {
		u, err := url.Parse(PROXY)
		if err != nil {
			panic(err)
		}
		op = append(op, thesportsgo.WithHTTPClient(&http.Client{
			Transport: &http.Transport{Proxy: http.ProxyURL(u)},
		}))
	}
	client = thesportsgo.NewClient(USER, SECRET, op...)
}
