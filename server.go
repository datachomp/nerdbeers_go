package main

import (
    "fmt"
    "html"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/suggestions", SuggestionsIndex)
    router.HandleFunc("/suggestions/{suggestionId}", SuggestionsShow)

    log.Println("Listening on localhost:8080...")
    log.Fatal(http.ListenAndServe(":8080", router))
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
