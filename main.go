package main

import (
	"fmt"
	"log"
	"net/http"
	router "sendemail/Router"
)

func main() {
	fmt.Println("Server is Starting...")
	log.Fatal(http.ListenAndServe(":8000", router.Routers()))
	fmt.Println("Server has Started")
}
