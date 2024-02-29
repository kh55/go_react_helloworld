package controller

import (
	"go-sample/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"log"
	"context"
)

var ctx = context.Background()

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is index!!!!",
	})
}

func Hello(context *gin.Context) {
	var c = service.Init()
	fmt.Printf("%T\n", c)
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func RedisTest(context *gin.Context) {
	rdb := service.Init()
	fmt.Printf("%T\n", rdb)

	fmt.Printf("%s\n", "set")
	key := "SetTestKey"
	value := "Hallo World Go!"
	
	// SETコマンドでkey-valueペアを設定
	err := rdb.Set(key, value, 0).Err()
	if err != nil {
		log.Fatalf("Failed to set key: %v", err)
	}

	fmt.Printf("%s\n", "get")
	// GETコマンドでkeyのvalueを取得
	val, err := rdb.Get(key).Result()
	if err != nil {
		log.Fatalf("Failed to get key: %v", err)
	}
	fmt.Println("key:", val)

	context.JSON(http.StatusOK, gin.H{
		"message": val,
	})
}

type InputString struct{
	InputData string
}

func RedisInput(context *gin.Context) {
	rdb := service.Init()
	fmt.Printf("%T\n", rdb)

	var hoge InputString
	context.BindJSON(&hoge)

	key := "InputData"
	value := hoge.InputData
	
	// SETコマンドでkey-valueペアを設定
	err := rdb.Set(key, value, 0).Err()
	if err != nil {
		log.Fatalf("Failed to set key: %v", err)
	}

	fmt.Printf("%s\n", "get")
	// GETコマンドでkeyのvalueを取得
	val, err := rdb.Get(key).Result()
	if err != nil {
		log.Fatalf("Failed to get key: %v", err)
	}
	fmt.Println("key:", val)

	context.JSON(http.StatusOK, gin.H{
		"message": val,
	})
}
