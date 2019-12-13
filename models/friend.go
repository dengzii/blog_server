package models

import "server/db"

type Friends struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Url    string `json:"url"`
	Email  string `json:"emil"`
	Avatar string `json:"avtar"`
	Alt    string `json:"alt"`
}

func AddFriend(f Friends) *Friends {

	db.Insert(f)
	return &f
}
