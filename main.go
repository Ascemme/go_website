package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
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
	file, handler, err := r.FormFile("myfile")
	if error == nil {
		tmpl.Execute(w, "")
	}
	if err != nil {
		print(err)
		print("error")
		return
	}
	defer file.Close()

	dst, err := os.Create(handler.Filename)
	defer dst.Close()
	if err != nil {
		return
	}
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func auth(w http.ResponseWriter, r *http.Request) {
	newdata := &User{}
	tmpl, error := template.ParseFiles("webView/templates/auth.html")
	if error != nil {
		return
	}
	data := *newdata
	data = User{
		UserName: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	print(data.Success)
	if data.Password == "qwe" {
		data.Success = true
	} else {
		data.Success = false
	}
	data.Storege = "Welcome to my html " + data.UserName
	tmpl.Execute(w, data)
}
