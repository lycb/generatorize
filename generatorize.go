package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
	"math/rand"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

type Response struct {
	GeneratedName string
}

type Name struct {
	Prefix string
	Postfix string
}

func generateName(delimiter, prefix, postfix string) string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s%s%s%s%s", prefix, delimiter, postfix, delimiter, funthings[rand.Intn(len(funthings))])
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Error parsing form")
	}
	prefix := r.FormValue("prefix")
	postfix := r.FormValue("postfix")

	resp := Response{
		GeneratedName: generateName("-", prefix, postfix),
	}

	tpl.Execute(w, resp)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", indexHandler)
	// mux.HandleFunc("/return-generate", generateHandler)
	mux.HandleFunc("/generate", generateHandler)
	http.ListenAndServe(":"+port, mux)
}
