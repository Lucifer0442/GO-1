package main // "This is the main part of my program." Every Go program starts here.

import "fmt" // "fmt" stands for format. It is a package that contains functions for formatting text, including printing to the console.
import "net/http" // "net/http" is a package that provides HTTP client and server implementations. It allows you to make HTTP requests and handle HTTP responses.

func main (){ // is the engine of the program
	
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w,"Hello World!")
	})

	http.ListenAndServe(":8080",nil)
}
