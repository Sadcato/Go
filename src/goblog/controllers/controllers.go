package pages

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"text/template"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type ArticlesFormData struct {
	Title, Body string
	URL         *url.URL
	Errors      map[string]string
}

func SetRouter(r *mux.Router) {
	Router = r
}

func ArticlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "访问文章列表")
}

func ArticlesCreateHandler(w http.ResponseWriter, r *http.Request) {

	storeURL, _ := Router.Get("articles.store").URL()
	data := ArticlesFormData{
		Title:  "",
		Body:   "",
		URL:    storeURL,
		Errors: nil,
	}
	tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 欢迎来到 goblog</h1>")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</p>")
}
