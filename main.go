package main

import (
	"log"
	"net/http"
	"question/mysql"
	"question/service"
)

func main() {
	mysql.Setup()
	go func() {
		http.HandleFunc("/createQuestion" , service.Create )
		http.HandleFunc("/updateQuestion" , service.Update )
		http.HandleFunc("/listQuestion" , service.List )
		http.HandleFunc("/" , service.Show)
	}()
	log.Fatal(http.ListenAndServe(":88" , nil))
}