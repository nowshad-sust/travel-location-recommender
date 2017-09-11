package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"../ibelongto/model"

	"github.com/julienschmidt/httprouter"
)

/**
this is a location recommendation system
this system will suggest user where should
they belong based on their feedback.

# every question should have multiple choices
# each option will consist some text or image
# eatch option should value some tags
# tag subjects: weather,
				food,
				culture,
				men/women,
				places/travel,
				jobs/occupation,

# based on the tags, find the most closet matched
	location/place implementing some basic machine
	learning techniques.

# questions can be static, and all of the form should
	be loaded in the front end without creating any
	backend load. Just the final tags array should be
	posted to the back end and wait for response.

# sleek UI/UX must be implemented for of all the front end
	effects and functionalities.

*/
var DATA []model.Question
var JSON []byte

func getData() []model.Question {
	// load data from json file
	file, e := ioutil.ReadFile("./database/data.json")

	JSON = file
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	var qs []model.Question
	json.Unmarshal([]byte(file), &qs)
	DATA = qs
	return qs
}

func main() {

	// load data
	getData()

	r := httprouter.New()
	r.GET("/", HomeHandler)
	r.GET("/json", JsonResponse)
	r.POST("/post", PostsCreateHandler)
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

	if err := tmpl.Execute(w, DATA); err != nil {
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

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	for _, val := range r.Form {
		for _, tag := range val {
			strs := strings.Split(strings.TrimSpace(tag), " ")
			fmt.Println(strs)
			// for _, str := range strs {
			// 	fmt.Println(str, len(str))
			// }
		}
	}
	fmt.Fprintln(rw, r.Form)
}
