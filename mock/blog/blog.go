package blog

type Post struct{}

type Blog interface {
	ListPosts() []Post
}

type jekyll struct{}

func (b *jekyll) ListPosts() []Post {
	return []Post{}
}

type wordpress struct{}

func (b *wordpress) ListPosts() []Post {
	return []Post{}
}
