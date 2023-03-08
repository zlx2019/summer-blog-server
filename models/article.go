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
	Desc         string `json:"desc" gorm:"size:256; not null; comment: 文章简介"`
	Content      string `json:"content" gorm:"type:text;comment: 文章内容"`
	LookCount    int    `json:"look_count" gorm:"default:0; comment: 浏览量"`
	CommentCount int    `json:"comment_count" gorm:"default:0; comment: 评论量"`
	LikeCount    int    `json:"like_count" gorm:"default:0; comment: 点赞量"`
	CollectCount int    `json:"collect_count" gorm:"default:0; comment: 收藏量"`
	Category     string `json:"category" gorm:"comment: 文章分类"`
	Source       string `json:"source" gorm:"comment: 文章来源"`
	Link         string `json:"link" gorm:"comment: 文章链接"`
	WordCount    int    `json:"word_count" gorm:"comment: 文章字数"`
	Author       string `json:"author" gorm:"size:42;comment: 作者昵称"`
	FileUrl      string `json:"file_url" gorm:"comment: 封面图片url"`

	// 封面文件
	File File `json:"file" gorm:"foreignKey:FileID"`
	// 封面文件ID
	FileID uint64 `json:"file_id" gorm:"comment: 封面图片ID"`
	// 所属用户
	User User `json:"user" gorm:"foreignKey:UserID"`
	// 所属用户ID
	UserID uint64 `json:"user_id" gorm:"comment: 所属用户ID"`
	// 文章标签列表;指定自定义的关系表名,和两个关联ID
	Tags []Tag `json:"tags" gorm:"many2many:article_tag;joinForeignKey:ArticleID;JoinReferences:TagID"`
	// 文章评论列表
	Comments []Comment `json:"comments" gorm:"foreignKey:ArticleID"`
}
