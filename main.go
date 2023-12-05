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
	router.POST("/user", uc.CreateUser)
	router.DELETE("/user/:id", uc.DeleteUser)

	http.ListenAndServe("localhost:9000", router)
}

// Connecting to DB
// What does mgo.Session do here?
// Dial returns a *mgo.Session which is stored in "session" and returned
func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB!😁")
	return session
}
