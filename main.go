package main

import (
	"log"
	"net/http"

	"helloworld/helloworld"
)

func main() {
	helloServeMux := http.ServeMux{}
	helloServeMux.HandleFunc("/hello", helloworld.HelloWorld)
	log.Fatal(http.ListenAndServe(":3000", &helloServeMux))
}
