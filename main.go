package main

import (
	"fmt"
	"net/http"

	"./controller"
	db "./database"

	"github.com/julienschmidt/httprouter"
)

func main() {

	// load data
	go db.GetData()

	r := httprouter.New()
	r.GET("/", controller.HomeHandler)
	r.GET("/json", controller.JsonResponse)
	r.POST("/v2/post", controller.PostHandler)

	r.ServeFiles("/static/*filepath", http.Dir("./static"))
	r.NotFound = http.FileServer(http.Dir("static/"))

	fmt.Println("Starting	server	on	:8080")
	http.ListenAndServe(":8080", r)

}