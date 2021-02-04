package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	html := template.Must(template.ParseFiles("file1", "file2"))
	router.SetHTMLTemplate(html)
	router.Run(":8080")
}
