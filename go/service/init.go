package service

import (
	"fmt"
	"os"
	"strconv"
	"log"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

var (
	client *redis.Client //小文字だとNG
)

func Init() *redis.Client{
	// 空のコンテキスト生成
	//ctx := context.Background()
	
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env target")
	}

	addr := os.Getenv("REDIS_ADDR")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")
	dbStr := os.Getenv("REDIS_DB")

	db, err := strconv.Atoi("0")
	if err != nil{
		panic(err)
	}

	fmt.Printf("%s\n", "aaaaaa")
	fmt.Printf("%s\n", addr)
	fmt.Printf("%s\n", port)
	fmt.Printf("%s\n", dbStr)
	fmt.Printf("%s\n", "bbbbbb")
	fmt.Printf("%s\n", err)

	host := fmt.Sprintf("%s:%s", addr, port)
	rdb := redis.NewClient(&redis.Options{
		Addr: host,
		Password: password,
		DB: db,
	})

	// PINGコマンドを送信してRedisサーバの動作を確認
	pong, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println(pong)

	return rdb
}

