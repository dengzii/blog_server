package friend

import "server/db"

type Friend struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Url    string `json:"url"`
	Email  string `json:"emil"`
	Avatar string `json:"avtar"`
	Alt    string `json:"alt"`
}

func AddFriend(f Friend) *Friend {

	db.Insert(f)
	return &f
}

func GetFriend(f Friend) *Friend {
	return nil
}
