package controller

import (
	"fmt"
	"encoding/json"
	"net/http"

	"../lib"

	db "../database"

	"github.com/julienschmidt/httprouter"
)

func JsonResponse(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", db.JSON)
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

	for _, place := range db.PLACES {
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

	res, _ := json.Marshal(db.PLACES[suggestionIndex])
	
	rw.WriteHeader(200)
	fmt.Fprintln(rw, string(res))
}