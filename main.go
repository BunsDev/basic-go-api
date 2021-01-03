package main

import (
	"fmt"
	"log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"

    "github.com/jonaferreira/go-api/dto"
    
    "github.com/jonaferreira/go-util/util"
)


var Articles = []dto.Article{
    {Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
    {Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}


func homePage(w http.ResponseWriter, r *http.Request){
    isTrue := util.MaxInt(10,20)
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage", isTrue)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    // Loop over all of our Articles
    // if the article.Id equals the key we pass in
    // return the article encoded as JSON
    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    // add our articles route and map it to our 
    // returnAllArticles function like so
    http.HandleFunc("/all", returnAllArticles)
    
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}