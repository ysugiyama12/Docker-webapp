package main

import(
	"log"
	"net/http"
	// "path/filepath"
	"sync"
	"text/template"
	// "fmt"
	"encoding/json"
)

type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}

type Ping struct {
    Message string	`json:"message"`
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ping := Ping{"Hello, World!!"}
		res, err := json.Marshal(ping)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
		return
	}
	
	// t.once.Do(func() {
	// 	t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	// })
	// t.templ.Execute(w, nil)
}

func main() {
	http.Handle("/", &templateHandler{})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}