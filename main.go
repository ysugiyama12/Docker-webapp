package main

import(
	"log"
	"net/http"
	// "path/filepath"
	"sync"
	"text/template"
	"fmt"
	"encoding/json"
	_ "github.com/lib/pq"
	"database/sql"
	"strings"
	"os"
	"strconv"
)

type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}

type Ping struct {
    Message string	`json:"message"`
}

type User struct {
    Id   int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type post_res struct {
	Name string
	Email string
}


func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var Db *sql.DB
	Db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_NAME")))
    if err != nil {
        log.Fatal(err)
	}
	url := r.URL.String()
	if r.Method == "GET" {
		slice := strings.Split(url, "/")
		if url == "/" {
			ping := Ping{"Hello, World!!"}
			res, err := json.Marshal(ping)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(res)
	
		}else if len(slice) == 2 || (len(slice) == 3 && slice[2] == ""){
			if slice[1] == "users" {
				rows, err := Db.Query("select id, name, email from my_user")
				if err != nil {
					log.Fatal(err)
				}
				var es []User
				for rows.Next() {
					var e User
					rows.Scan(&e.Id, &e.Name, &e.Email)
					es = append(es, e)
				}
				res, err := json.Marshal(es)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(res)
				return
			}
		}else if len(slice) == 3 {
			if slice[1] == "users" {
				rows, err := Db.Query("select id, name, email from my_user where id = " + slice[2])
				if err != nil {
					log.Fatal(err)
				}
				var es []User
				for rows.Next() {
					var e User
					rows.Scan(&e.Id, &e.Name, &e.Email)
					es = append(es, e)
				}
				res, err := json.Marshal(es[0])
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(res)
				return
			}

		}

	}else if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var t post_res
		err := decoder.Decode(&t)
		if err != nil {
			log.Fatal(err)
		}
		query := "insert into my_user (name, email) values ($1,$2) returning id"
		stmt, err := Db.Prepare(query)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()
		var tt User
		err = stmt.QueryRow(t.Name, t.Email).Scan(&tt.Id)
		tt.Name = t.Name
		tt.Email = t.Email
		res, err := json.Marshal(tt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(res)
		return
	}
}

func main() {
	// port := os.Getenv("PORT")
	port, _ := strconv.Atoi(os.Args[1])
	// if port == "" {
	// 	log.Fatal("$POST must be set")
	// }

	http.Handle("/", &templateHandler{})
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}