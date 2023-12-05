package controllers

import (
	"GO-CRUD-NoSQL/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

// NEW USER CONTROLLER
func NewUserController(session *mgo.Session) *UserController {
	return &UserController{session}
}

// GET USER BY ID HANDLER
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	//Checking if the id is Hex
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	//oid := bson.ObjectHex(id) DOESN'T EXIST
	oid := bson.ObjectIdHex(id)

	user := models.User{}

	//Selecting database from session [here uc struct has the session]
	connection := uc.session.DB("Mongo-Golang-CRUD")
	//Selecting "Collection" from the database
	collection := connection.C("Users")
	//Looking for my required data
	requiredData := collection.FindId(oid)
	//Populating the "user" struct with the required data
	err := requiredData.One(&user)

	// err := uc.session.DB("Mongo-Golang-CRUD").C("Users").FindId(oid).One(&user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//Converting the gobson to json
	user_json, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error while Parsing inside GetUser!")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s/n", user_json)
	w.Write(user_json)
}
