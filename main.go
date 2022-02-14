package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Hobby string `json:"hobby"`
}

type ResponseFormat struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var errorJson = ResponseFormat{
	Code:    400,
	Message: "Got some errors",
}

func main() {
	db, err := sql.Open("mysql", "root@tcp(mysql_db)/docker")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Second * 120)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	defer db.Close()

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")

		if r.Method == http.MethodGet {
			var users []User
			query, err := db.Query("SELECT id, name, email, hobby FROM users")
			if err != nil {
				panic(err)
			}

			for query.Next() {
				var user User
				query.Scan(&user.Id, &user.Name, &user.Email, &user.Hobby)

				users = append(users, user)
			}

			defer query.Close()

			err = json.NewEncoder(w).Encode(users)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			return
		} else if r.Method == http.MethodPost {
			var user struct {
				Name  string `json:"name"`
				Email string `json:"email"`
				Hobby string `json:"hobby"`
			}
			err := json.NewDecoder(r.Body).Decode(&user)
			if err == io.EOF {
				json.NewEncoder(w).Encode(ResponseFormat{http.StatusBadRequest, "Body cannot be empty"})
				return
			}

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			query, err := db.Prepare("INSERT INTO users (name, email, hobby) VALUES (?, ?, ?)")
			query.Exec(&user.Name, &user.Email, &user.Hobby)
			if err != nil {
				panic(err)
			}

			json.NewEncoder(w).Encode(ResponseFormat{201, "Data inserted"})

			return
		}

		json.NewEncoder(w).Encode(errorJson)
	})

	fmt.Println("Server running on port 3000")

	http.ListenAndServe("0.0.0.0:3000", nil)
}
