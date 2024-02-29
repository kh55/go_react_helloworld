package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
	"log"
	"go-sample/controller"
)

func main() {
	log.Println("start server...")
	r := gin.Default()
	// ここからCorsの設定
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:3001",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	  }))
	
	r.GET("/", controller.Index)
	r.GET("/hello", controller.Hello)
	r.GET("/redis", controller.RedisTest)
	r.POST("/input", controller.RedisInput)
	log.Fatal(r.Run(":3001"))
}
