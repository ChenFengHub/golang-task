package store

import (
	"errors"

	"github.com/ChenfengHub/golang-task/task04/entity"
	"gorm.io/gorm"
)

type PostStore interface {
	GetPostList() []SimplePost
	GetPostDetail(postID uint) entity.Post
	CreatePost(post *entity.Post) error
	// 更新文章。userId是更新操作者的id
	UpdatePost(userId uint, post *entity.Post) error
	// 删除文章。userId是删除操作者的id
	DeletePost(userId uint, postId uint) error
}

type SimplePost struct {
	Id    uint
	Title string
}

type postStore struct {
	db *gorm.DB
}

func NewPostStore(db *gorm.DB) PostStore {
	return &postStore{db: db}
}

func (ps *postStore) GetPostList() []SimplePost {
	posts := []entity.Post{}
	ps.db.Model(&entity.Post{}).Where("1=1").Find(&posts)

	res := make([]SimplePost, len(posts))
	if len(posts) == 0 {
		return res
	}

	for i, v := range posts {
		res[i].Id = v.ID
		res[i].Title = v.Title
	}

	return res
}

func (ps *postStore) GetPostDetail(postID uint) entity.Post {
	post := entity.Post{}
	ps.db.Select("id", "title", "content", "user_id", "created_at").Where("id = ?", postID).First(&post)
	return post
}

func (ps *postStore) CreatePost(post *entity.Post) error {
	if err := ps.db.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

func (ps *postStore) UpdatePost(userId uint, post *entity.Post) error {
	if userId == 0 || post.ID == 0 {
		return errors.New("userId or post.ID is miss")
	}

	exist := entity.Post{}
	ps.db.Where("id = ?", post.ID).First(&exist)
	if exist.ID == 0 {
		return errors.New("post is not exist")
	}
	if exist.UserID != userId {
		return errors.New("只有文章的作者才能更新自己的文章")
	}

	if err := ps.db.Where("id = ?", post.ID).Updates(entity.Post{Title: post.Title, Content: post.Content}).Error; err != nil {
		return err
	}
	return nil
}

func (ps *postStore) DeletePost(userId uint, postId uint) error {
	if userId == 0 || postId == 0 {
		return errors.New("userId or postId is miss")
	}

	exist := entity.Post{}
	ps.db.Where("id = ?", postId).First(&exist)
	if exist.ID == 0 {
		return errors.New("post is not exist")
	}
	if exist.UserID != userId {
		return errors.New("只有文章的作者才能删除自己的文章")
	}

	if err := ps.db.Where("id = ?", postId).Delete(&entity.Post{}).Error; err != nil {
		return err
	}

	return nil
}
