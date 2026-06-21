package thesportsgo_test

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestJSON(t *testing.T) {
	jsonAnalytics := func(name string) error {
		filePath := getPath(name)
		data, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("讀取: %s 失敗: %w", name, err)
		}
		type jData struct {
			Results any `json:"results"`
		}
		var j jData
		if err := json.Unmarshal(data, &j); err != nil {
			return fmt.Errorf("反序列化: %s 失敗: %w", name, err)
		}

		l := reflect.ValueOf(j.Results).Len()

		if l == 0 {
			return fmt.Errorf("%s 資訊異常", name)
		}

		return nil
	}
	dir := must(os.ReadDir(getPath()))
	for _, entry := range dir {
		if strings.HasSuffix(entry.Name(), ".json") {
			if err := jsonAnalytics(entry.Name()); err != nil {
				t.Error(err)
			}
		}
	}
}
