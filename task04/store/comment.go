package store

import (
	"errors"

	"github.com/ChenfengHub/golang-task/task04/entity"
	"gorm.io/gorm"
)

type CommentStore interface {
	AddComment(comment *entity.Comment) error
	GetCommentList(postId uint) []entity.Comment
}

type commentStore struct {
	db *gorm.DB
}

func NewCommentStore(db *gorm.DB) CommentStore {
	return &commentStore{db: db}
}

func (cs *commentStore) AddComment(comment *entity.Comment) error {
	post := entity.Post{}
	cs.db.Where("id = ?", comment.PostID).First(&post)
	if post.ID == 0 {
		return errors.New("评论的文章不存在")
	}

	if err := cs.db.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func (cs *commentStore) GetCommentList(postId uint) []entity.Comment {
	comments := []entity.Comment{}
	cs.db.Where("post_id = ?", postId).Order("updated_at desc, created_at desc, id desc").Find(&comments)

	return comments
}
