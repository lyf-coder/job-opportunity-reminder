package tpl

import (
	"strconv"
	"testing"
)

func TestGetTplPath(t *testing.T) {
	type args struct {
		tmplFileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试路径拼接",
			args: args{
				"test.json",
			},
			want: "receiver/tpl/test.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTplPath(tt.args.tmplFileName); got != tt.want {
				t.Errorf("GetTplPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTemplateResultStr(t *testing.T) {
	type args struct {
		name     string
		filePath string
		data     any
	}
	type P struct {
		Name string
		Age  int
	}
	p := &P{
		Name: "张三",
		Age:  18,
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试根据模版生成字符串",
			args: args{
				name:     "test.json",
				filePath: "./test.json",
				data:     p,
			},
			want: `{
  "name": ` + p.Name + `,
  "age": ` + strconv.Itoa(p.Age) + `
}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTemplateResultStr(tt.args.name, tt.args.filePath, tt.args.data); got != tt.want {
				t.Errorf("GetTemplateResultStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
