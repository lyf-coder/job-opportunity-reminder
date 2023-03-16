package tpl

import (
	"bytes"
	"log"
	"text/template"
)

// 缓存模版
var tplMap = make(map[string]*template.Template)

func Load(name string, filePath string) (*template.Template, error) {
	// 加载模版
	replyTemplate, err := template.New(name).ParseFiles(filePath)
	if err != nil {
		log.Println("加载模版失败！", name, filePath, err)
		return nil, err
	} else {
		tplMap[name] = replyTemplate
	}
	return replyTemplate, nil
}

// Get 获取模版
func Get(name string, filePath string) (*template.Template, error) {
	if replyTemplate, ok := tplMap[name]; ok {
		return replyTemplate, nil
	}
	return Load(name, filePath)
}

// GetTemplateResultStr 处理模版-name是文件名
func GetTemplateResultStr(name string, filePath string, data any) string {
	replyTemplate, err := Get(name, filePath)
	if err != nil {
		log.Println("模版获取失败！", err)
	}
	replyBuff := new(bytes.Buffer)
	err = replyTemplate.Execute(replyBuff, data)
	if err != nil {
		log.Println(err)
	}
	return replyBuff.String()
}

// GetTplPath 获取模版文件路径-相对于main函数位置
func GetTplPath(tmplFileName string) string {
	return "tpl/" + tmplFileName
}
