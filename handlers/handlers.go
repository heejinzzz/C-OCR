package handlers

import (
	"C-OCR/paddleOCRclient"
	"C-OCR/redisDB"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"strings"
)

// HomePage C-OCR主页
func HomePage(c *gin.Context) {
	c.HTML(200, "homePage.html", gin.H{})
}

// OCR 保存前端传回的图片，对其进行OCR文字识别，将识别出的文本内容和结果图片的保存路径返回给前端
func OCR(c *gin.Context) {
	image, err := c.FormFile("image")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	imageName := image.Filename
	imageNameSplit := strings.Split(imageName, ".")
	imageType := imageNameSplit[len(imageNameSplit)-1]

	// 将图片按序号重命名，保存在 static/originImages/ 路径下
	imageID := redisDB.GetImageNum()
	err = redisDB.ImageNumAddOne()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	imageNewName := "./static/originImages/" + imageID + "." + imageType
	err = c.SaveUploadedFile(image, imageNewName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// 对图片进行OCR文字识别，将结果图片保存到 static/resultImages/ 路径下
	imageBytes, err := ioutil.ReadFile(imageNewName)
	resultText, resultImageBytes := paddleOCRclient.DoOCR(imageBytes, imageType)
	resultImageName := "./static/resultImages/" + imageID + "." + imageType
	_, err = os.OpenFile(resultImageName, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = ioutil.WriteFile(resultImageName, resultImageBytes, 0777)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// 将识别出的文本内容和结果图片的保存路径返回给前端
	response := gin.H{
		"resultText":      resultText,
		"resultImageName": "/resultImages/" + imageID + "." + imageType,
	}
	c.JSON(200, response)
}
