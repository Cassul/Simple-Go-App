package mongo

import (
	"github.com/cassul/root"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type userModel struct {
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Email    string        `json:"email"`
	Password string        `json:"password"`
}

func userModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

func newUserModel(u *root.User) *userModel {
	return &userModel{
		Email:    u.Email,
		Password: u.Password}
}

func (u *userModel) toRootUser() *root.User {
	return &root.User{
		Id:       u.Id.Hex(),
		Email:    u.Email,
		Password: u.Password}
}
