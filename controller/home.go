package controller

import (
	"net/http"
	"html/template"
)

type TplVars struct {
	H1 string
}

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tplVars := TplVars{
		H1: "Go web!",
	}

	t.Execute(w, tplVars)
}
