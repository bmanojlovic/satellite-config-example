package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		hostname, err = os.Hostname()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, "Hello, World on %s!", hostname)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
