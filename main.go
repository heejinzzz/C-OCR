package main

import (
	"C-OCR/handlers"
	"github.com/gin-gonic/gin"
)

var C_OCRAddress = "127.0.0.1:9091"

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./static/htmlFiles/*")

	r.GET("/", handlers.HomePage)
	r.POST("/ocr", handlers.OCR)

	r.Run(C_OCRAddress)
}
