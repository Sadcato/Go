package main

import (
	pages "goblog/controllers"
	"goblog/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// 先注册所有路由
	router.HandleFunc("/", pages.HomeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", pages.AboutHandler).Methods("GET").Name("about")
	router.HandleFunc("/articles/{id:[0-9]+}", pages.ArticlesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", pages.ArticlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", pages.ArticlesStoreHandler).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create", pages.ArticlesCreateHandler).Methods("GET").Name("articles.create")

	// 设置路由器到控制器包中
	pages.SetRouter(router)

	// 自定义 404 页面
	router.NotFoundHandler = http.HandlerFunc(pages.NotFoundHandler)

	// 中间件
	router.Use(middlewares.ForceHTML)

	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
