package service

import (
	"github.com/ChenfengHub/golang-task/task04/entity"
	"github.com/ChenfengHub/golang-task/task04/store"
)

type PostService struct {
	postStore store.PostStore
}

func NewPostService(store store.PostStore) *PostService {
	return &PostService{postStore: store}
}

func (ps *PostService) GetPostList() []store.SimplePost {
	return ps.postStore.GetPostList()
}

func (ps *PostService) GetPostDetail(postID uint) entity.Post {
	return ps.postStore.GetPostDetail(postID)
}

func (ps *PostService) CreatePost(post *entity.Post) error {
	return ps.postStore.CreatePost(post)
}

func (ps *PostService) UpdatePost(userId uint, post *entity.Post) error {
	return ps.postStore.UpdatePost(userId, post)
}

func (ps *PostService) DeletePost(userId uint, postId uint) error {
	return ps.postStore.DeletePost(userId, postId)
}
