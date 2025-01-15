package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rushabh2390/gousersmodule/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterUserstore(r)
	http.Handle("/", r)
	fmt.Println("Starting server at 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal((err))
	}
}
