package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var static string

	flag.StringVar(&static, "static", ".", "the directory to serve static files from.")
	flag.Parse()

	// Thanks to https://blog.questionable.services/article/vue-react-ember-server-golang/
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(static+"/dist/"))))
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe("localhost:3000", nil),
	)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
