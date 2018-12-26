package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func debugHandler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	log.Printf("%s", dump)
	fmt.Fprintf(w, "%s", dump)
}
func main() {
	http.HandleFunc("/", debugHandler)
	log.Println("server starting at 6969")
	log.Fatal(http.ListenAndServe(":6969", nil))
}
