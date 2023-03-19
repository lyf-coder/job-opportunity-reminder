package config

import (
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// 环境变量要大写生效
	os.Setenv("PROXY_URL", "test")
	os.Setenv("DURATION_SEC", "400")
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "测试加载文件",
			args: args{
				filePath: "../",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoadConfig(tt.args.filePath)
			assert.Equal(t, viper.Get("proxy_url"), "test")
			assert.Equal(t, viper.Get("duration_sec"), "400")
			assert.Equal(t, viper.Get("fei_shu_bot_webhook_url"), "")
		})
	}
}
