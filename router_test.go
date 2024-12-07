package myrouter

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRouter_GET(t *testing.T) {
	testcases := []struct {
		name string
		endpoint string
		handler http.Handler
	}{
		{
			name: "/のエンドポイントにハンドラを追加する",
			endpoint: "/",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		},
		{
			name: "/helloのエンドポイントにハンドラを追加する",
			endpoint: "/hello",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		},
		{
			name: "/hogeのエンドポイントにハンドラを追加する",
			endpoint: "/hoge",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		},
		{
			name: "/fugaのエンドポイントにハンドラを追加する",
			endpoint: "/fuga",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		},
	}

	r := NewRouter()

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			r.Get(testcase.endpoint, testcase.handler)
		})
	}
}

func TestRouter_Search(t *testing.T) {
	testcases := []struct {
		name string
		endpoint string
		handler http.Handler
	}{
		{
			name: "/のエンドポイントを検索する",
			endpoint: "/",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		},
		{
			name: "/hoge/fugaのエンドポイントを検索する",
			endpoint: "/hoge/fuga",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		},
		{
			name: "/helloのエンドポイントを検索する",
			endpoint: "/hello",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		},
		{
			name: "/hogeのエンドポイントを検索する",
			endpoint: "/hoge",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}),
		},
	}

	r := NewRouter()

	for _, testcase := range testcases {
		r.Get(testcase.endpoint, testcase.handler)
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			handler := r.Search(http.MethodGet, testcase.endpoint)
			// 関数のポインタを比較する
			if reflect.ValueOf(handler).Pointer() != reflect.ValueOf(testcase.handler).Pointer() {
				t.Errorf("ハンドラが異なります\nexpected: %v\nactual: %v", testcase.handler, handler)
			}
		})
	}
}
