package receiver

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestPost(t *testing.T) {
	const resultStr = "Hello World!"
	// 模拟一个返回
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {

		switch strings.TrimSpace(r.URL.Path) {
		case "/":
			_, _ = io.WriteString(w, resultStr)

		default:
			http.NotFoundHandler().ServeHTTP(w, r)
		}
	}))

	type args struct {
		url  string
		body interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "测试发送给接收者信息",
			args: args{
				url:  server.URL,
				body: "test",
			},
			want:    []byte(resultStr),
			wantErr: false,
		},
		{
			name: "测试发送给接收者信息-错误的url",
			args: args{
				url: "/test",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Post(tt.args.url, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Post() got = %v, want %v", got, tt.want)
			}
		})
	}
}
