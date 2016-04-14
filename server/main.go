package main

import (
	"html/template"
	"io"
	"net/http"
)

var tmpl *template.Template

type someData struct {
	Name   string
	People []string
	Male   bool
}

// DogHandler to handle mux
type DogHandler int

// IndexHandler to handle mux
type IndexHandler int

func (h DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
		<h1>`+req.URL.Path+`</h1>`)
}

func (h IndexHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := someData{
		Name:   "Harsh",
		People: []string{"Gandhi", "Buddha", "MK"},
		Male:   true,
	}
	tmpl.Execute(res, data)
	// io.WriteString(res, `<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
}

// Always load templates in init
func init() {
	// var err error
	// tmpl, err = template.ParseFiles("templates/hello.gohtml")
	// if err != nil {
	// 	panic(err)
	// }
	tmpl = template.Must(template.ParseFiles("templates/hello.gohtml"))

}

func main() {

	var dog DogHandler
	var index IndexHandler
	mux := http.NewServeMux()
	mux.Handle("/dog", dog)
	mux.Handle("/", index)
	// http.HandleFunc("/", index)

	// serve css and images
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", mux)
}

// func index(res http.ResponseWriter, req *http.Request) {
// 	data := someData{
// 		Name:   "Harsh",
// 		People: []string{"Gandhi", "Buddha", "MK"},
// 		Male:   true,
// 	}
// 	tmpl.Execute(res, data)
// 	// _, err := io.WriteString(res, "Hellssdssso")
// 	// if err != nil {
// 	// 	log.Fatal(err) // Prints to terminal and exits
// 	// 	panic(err)
// 	// }
// 	// fmt.Fprint(res, "a ...interface{}")
// 	// res.Write([]byte("Hello"))
//
// }
