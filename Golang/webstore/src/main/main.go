package main

import (
	"net/http"
	"os"
	"text/template"
	"controllers"
)

func main() {

	templates := populateTemplates() // 缓存相关的模板
	controllers.Register(templates)  // 路由对应的处理函数的注册

	http.ListenAndServe(":8989", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates")
	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)

	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}
	result.ParseFiles(*templatePaths...)

	return result
}
