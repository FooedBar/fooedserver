package main

import (
	"fmt"
	"github.com/FooedBar/fooedserver/controllers"
	"github.com/FooedBar/fooedserver/models"
	"log"
	"net/http"
)

func main() {
	err := models.Setup()
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port", models.Conf.Port[1:])
	err = http.ListenAndServe(models.Conf.Port, controllers.CreateRouter())
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
