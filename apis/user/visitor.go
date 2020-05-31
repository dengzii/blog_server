package user

import (
	"github.com/dengzii/blog_server/models/base"
)

type Visitor struct {
	base.CommonModel
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
	Location  string `json:"location"`
	OS        string `json:"is_mobile"`
}

func AddVisitor(visitor Visitor) {

}
