package main

import (
	"fmt"
	"log"
	"net/http"
	helpers "sendemail/Helpers"
	router "sendemail/Router"
	"time"
)

func main() {
	fmt.Println("Server is Starting...")
	go func() {
		log.Fatal(http.ListenAndServe(":8000", router.Routers()))
	}()
	for {
		fmt.Println("Ticker ticked at", time.Now())
		helpers.DeleteVcode()
		time.Sleep(10 * time.Minute)
	}

}
