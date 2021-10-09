package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/koseburak/marvel-universe-web/config"
	"github.com/koseburak/marvel-universe-web/marvel"
	"github.com/koseburak/marvel-universe-web/model"
)

var tpl = template.Must(template.ParseFiles("index.html"))

type Search struct {
	Query      string
	TotalPages int
	Results    *model.MarvelResponse `json:"MarvelResponse"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buf.WriteTo(w)
}

func searchHandler(marvelClient *marvel.MarvelClient) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		u, err := url.Parse(r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := u.Query()
		character := params.Get("character")
		log.Println("Entered character: ", character)

		resultCharacters, err := marvelClient.GetCharacters(character)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		search := &Search{
			Query:      character,
			TotalPages: resultCharacters.Data.Total,
			Results:    resultCharacters,
		}

		buf := &bytes.Buffer{}
		err = tpl.Execute(buf, search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		buf.WriteTo(w)
	}
}

func main() {

	conf, err := config.Config()
	if err != nil {
		log.Println("Got error while loading env config: ", err)
	}

	log.Println(conf)
	log.Println("Listening Serve Port: ", conf.Port)

	defaultHTTPClient := &http.Client{}
	marvelClient := marvel.NewMarvelClient(conf, defaultHTTPClient)

	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/search", searchHandler(marvelClient))
	http.ListenAndServe(":"+conf.Port, mux)
}
