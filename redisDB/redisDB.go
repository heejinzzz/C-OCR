package redisDB

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var redisServerAddress = "127.0.0.1:6379"
var Conn redis.Conn

func init() {
	var err error
	Conn, err = redis.Dial("tcp", redisServerAddress)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

// GetImageNum 获取当前已保存的图片数量
func GetImageNum() string {
	num, err := redis.String(Conn.Do("get", "C-OCR_ImageNum"))
	if err != nil && err.Error() == "redigo: nil returned" {
		_, err = Conn.Do("set", "C-OCR_ImageNum", 0)
		if err != nil {
			panic(err)
		}
		return "0"
	}
	if err != nil {
		panic(err)
	}
	return num
}

// ImageNumAddOne 将已保存的图片数量加1
func ImageNumAddOne() error {
	num, err := redis.Int(Conn.Do("get", "C-OCR_ImageNum"))
	if err != nil {
		return err
	}
	_, err = Conn.Do("set", "C-OCR_ImageNum", num+1)
	return err
}

// ImageNumMinusOne 将已保存的图片数量减1
func ImageNumMinusOne() error {
	num, err := redis.Int(Conn.Do("get", "C-OCR_ImageNum"))
	if err != nil {
		return err
	}
	_, err = Conn.Do("set", "C-OCR_ImageNum", num-1)
	return err
}

// ImageNumSetZero 将已保存的图片数量置为0
func ImageNumSetZero() error {
	_, err := Conn.Do("set", "C-OCR_ImageNum", 0)
	return err
}
