package thesportsgo

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Response[T any] struct {
	Code    int `json:"code"`
	Results T   `json:"results"`
	raw     []byte
}

// Raw 取得原數據
func (r Response[T]) Raw() []byte {
	return append([]byte(nil), r.raw...)
}

// MetaResponse 原始返回結構
type MetaResponse[T any] struct {
	raw     []byte
	Code    int     `json:"code"`
	Results T       `json:"results"`
	Err     *string `json:"err,omitempty"`
	Msg     *string `json:"msg,omitempty"`
}

// Raw 取得原始值
func (r MetaResponse[T]) Raw() []byte {
	return append([]byte(nil), r.raw...)
}

func (r *MetaResponse[T]) UnmarshalJSON(bytes []byte) error {
	type tmp struct {
		Code    int     `json:"code"`
		Results T       `json:"results"`
		Err     *string `json:"err,omitempty"`
		Msg     *string `json:"msg,omitempty"`
	}
	var data tmp
	if err := json.Unmarshal(bytes, &data); err != nil {
		return fmt.Errorf("json unmarshal fail: %w, raw: %s", err, string(bytes))
	}
	r.Code = data.Code
	r.Results = data.Results
	r.Err = data.Err
	r.Msg = data.Msg
	r.raw = append([]byte(nil), bytes...)
	return nil
}

func (r MetaResponse[T]) toResponse() Response[T] {
	return Response[T]{
		Code:    r.Code,
		Results: r.Results,
		raw:     r.raw,
	}
}

func (r MetaResponse[T]) Error() string {
	var msg []string
	if r.Code != 0 {
		msg = append(msg, "code: "+strconv.Itoa(r.Code))
	}
	if r.Err != nil && *r.Err != "" {
		msg = append(msg, "err: "+*r.Err)
	}
	if r.Msg != nil && *r.Msg != "" {
		msg = append(msg, "msg: "+*r.Msg)
	}
	msg = append(msg, "raw: "+string(r.raw))
	return strings.Join(msg, ", ")
}

func (r MetaResponse[T]) hasError() bool {
	return r.Err != nil && *r.Err != ""
}
