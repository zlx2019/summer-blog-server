/**
  @author: Zero
  @date: 2023/3/4 20:44:55
  @desc: 菜单表

**/

package models

import "summer/models/custom"

// Menu 菜单表结构
type Menu struct {
	custom.BaseModel
	// 标题
	Title string `json:"title" gorm:"size:32; comment: 标题"`
	// slogan
	Slogan string `json:"slogan" gorm:"size:32; comment: slogan"`
	// 简介
	Desc string `json:"desc" gorm:"size:256; comment: 简介"`
	// 菜单图片列表
	Banners []File `json:"menuImages" gorm:"many2many:menu_file;joinForeignKey:MenuID;JoinReferences:FileID"`
	// 菜单图片切换时间
	BannerTime int `json:"img_time" gorm:"comment: 图片切换间隔时间"`
	// 菜单优先级排序
	Sort int `json:"sort" gorm:"size:10; comment: 菜单优先级排序"`
}
