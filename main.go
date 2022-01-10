package main

import (
	"html/template"
	"net/http"
)

type User struct {
	UserName string
	Password string
	Success  bool
	Storege  string
}

func main() {
	handlerRequest()
}

func handlerRequest() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/static", http.NotFoundHandler())
	//http.Handle("/assets/", http.StripPrefix("/assets", fs))
	mux.Handle("/assets/", http.StripPrefix("/assets", fs))
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/contacts/", contactPage)
	mux.HandleFunc("/auth/", auth)
	//http.HandleFunc("/auth/", auth)

	http.ListenAndServe(":8000", mux)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, error := template.ParseFiles("webView/templates/firstPage.html")
	if error == nil {
		tmpl.Execute(w, "it works")
	}
}
func contactPage(w http.ResponseWriter, r *http.Request) {
	tmpl, error := template.ParseFiles("webView/templates/secondPage.html")
	if error == nil {
		tmpl.Execute(w, "")
	}
}

func auth(w http.ResponseWriter, r *http.Request) {
	tmpl, error := template.ParseFiles("webView/templates/auth.html")
	if error != nil {
		return
	}

	data := User{
		UserName: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	if data.Password == "qwe" {
		data.Success = true
	} else {
		data.Success = false
	}

	data.Storege = "Welcome to my html" + data.UserName

	tmpl.Execute(w, data)
}
