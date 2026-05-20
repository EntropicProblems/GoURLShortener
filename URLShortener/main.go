package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
)

type Page struct{
	Title string
	Body []byte
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
    //need to take a request for a specific path in our domain 
    //and check and see if it maps to any url's
    if r.Method == http.MethodGet {
        if r.URL.Path == "/" {
            //do nothing
        } else {
            http.Redirect(w,r,"https://google.com",301) //half-built redirect function
        }
        //if the domain is root, direct them to a page to designate a new url to shorten
        fmt.Fprint(w, "Welcome to root")
    }
    if r.Method == http.MethodPost {
        //this will handle our shortener submission, adding to our map
    }
}

func main(){
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}