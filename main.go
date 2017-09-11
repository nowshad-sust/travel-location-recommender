package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/julienschmidt/httprouter"
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
	// r.POST("/post", PostsCreateHandler)

	r.ServeFiles("/static/*filepath", http.Dir("./static"))

	fmt.Println("Starting	server	on	:8080")
	http.ListenAndServe(":8080", r)

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
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", JSON)
}

// func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	r.ParseForm()
// 	for _, val := range r.Form {
// 		for _, tag := range val {
// 			strs := strings.Split(strings.TrimSpace(tag), " ")
// 			fmt.Println(strs)
// 			// for _, str := range strs {
// 			// 	fmt.Println(str, len(str))
// 			// }
// 		}
// 	}
// 	fmt.Fprintln(rw, r.Form)
// }
