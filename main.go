package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	log.Print("HTTPDump listening at :8080")

	http.HandleFunc("/", Dump)
	http.ListenAndServe(":8080", nil)
}

func Dump(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	log.Printf("%s", dump)
	fmt.Fprint(w, "Hello from HTTPDump")
}
