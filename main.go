package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

var JSON []byte

func getData() {
	// load data from json file
	file, e := ioutil.ReadFile("./database/data.json")

	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	JSON = file
}

func main() {

	// load data
	getData()

	r := httprouter.New()
	r.GET("/", HomeHandler)
	r.GET("/json", JsonResponse)
	r.POST("/post", PostsCreateHandler)

	r.ServeFiles("/static/*filepath", http.Dir("./static"))

	handler := cors.Default().Handler(r)

	fmt.Println("Starting	server	on	:8080")
	http.ListenAndServe(":8080", handler)

}

func HomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fp := path.Join("views", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, false); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func JsonResponse(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", JSON)
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

	data := r.FormValue("data")
	var jsonBlob = []byte(data)

	// fmt.Println(data)
	arr := [][]interface{}{}
	err := json.Unmarshal(jsonBlob, &arr)
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range arr {
		for _, val := range value {
			tag := fmt.Sprint(val)
			fmt.Println(tag)
		}
	}

	rw.WriteHeader(200)
	fmt.Fprintln(rw, data)
}
