/**
  @author: Zero
  @date: 2023/3/3 16:00:32
  @desc: 用户实体

**/

package models

import . "summer/models/custom"

// User 用户表结构
type User struct {
	BaseModel
	NickName string         `json:"nickName" gorm:"type: varchar(36); comment: 用户昵称"`
	UserName string         `json:"userName" gorm:"type: varchar(36); uniqueIndex ;comment: 用户名"`
	Password string         `json:"password" gorm:"type: varchar(128); not null; comment: 用户密码"`
	Avatar   string         `json:"avatar" gorm:"type: varchar(256); comment: 用户头像"`
	Email    string         `json:"email" gorm:"type: varchar(128); comment: 用户邮箱"`
	Phone    string         `json:"phone" gorm:"size: 18; comment: 手机号码"`
	Addr     string         `json:"addr" gorm:"size:64;comment: 用户地址"`
	OtherId  string         `json:"otherId" gorm:"size: 64; comment: 其他认证平台唯一ID"`
	IP       string         `json:"Ip" gorm:"size: 20; comment: 用户所在IP"`
	Role     Role           `json:"role" gorm:"size:4; default:1; comment: 用户角色(1: 管理员 2:普通用户 3:游客)"`
	Source   RegisterManner `json:"source" gorm:"comment: 注册来源(1: QQ注册 2: 码云注册 3:github注册 4:邮箱注册)"`
	// 用户发布的所有文章
	// 一对多关系映射
	// 用户(一) --- 文章(多)
	// foreignKey 指定user表的Id,在articles表中的名字.
	Articles []Article `json:"articles" gorm:"foreignKey: UserID"`
	// 用户收藏的文章列表
	// 多对多关系映射
	// many2many: user_like_article 关系表的名字
	// joinForeignKey: UserID 关系表中User表的主键
	// JoinReferences: ArticleID 关系表中Article标的主键
	LikeArticles []Article `gorm:"many2many: user_like_article;joinForeignKey:UserID;JoinReferences:ArticleID" json:"likeArticles"`
}
