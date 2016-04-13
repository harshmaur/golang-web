package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {
	// name := "Harsh Maur"
	name := os.Args[1] // to get the command line arguments

	// Sprint to add the things to the variable
	// str := fmt.Sprint(`
	//     <!DOCTYPE html>
	//     <html>
	//         <head>
	//             <meta charset="utf-8">
	//             <title>Hello</title>
	//         </head>
	//         <body>
	//             <h1>` + name + `</h1>
	//         </body>
	//     </html>
	//     `)

	// Parse a predefined file.
	tmpl, err := template.ParseFiles("template.gohtml") // Parses files that are mentioned
	fmt.Printf("%T\n", tmpl)
	if err != nil {
		log.Fatal(err)
	}

	nf, err := os.Create("index.html") // create a new file and get pointer to that file
	// fmt.Printf("%T\n", *nf)
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close() // defer close the file

	err = tmpl.Execute(nf, name) // execute the template and pass in the data
	if err != nil {
		panic(err)
	}

	// io.Copy(nf, strings.NewReader(str)) // copy the contents of str, get a reader interface and paste to nf
}
