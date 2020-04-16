package user

import (
	"github.com/dengzii/blog_server/models/base"
)

const (
	del           = 1 >> 8
	update        = 1 >> 7
	read          = 1 >> 6
	manageArticle = 1 >> 5
	manageUser    = 1 >> 4
	manageSite    = 1 >> 3
)

type Role struct {
	base.CommonModel
	Name        string
	Description string
	Permission  int
	Parent      *Role
}
