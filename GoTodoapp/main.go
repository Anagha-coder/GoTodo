package main

import (
	// always give org package name while importing
	"fmt"
	"log"
	"myapi/controller"
	"myapi/model"

	"net/http"

	_ "github.com/go-sql-driver/mysql" //mysql driver
)

func main() {

	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":5000", mux))

}

// mux := http.NewServeMux()
