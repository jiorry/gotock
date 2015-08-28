package user

import (
	"github.com/kere/gos/db"
)

type UserModel struct {
	db.BaseModel
}

func NewUserModel() *UserModel {
	m := &UserModel{}
	m.Init(&UserVO{})
	return m
}
