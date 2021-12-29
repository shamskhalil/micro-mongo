package routes

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	controller "github.com/shamskhalil/micro-mongo/controllers"
	model "github.com/shamskhalil/micro-mongo/models"
	"gopkg.in/mgo.v2/bson"
)

type UserRoute struct {
	UserCtrl *controller.UserCtrl
}

type CustomResponse struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

func (ur *UserRoute) GetUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	id := param.ByName("id")
	user, err := ur.UserCtrl.GetUser(id)
	if err != nil {
		resp := CustomResponse{Message: err.Error(), Description: "Error finding user by id"}
		json.NewEncoder(w).Encode(resp)
	}
	json.NewEncoder(w).Encode(user)
}
func (ur *UserRoute) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")

	users, err := ur.UserCtrl.GetUsers()
	if err != nil {
		resp := CustomResponse{Message: err.Error(), Description: "Error fetching all users"}
		json.NewEncoder(w).Encode(resp)
	}
	json.NewEncoder(w).Encode(users)
}
func (ur *UserRoute) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		resp := CustomResponse{Message: err.Error(), Description: "Error Decoding request body"}
		json.NewEncoder(w).Encode(resp)
	}
	userId, err2 := ur.UserCtrl.CreateUser(&user)
	if err2 != nil {
		resp := CustomResponse{Message: err2.Error(), Description: "Error Saving User to Database"}
		json.NewEncoder(w).Encode(resp)
	}
	user.Id = bson.ObjectIdHex(userId)
	json.NewEncoder(w).Encode(user)
}
func (ur *UserRoute) UpdateUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	user := model.User{}
	id := param.ByName("id")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		resp := CustomResponse{Message: err.Error(), Description: "Error Decoding request body"}
		json.NewEncoder(w).Encode(resp)
	}
	err = ur.UserCtrl.UpdateUser(&user)
	if err != nil {
		resp := CustomResponse{Message: err.Error(), Description: "Error Updating User"}
		json.NewEncoder(w).Encode(resp)
	}
	obj := CustomResponse{Message: "User id:" + id, Description: "User has been updated successfully!"}
	json.NewEncoder(w).Encode(obj)
}

func (ur *UserRoute) DeleteUser(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	id := param.ByName("id")
	err := ur.UserCtrl.DeleteUser(id)
	if err != nil {
		resp := CustomResponse{Message: err.Error(), Description: "Error Deleting User"}
		json.NewEncoder(w).Encode(resp)
	}
	obj := CustomResponse{Message: "User id:" + id, Description: "User has been deleted successfully!"}
	json.NewEncoder(w).Encode(obj)
}
