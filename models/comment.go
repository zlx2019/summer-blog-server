/**
  @author: Zero
  @date: 2023/3/3 19:37:56
  @desc: 文章评论

**/

package models

import "summer/models/custom"

// Comment 文章评论结构
type Comment struct {
	custom.BaseModel

	// 该评论下的所有子级评论
	ChildComments []*Comment `json:"child_comments" gorm:"foreignKey:ParentCommentID"`
	// 该评论的父级评论
	ParentComment *Comment `json:"parent_comment" gorm:"foreignKey:ParentCommentID"`
	// 父级评论的ID
	ParentCommentID *uint64 `json:"parent_comment_id" gorm:"default:0; comment: 父级评论ID"`
	// 评论内容
	Content string `json:"content" gorm:"size:256; not null; comment: 评论内容"`
	// 点赞量
	LikeCount int `json:"like_count" gorm:"default:0; comment: 点赞量"`
	// 子级评论数量
	ChildCount int `json:"child_count" gorm:"default:0; comment: 子级评论数量"`
	// 评论者IP
	IP string `json:"ip" gorm:"size:20;comment: 评论者IP"`
	// 所在地
	Addr string `json:"addr" gorm:"size:20; comment: 所在地"`
	// 所关联的文章
	Article Article `json:"article" gorm:"foreignKey:ArticleID"`
	// 文章ID
	ArticleID uint64 `json:"article_id" gorm:"comment: 文章ID"`
	// 评论者 所关联的用户
	User User `json:"user"`
	// 评论者ID
	UserID uint64 `json:"user_id" gorm:"comment: 评论者ID"`
}
