package controller

import (
	"fmt"
	"log"

	model "github.com/shamskhalil/micro-mongo/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserCtrl struct {
	Session *mgo.Session
}

func NewUserCtrl(host string, port int) *UserCtrl {
	url := fmt.Sprintf("mongodb://%s:%d", host, port)
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal("Error connecting to mogo database", err)
	}
	return &UserCtrl{Session: session}
}

func (u *UserCtrl) CreateUser(user *model.User) (string, error) {
	user.Id = bson.NewObjectId()
	err := u.Session.DB("usersdb").C("users").Insert(user)
	if err != nil {
		fmt.Println("Error inserting new user ", err)
		return "", err
	}
	fmt.Println("User inserted successfully!")
	return user.Id.Hex(), nil
}

func (u *UserCtrl) GetUser(id string) (*model.User, error) {
	user := model.User{}
	err := u.Session.DB("usersdb").C("users").FindId(bson.ObjectIdHex(id)).One(&user)
	if err != nil {
		fmt.Println("Error getting user by id ", err)
		return &user, err
	}
	fmt.Println("Sigle User found !")
	return &user, nil
}

func (u *UserCtrl) GetUsers() ([]model.User, error) {
	var users []model.User
	err := u.Session.DB("usersdb").C("users").Find(bson.M{}).All(&users)
	if err != nil {
		fmt.Println("Error getting all users ", err)
		return users, err
	}
	fmt.Println("All Users found !")
	return users, nil
}

func (u *UserCtrl) UpdateUser(user *model.User) error {
	err := u.Session.DB("usersdb").C("users").UpdateId(user.Id, user)
	if err != nil {
		fmt.Println("Error updating user ", err)
		return err
	}
	fmt.Println("User updated successfully!")
	return nil
}

func (u *UserCtrl) DeleteUser(id string) error {
	err := u.Session.DB("usersdb").C("users").RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		fmt.Println("Error deleting user by id ", err)
		return err
	}
	fmt.Println("User deleted !")
	return nil
}
