package entity

import "gorm.io/gorm"

type Log struct {
	gorm.Model
	ErrMsg     string
	UserId     uint   `gorm:"commnet:导致日志报错的操作者id"`
	Method     string `gorm:"commnet:调用接口方法:POST、GET等"`
	Path       string `gorm:"commnet:调用url信息"`
	StatusCode int    `gorm:"commnet:状态码"`
	ClientIP   string `gorm:"commnet:客户端ip"`
	UserAgent  string `gorm:"commnet:agent信息"`
	RequestID  string `gorm:"commnet:RequestID信息"`
}
