package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"<p> Hello World </p>")
}

func main(){
	port := ":8080"
	http.Handle("/",http.FileServer(http.Dir("./static")))
	fmt.Printf("Listening on port %s\n",port)
	http.ListenAndServe(port,nil)
}
