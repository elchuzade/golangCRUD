package main

import (
	"fmt"
	"log"
	"net/http"

	"cars/router"
)

func main() {
	fmt.Println("Running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router.Router()))
}
