package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		hostname, err_hostname := os.Hostname()
		if err_hostname != nil {
			fmt.Println(err_hostname)
		}
		fmt.Fprintf(w, "Hello, World on %s!", hostname)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
