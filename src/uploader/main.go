package main

import (
	"log"
	"net/http"
	"os"
)

var BasePath string

func checkError(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func main() {
	log.Println(os.Args)
	if len(os.Args) > 1 {
		log.Println("Base directory spacified:", os.Args[1])
		BasePath = os.Args[1]
	} else {
		log.Println("No Path found use '.'")
		BasePath = "."
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", upload) //the default action is upload

	log.Println("Listen on 6666")
	err := http.ListenAndServe(":6666", mux)
	checkError(err)
}
