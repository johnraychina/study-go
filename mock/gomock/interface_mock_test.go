package gomock

import (
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github/johnraychina/study-go/mock/blog"
	"github/johnraychina/study-go/mock/service"
	"testing"
)

// see https://github.com/golang/mock

func TestListPosts(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBlog := NewMockBlog(ctrl)
	mockBlog.EXPECT().ListPosts().Return([]blog.Post{})

	s := service.NewService(mockBlog)
	assert.Equal(t, []blog.Post{}, s.ListPosts())
}
