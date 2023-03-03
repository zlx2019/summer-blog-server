/**
  @author: Zero
  @date: 2023/3/3 18:36:59
  @desc: 文章实体

**/

package models

import . "summer/models/custom"

// Article 文章表结构
type Article struct {
	BaseModel
	Title        string `json:"title" gorm:"size:32; not null; comment: 文章标题"`
	Desc         string `json:"desc" gorm:"size: 256; comment: 文章简介"`
	Content      string `json:"content" gorm:"comment: 文章内容"`
	LookCount    int    `json:"look_count" gorm:"default:0; comment: 浏览量"`
	CommentCount int    `json:"comment_count" gorm:"default:0; comment: 评论量"`
	LikeCount    int    `json:"like_count" gorm:"default:0; comment: 点赞量"`
	CollectCount int    `json:"collect_count" gorm:"default:0; comment: 收藏量"`
	Category     string `json:"category" gorm:"comment: 文章分类"`
	Source       string `json:"source" gorm:"comment: 文章来源"`
	Link         string `json:"link" gorm:"comment: 文章链接"`
	WordCount    int    `json:"word_count" gorm:"comment: 文章字数"`

	// 文章标签列表
	Tags []Tag `json:"tags" gorm:"many2many:article_tag"`
	// 文章评论列表
	Comments []Comment `json:"comments" gorm:"foreignKey:ArticleID"`
	// 所属用户
	User User `json:"user" gorm:"foreignKey:UserID"`
	// 所属用户ID
	UserID uint64 `json:"user_id"`
}
