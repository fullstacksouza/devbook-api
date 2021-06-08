package main

import (
	"devbook-api/src/config"
	"devbook-api/src/router"
	"fmt"
	"log"
	"net/http"
)

func init() {
	// loads values from .env into the system
	config.Load()

}

func main() {
	fmt.Printf("Running api in %d\n", config.ApiPort)
	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApiPort), r))
}
