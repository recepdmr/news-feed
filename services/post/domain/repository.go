package domain

type PostRepository interface {
	GetAll() []Post
	GetById(id string) (*Post, error)
	Insert(Post) error
	Delete(Post) error
}
