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

type PostHandler struct {
	postService *service.PostService
}

type PostResp struct {
	ID      uint
	Title   string
	Content string
	UserId  uint
}

func postToPostResp(post entity.Post) PostResp {
	resp := PostResp{}

	resp.ID = post.ID
	resp.Title = post.Title
	resp.Content = post.Content
	resp.UserId = post.UserID

	return resp
}

// func postListToPostRespList(posts []entity.Post) []PostResp {
// 	resps := []PostResp{}

// 	if len(posts) == 0 {
// 		return resps
// 	}

// 	for _, post := range posts {
// 		resps = append(resps, postToPostResp(post))
// 	}

// 	return resps
// }

func newPostHandler(ps *service.PostService) *PostHandler {
	return &PostHandler{postService: ps}
}

func SetupPostRoutes(r *gin.Engine, db *gorm.DB) {
	ps := service.NewPostService(store.NewPostStore(db))
	ph := newPostHandler(ps)

	api := r.Group("/v1")

	postGroup := api.Group("/post")
	// 代码块中批量操作
	{
		postGroup.POST("/getList", ph.GetPostList)
		postGroup.POST("/getDetail", ph.GetPostDetail)
		postGroup.POST("/create", ph.CreatePost)
		postGroup.POST("/update", ph.UpdatePost)
		postGroup.POST("/delete", ph.DeletePost)
	}
}

func (ph *PostHandler) GetPostList(c *gin.Context) {
	data := ph.postService.GetPostList()
	c.JSON(http.StatusCreated, gin.H{"data": data})
}

func (ph *PostHandler) GetPostDetail(c *gin.Context) {
	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postID := post.ID
	detail := ph.postService.GetPostDetail(postID)
	if detail.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "详情不存在"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": postToPostResp(detail)})
}

func (ph *PostHandler) CreatePost(c *gin.Context) {
	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, _ := strconv.ParseUint(c.GetHeader("Userid"), 10, 64)
	post.UserID = uint(userId)
	if err := ph.postService.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"data": postToPostResp(post)})
}

func (ph *PostHandler) UpdatePost(c *gin.Context) {
	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, _ := strconv.ParseUint(c.GetHeader("Userid"), 10, 64)
	if err := ph.postService.UpdatePost(uint(userId), &post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": postToPostResp(post)})
}

func (ph *PostHandler) DeletePost(c *gin.Context) {
	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, _ := strconv.ParseUint(c.GetHeader("Userid"), 10, 64)
	postId := post.ID

	if err := ph.postService.DeletePost(uint(userId), postId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"data":    postToPostResp(post),
		"message": "删除完成",
	})
}
