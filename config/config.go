package config

import (
	"github.com/spf13/viper"
	"log"
)

// IsLoadKey 是否以及加载
const IsLoadKey = "IS_LOAD_KEY"

// LoadConfig 加载配置文件- 传入配置文件路径- 相对路径-相对于执行程序的路径
func LoadConfig(filePath string) {
	// 判断是否已经加载-假如已经加载则不用继续
	if viper.GetBool(IsLoadKey) {
		return
	}

	// 设置文件名
	viper.SetConfigName("config") // name of config file (without extension)
	// 设置配置文件的类型
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name

	// 设置配置文件路径为当前工作目录
	viper.AddConfigPath(filePath) // optionally look for config in the working directory

	log.Println("加载环境变量...")
	log.Println(viper.AllSettings())
	// 环境变量需要大写
	viper.AutomaticEnv()

	// 读取配置
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Println("error", "无法读取配置文件", err)
	}

	// 设置为已加载文件
	viper.Set(IsLoadKey, true)
}
