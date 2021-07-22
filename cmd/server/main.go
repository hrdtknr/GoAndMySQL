package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("start")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "hello, docker container")
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
