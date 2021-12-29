package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	controller "github.com/shamskhalil/micro-mongo/controllers"
	"github.com/shamskhalil/micro-mongo/routes"
)

func main() {
	userCtrl := controller.NewUserCtrl("localhost", 27017)
	myroutes := routes.UserRoute{UserCtrl: userCtrl}

	r := httprouter.New()
	r.GET("/user", myroutes.GetUsers)
	r.GET("/user/:id", myroutes.GetUser)
	r.POST("/user", myroutes.CreateUser)

	fmt.Println("Server listening on port 3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Error starting server !! ", err)
	}

}
