package main

import (
	"log"
	"net/http"
	"try/api"
)

func main() {
	http.HandleFunc("/v1/hello", api.HelloHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}
