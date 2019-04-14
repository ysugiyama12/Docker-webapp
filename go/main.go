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
    Id   int
	Name string
	Email string
}

type User2 struct {
    Id   int
	Name string
	Email string
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var Db *sql.DB
    Db, err := sql.Open("postgres", "host=db-sugiyama user=root password=root dbname=app_db sslmode=disable")
    if err != nil {
        log.Fatal(err)
	}
	fmt.Println(r.URL)
	url := r.URL.String()
	if r.Method == "GET" {
		slice := strings.Split(url, "/")
		if len(slice) == 2 || (len(slice) == 3 && slice[2] == ""){
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
				fmt.Printf("%v", es)
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
				fmt.Printf("%v", es)
				res, err := json.Marshal(es[0])
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(res)
				return
			}

		}
		// for _, str := range slice {
		// 	fmt.Println(str)
		// }

	}
	// sql := "SELECT id, name, email FROM my_user WHERE id=$1;"
	// pstatement, err := Db.Prepare(sql)
    // if err != nil {
    //     log.Fatal(err)
    // }
	// 検索パラメータ（ユーザID）
	// queryID := 1
	// 検索結果格納用の TestUser
	// var User1 User

	// // queryID を埋め込み SQL の実行、検索結果1件の取得
	// err = pstatement.QueryRow(queryID).Scan(&User1.Id, &User1.Name, &User1.Email)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 検索結果の表示
	// fmt.Println(User.Id, User.Name, User.Email)


	if r.Method == "GET" {
		// ping := Ping{"Hello, World!!"}


	}

}

func main() {
	http.Handle("/", &templateHandler{})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}