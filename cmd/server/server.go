package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
	"os"
	"nerdbeers_go/Godeps/_workspace/src/github.com/gorilla/mux"
)

type Suggestion struct {
	Suggestion string
	CreatedAt  time.Time
}

type Suggestions []Suggestion

func main() {
	port := os.Getenv("PORT")
	connector := fmt.Sprint(":", port)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/suggestions", SuggestionsIndex)
	router.HandleFunc("/suggestions/{suggestionId}", SuggestionsShow)
	log.Println(port)
	log.Println("Listening on localhost:%v...", port)
	log.Fatal(http.ListenAndServe(connector, router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello NerdBeer, %q", html.EscapeString(r.URL.Path))
}

func SuggestionsIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Suggestion Index!")
}

func SuggestionsShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	suggestionId := vars["suggestionId"]
	fmt.Fprintln(w, "Suggestion show:", suggestionId)
}
