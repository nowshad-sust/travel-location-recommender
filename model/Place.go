package model

type Place struct {
	Id          int      `josn:"id"`
	Tags        []string `json:"tags"`
	Name        string   `json:"name"`
	Description string   `json:description`
	Images      []string `json:images`
	Video       string   `json:video`
}
