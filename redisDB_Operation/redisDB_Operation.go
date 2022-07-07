package main

import (
	"C-OCR/redisDB"
	"fmt"
)

func main() {
	//err := redisDB.ImageNumSetZero()
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}

	fmt.Println(redisDB.GetImageNum())
}
