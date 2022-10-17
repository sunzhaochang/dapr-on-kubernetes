package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", echo)
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	log.Println("recv: ", r.Form.Get("value"))

	fmt.Fprintf(w, "echo: %s", "this is server")
}
