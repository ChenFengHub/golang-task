package hv1

import (
	"net/http"
	"strconv"

	"github.com/ChenfengHub/golang-task/task04/entity"
	"github.com/ChenfengHub/golang-task/task04/service"
	"github.com/ChenfengHub/golang-task/task04/store"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentHandler struct {
	commentService *service.CommentService
}

type CommentResp struct {
	Id      uint
	Content string
	UserId  uint
	PostId  uint
}

func commentToCommentResp(comment entity.Comment) CommentResp {
	resp := CommentResp{}

	resp.Id = comment.ID
	resp.Content = comment.Content
	resp.UserId = comment.UserID
	resp.PostId = comment.PostID

	return resp
}

func commentListToRespList(comments []entity.Comment) []CommentResp {
	if len(comments) == 0 {
		return []CommentResp{}
	}
	resp := make([]CommentResp, 0)
	for _, comment := range comments {
		resp = append(resp, commentToCommentResp(comment))
	}
	return resp
}

func newCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

func SetupCommentRoutes(r *gin.Engine, db *gorm.DB) {
	cs := service.NewCommentService(store.NewCommentStore(db))
	ch := newCommentHandler(cs)

	api := r.Group("/v1")

	commentGroup := api.Group("/comment")
	// 代码块中批量操作
	{
		commentGroup.POST("/add", ch.AddComment)
		commentGroup.POST("/getList", ch.GetCommentList)
	}
}

func (cs *CommentHandler) AddComment(c *gin.Context) {
	var comment entity.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, _ := strconv.ParseUint(c.GetHeader("Userid"), 10, 64)
	comment.UserID = uint(userId)

	if err := cs.commentService.AddComment(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": commentToCommentResp(comment)})
}

func (cs *CommentHandler) GetCommentList(c *gin.Context) {
	var comment entity.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comments := cs.commentService.GetCommentList(comment.PostID)
	c.JSON(http.StatusCreated, gin.H{"data": commentListToRespList(comments)})
}
