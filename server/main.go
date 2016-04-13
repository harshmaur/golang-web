package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

// Always load templates in init
func init() {
	var err error
	tmpl, err = template.ParseFiles("templates/hello.gohtml")
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", index)

	// serve css and images
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {

	tmpl.Execute(res, nil)
	// _, err := io.WriteString(res, "Hellssdssso")
	// if err != nil {
	// 	log.Fatal(err) // Prints to terminal and exits
	// 	panic(err)
	// }
	// fmt.Fprint(res, "a ...interface{}")
	// res.Write([]byte("Hello"))

}
