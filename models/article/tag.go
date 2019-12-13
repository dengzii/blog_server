package article

import "server/db"

type Tag struct {
	Id           int
	Belong       Class
	Name         string
	ArticleCount int
	Display      bool
}

func AddTag() (err error) {

	db.Insert(&Tag{})
	return err
}

func GetTag() (tag *Tag, err error) {

	return nil, nil
}
