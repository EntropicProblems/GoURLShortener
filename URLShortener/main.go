package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
    "encoding/json"
)

type Page struct{
	Title string
	Body []byte
}

func (p *Page) save() error {
    filename := p.Title + ".html"
    return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".html" //pages will be in html
    body, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

var myshortcuts = make(map[string]string)//fix this and then we can get cooking with POST

func handler(w http.ResponseWriter, r *http.Request) {
    //need to take a request for a specific path in our domain 
    //and check and see if it maps to any url's
    if r.Method == http.MethodGet {
        if r.URL.Path == "/" || r.URL.Path == "/login" || r.URL.Path =="/register" {
            //do nothing
        } else {
            if myshortcuts == nil {
                body, err := os.ReadFile("shortcuts.json")
                if err != nil {
                    fmt.Fprint(w, "Something went wrong") //adjust this to a nice webpage eventually
                }
                err = json.Unmarshal(body, &myshortcuts)
                if err != nil {
                    fmt.Fprint(w, "Something went wrong") //adjust this to a nice webpage eventually
                }
            }
            http.Redirect(w,r,myshortcuts[r.URL.Path],301) //half-built redirect function
        }
        //if the domain is root, direct them to a page to designate a new url to shorten
        fmt.Fprint(w, "Welcome to root")
        //loadPage("index") //this will be the final idea
    }
    if r.Method == http.MethodPost {
        //this will handle our shortener submission, adding to our map
        if r.FormValue("URLShortcut") == "/login" || r.FormValue("URLShortcut") == "/register" || r.FormValue("URLShortcut") == "/" {
            //these are invalid inputs, so throw them back an error in the form
        } else {
            //add to the map
            if myshortcuts == nil {
                body, err := os.ReadFile("shortcuts.json")
                if err != nil {
                    fmt.Fprint(w, "Something went wrong") //adjust this to a nice webpage eventually
                }
                err = json.Unmarshal(body, &myshortcuts)
                if err != nil {
                    fmt.Fprint(w, "Something went wrong") //adjust this to a nice webpage eventually
                }
            }
            _, exists := myshortcuts[r.FormValue("URLShortcut")]
            if exists {
                fmt.Fprint(w, "Whoops! Someone already is using this shortcut!")
            } else {
                myshortcuts[r.FormValue("URLShortcut")] = r.FormValue("ShortcuttedURL")
                //update the json now
                byties, err :=json.Marshal(myshortcuts)
                if err != nil {
                    fmt.Fprint(w, "Something went wrong") //adjust this to a nice webpage eventually
                }
                os.WriteFile("shortcuts.json", byties, 0600)
            }
        }
    }
}

func main(){
    myshortcuts = nil
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}