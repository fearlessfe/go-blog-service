package model

import "github.com/fearlessfe/go-blog-service/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

// tag.go
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

func (t Tag) TableName() string {
	return "blog_tag"
}
