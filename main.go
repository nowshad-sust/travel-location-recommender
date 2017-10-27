package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"../ibelongto/lib"

	"../ibelongto/model"

	"github.com/julienschmidt/httprouter"
	// "github.com/rs/cors"
)

var JSON []byte
var PLACES []model.Place

func getData() {
	// load options data from json file
	file, e := ioutil.ReadFile("./database/data.json")

	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	JSON = file

	// load places data from json file
	file, er := ioutil.ReadFile("./database/places.json")

	if er != nil {
		fmt.Printf("File error: %v\n", er)
		os.Exit(1)
	}

	json.Unmarshal(file, &PLACES)
}

func main() {

	// load data
	go getData()

	r := httprouter.New()
	r.GET("/", HomeHandler)
	r.GET("/json", JsonResponse)
	r.POST("/v2/post", PostHandler)

	r.ServeFiles("/static/*filepath", http.Dir("./static"))
	r.NotFound = http.FileServer(http.Dir("static/"))

	// handler := cors.Default().Handler(r)

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

	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", JSON)
}

type test_struct struct {
	Tags  []string `json:"data"`
}

func PostHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	
	decoder := json.NewDecoder(r.Body)

	var data = &test_struct{}
	decoder.Decode(&data)

	arr := data.Tags

	var diffArray []int

	for _, place := range PLACES {
		z, ok := lib.Intersect(arr, place.Tags)
		if !ok {
			fmt.Println("Cannot find intersect")
		}
		slice, ok := z.Interface().([]string)
		if !ok {
			fmt.Println("Cannot convert to slice")
		}
		diffArray = append(diffArray, len(slice))
	}

	suggestionIndex, _ := lib.Max(diffArray)

	res, _ := json.Marshal(PLACES[suggestionIndex])
	
	rw.WriteHeader(200)
	fmt.Fprintln(rw, string(res))
}