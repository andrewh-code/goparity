package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "welcome to goparity server");

}

func main() {
	http.HandleFunc("/", handler)
	port := 9000;
	fmt.Printf("now servering on localhost: %d\n", port)
	if err := http.ListenAndServe(":9000", nil); err != nil {
	  panic(err)
	}
}    