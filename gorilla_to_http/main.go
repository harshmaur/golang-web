package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/harshmaur/golang-web/gorilla_to_http/utils"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("assets/templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)

	// 404 for favicon
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// serve assets
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	ck, err := req.Cookie("session-id") // get cookie from session
	if err != nil {                     // otherwise set a new one
		ck = utils.NewVisitor()
		http.SetCookie(res, ck)
	}

	if utils.Tampered(ck.Value) { // if tampered set a new one
		ck = utils.NewVisitor()
		http.SetCookie(res, ck)
	}

	if req.Method == "POST" {
		src, _, err := req.FormFile("data") // get image
		if err != nil {
			fmt.Println("err", err)
		}
		defer src.Close()

		fname := utils.UploadImage(src)       // upload and save image
		ck = utils.AddCookie(ck.Value, fname) // add new image data to exisiting cookie
		http.SetCookie(res, ck)               //set cookie

	}
	// set header to text/html
	res.Header().Set("Content-Type", "text/html; charset=utf-8")

	m := utils.GetModel(ck.Value)
	tmpl.ExecuteTemplate(res, "index.html", m)

}
