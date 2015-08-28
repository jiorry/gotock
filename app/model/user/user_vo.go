package user

import (
	"time"

	"github.com/kere/gos/db"
)

type UserVO struct {
	db.BaseVO
	Id      int64     `json:"id" skip:"all"`
	Nick    string    `json:"nick"`
	Token   string    `json:"token"`
	Salt    string    `json:"salt"`
	Status  int       `json:"status"`
	Created time.Time `json:"created_at"`
	LastSee time.Time `json:"last_see_at"`
}

func NewUserVO(nick string) *UserVO {
	vo := &UserVO{
		Nick:   nick,
		Status: 0,
	}

	vo.Init(vo)
	return vo
}

func (a *UserVO) Table() string {
	return "users"
}
