package main

import (
	"GO-CRUD-NoSQL/controllers"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	//Creating new router using httprouter
	router := httprouter.New()

	uc := controllers.NewUserController(getSession())

	router.GET("/user/:id", uc.GetUser)
	router.POST("", uc.CreateUser)
	router.DELETE("", uc.DeleteUser)

	http.ListenAndServe("localhost:9000", router)
}

// Connecting to DB
// What does mgo.Session do here?
func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil {
		fmt.Println("Could Not Connect to DB")
	}
	fmt.Println("Connected to DB!😁")
	return session
}
