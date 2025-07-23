package service

import (
	"github.com/ChenfengHub/golang-task/task04/entity"
	"github.com/ChenfengHub/golang-task/task04/store"
)

type CommentService struct {
	commentStore store.CommentStore
}

func NewCommentService(commentStore store.CommentStore) *CommentService {
	return &CommentService{commentStore: commentStore}
}

func (cs *CommentService) AddComment(comment *entity.Comment) error {
	return cs.commentStore.AddComment(comment)
}

func (cs *CommentService) GetCommentList(postId uint) []entity.Comment {
	return cs.commentStore.GetCommentList(postId)
}
