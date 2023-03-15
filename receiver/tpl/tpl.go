package tpl

import (
	"bytes"
	"log"
	"text/template"
)

// 缓存模版
var tplMap = make(map[string]*template.Template)

// GetTemplateResultStr 处理模版-name是文件名
func GetTemplateResultStr(name string, filePath string, data any) string {
	replyTemplate, ok := tplMap[name]
	var err error
	if !ok {
		log.Println("新建模版")
		replyTemplate, err = template.New(name).ParseFiles(filePath)
		if err != nil {
			log.Println(err)
		} else {
			tplMap[name] = replyTemplate
		}

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
	return "receiver/tpl/" + tmplFileName
}
