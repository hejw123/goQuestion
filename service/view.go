package service

import (
	"fmt"
	"html/template"
	"net/http"
)

func Show ( w http.ResponseWriter , r *http.Request ) {
	w.Header().Set("text/html","charset=utf-8")
	t , err := template.ParseFiles("template/question.html")
	if err != nil {
		fmt.Fprint(w , "file no find")
	}
	t.Execute(w , nil)
}