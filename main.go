package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	Id          int
	Login       string
	MoneyAmount int
	CardNumber  string
	Status      int
}

/* func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello! Welcome to online banking application by Anastasia Sheremeteva =]")
} */

func loginsPage(w http.ResponseWriter, r *http.Request) {
	login_param := r.URL.Query().Get("login")
	if login_param == "" {
		fmt.Fprint(w, "Bad login")
		return
	}

	db, err := sql.Open("mysql", "root:12345@/bank")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	prep, err := db.Prepare("SELECT * FROM users WHERE login = ?")
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "Bad login")
		panic(err)
	}
	defer prep.Close()
	rows, err := prep.Query(login_param)
	if err != nil {
		fmt.Fprint(w, "Bad login")
		panic(err)
	}
	defer rows.Close()
	if !rows.Next() {
		fmt.Fprint(w, "Bad login")
		return
	}
	var user User
	err = rows.Scan(&user.Id, &user.Login, &user.MoneyAmount, &user.CardNumber, &user.Status)
	if err != nil {
		fmt.Fprint(w, "Bad login")
		panic(err)
	}

	var status string
	if user.Status == 1 {
		status = "active"
	} else {
		status = "passive"
	}
	fmt.Fprintf(w, "Some information about requested user\nID: %d\nLogin: %s\nAmount of money on bank card: %d\nCard number: %s\nStatus: %v",
		user.Id, user.Login, user.MoneyAmount, user.CardNumber, status)
}

func idsPage(w http.ResponseWriter, r *http.Request) {
	id_param := r.URL.Query().Get("id")
	if id_param == "" {
		fmt.Fprint(w, "Bad id")
		return
	}

	db, err := sql.Open("mysql", "root:12345@/bank")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	prep, err := db.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer prep.Close()
	rows, err := prep.Query(id_param)
	if err != nil {
		fmt.Fprint(w, "Bad id")
		panic(err)
	}
	defer rows.Close()
	if !rows.Next() {
		fmt.Fprint(w, "Bad id")
		return
	}
	var user User
	err = rows.Scan(&user.Id, &user.Login, &user.MoneyAmount, &user.CardNumber, &user.Status)
	if err != nil {
		fmt.Fprint(w, "Bad id")
		panic(err)
	}

	var status string
	if user.Status == 1 {
		status = "active"
	} else {
		status = "passive"
	}
	fmt.Fprintf(w, "Some information about requested user\nID: %d\nLogin: %s\nAmount of money on bank card: %d\nCard number: %s\nStatus: %v",
		user.Id, user.Login, user.MoneyAmount, user.CardNumber, status)
}

func getUsersPage(w http.ResponseWriter, r *http.Request) {
	// active users : id, login
	db, err := sql.Open("mysql", "root:12345@/bank")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	usersRes, err := db.Query("SELECT id, login FROM users WHERE status=1")
	if err != nil {
		panic(err)
	}
	defer usersRes.Close()

	fmt.Fprint(w, "List of active users:\n\n")
	for usersRes.Next() {
		var id int
		var login string
		err := usersRes.Scan(&id, &login)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "id: %d\tlogin: %s\n", id, login)
	}
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./static")))
	//r.HandleFunc("/", homePage).Methods("GET")
	r.HandleFunc("/users", getUsersPage)
	r.HandleFunc("/by-login", loginsPage)
	r.HandleFunc("/by-id", idsPage)

	return r
}

func main() {
	db, err := sql.Open("mysql", "root:12345@/bank")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	r := NewRouter()
	http.ListenAndServe(":8080", r)
}
