package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(response http.ResponseWriter, req *http.Request) {
	response.Write([]byte("Hello World"))
}

func main() {
	fmt.Println("Hello World!")
	http.HandleFunc("/hello", helloWorld)
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}
