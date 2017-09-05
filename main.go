package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"

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

# sleek UI/UX must be implemented for all the front end
	effects and functionalities.

*/

type Book struct {
	Title  string
	Author string
}

func main() {

	r := httprouter.New()
	r.GET("/", HomeHandler)
	r.POST("/posts", PostsCreateHandler)
	r.ServeFiles("/static/*filepath", http.Dir("./static"))

	fmt.Println("Starting	server	on	:8080")
	http.ListenAndServe(":8080", r)

}

func HomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// load data from json file

	// f, err := os.Open("./database/data.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	str := `[
		{
			"question": "What type of weather you wanna experience",
			"question_type": "select",
			"options": [
				{
					"tags": [
						"cold",
						"cool",
						"ice"
					],
					"image": "static/images/data/1.jpeg"
				},
				{
					"tags": [
						"hot",
						"humid",
						"warm"
					],
					"image": "static/images/data/2.jpeg"
				},
				{
					"tags": [
						"medium",
						"normal"
					],
					"image": "static/images/data/2.jpeg"
				}
			]
		}
	]`

	var S []model.Question
	json.Unmarshal([]byte(str), &S)

	fp := path.Join("views", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, S); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "posts	create")
}
