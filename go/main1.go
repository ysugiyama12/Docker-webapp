package main

import(
	"log"
	"net/http"
	// "path/filepath"
	"sync"
	"text/template"
	"fmt"
	"encoding/json"
	"strconv"
    "os"
    "database/sql"
    _ "github.com/lib/pq"
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

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var Db *sql.DB
    // Dbの初期化
    Db, err := sql.Open("postgres", "host=localhost user=app_user password=password dbname=my_user sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(Db)
    sql := "SELECT id, name, email FROM my_user WHERE id=$1;"
    pstatement, err := Db.Prepare(sql)
    if err != nil {
        log.Fatal(err)
    }

    // 検索パラメータ（ユーザID）
    queryID := 1
    // 検索結果格納用の TestUser
    var User User

    // queryID を埋め込み SQL の実行、検索結果1件の取得
    err = pstatement.QueryRow(queryID).Scan(&User.Id, &User.Name, &User.Email)
    if err != nil {
        log.Fatal(err)
    }

    // 検索結果の表示
    fmt.Println(User.Id, User.Name, User.Email)



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
    port, _ := strconv.Atoi(os.Args[1])
    // fmt.Println(port)
    fmt.Printf("Starting server at Port %d", port)
    http.Handle("/", &templateHandler{})
    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}