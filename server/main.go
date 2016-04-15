package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var tmpl *template.Template

type someData struct {
	Name   string
	People []string
	Male   bool
}

type person struct {
	FirstName string
	LastName  string
}

// DogHandler to handle mux
type DogHandler int

// IndexHandler to handle mux
type IndexHandler int

// FormHandler to handle mux
type FormHandler int

// FormUploadHandler to handle mux
type FormUploadHandler int

// CookieHandler to handle mux
type CookieHandler int

func (h DogHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">
		<h1>`+req.URL.Path+`</h1>`)
	io.WriteString(res, `<h2>`+req.URL.Query().Get("q")+`</h2>`)
}

func (h IndexHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	data := someData{
		Name:   "Harsh",
		People: []string{"Gandhi", "Buddha", "MK"},
		Male:   true,
	}
	tmpl.ExecuteTemplate(res, "hello.gohtml", data)
	// io.WriteString(res, `<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
}

func (h FormHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fName := req.FormValue("first")
	lName := req.FormValue("last")

	err := tmpl.ExecuteTemplate(res, "form.gohtml", person{fName, lName})
	if err != nil {
		http.Error(res, err.Error(), 500)
		log.Fatalln(err)
	}
}

func (h FormUploadHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		fmt.Println(os.TempDir())
		src, _, err := req.FormFile("file") // get the uploaded file
		if err != nil {
			panic(err)
		}
		defer src.Close()

		io.LimitReader(src, 500) // Limit the file content if the file size is too big

		dst, err := os.Create(filepath.Join("./", "file.txt")) // destination file create
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}
		defer dst.Close()

		io.Copy(dst, src) // copy the contents
	}
	tmpl.ExecuteTemplate(res, "formupload.gohtml", nil)
}

func (h CookieHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("counter") // get a cookie counter to keep track of previous count
	if err == http.ErrNoCookie {
		// Since for the first time cookie wont be present, we set it to 0 value and return.
		c = &http.Cookie{
			Name:  "counter",
			Value: "0",
		}
	}

	val, _ := strconv.Atoi(c.Value) // convert the to int
	val++                           // increment the value
	c.Value = strconv.Itoa(val)     // conver the value to string

	// Set a cookie
	// http.SetCookie(res, &http.Cookie{
	// 	Name:  "my-cookie",
	// 	Value: "Some Value",
	// })

	// Set the counter cookie
	http.SetCookie(res, c)
}

// Always load templates in init
func init() {
	// var err error
	// tmpl, err = template.ParseFiles("templates/hello.gohtml")
	// if err != nil {
	// 	panic(err)
	// }
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

func main() {

	var dog DogHandler
	var index IndexHandler
	var form FormHandler
	var formUpload FormUploadHandler
	var cookie CookieHandler

	mux := http.NewServeMux()
	// mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/dog", dog)
	mux.Handle("/", index)
	mux.Handle("/form", form)
	mux.Handle("/form-upload", formUpload)
	mux.Handle("/cookie", cookie)
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
