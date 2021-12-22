package service

import (
	"github/johnraychina/study-go/mock/blog"
)

type Service interface {
	ListPosts() []blog.Post
}

type service struct {
	blog blog.Blog
}

func NewService(b blog.Blog) *service {
	return &service{
		blog: b,
	}
}

func (s *service) ListPosts() []blog.Post {
	return s.blog.ListPosts()
}
