package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func web(w http.ResponseWriter, r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Fprintf(w, "Error encountered: %v\n", err)
		return
	}
	fmt.Fprintf(w, "Below is the requestDump that the PCF app saw:\n\n\n%v\n", string(requestDump))
	fmt.Printf("%v\n", string(requestDump))
}

func main() {
	http.HandleFunc("/", web)
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
