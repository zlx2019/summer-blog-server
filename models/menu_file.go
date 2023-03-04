/**
  @author: Zero
  @date: 2023/3/4 20:54:50
  @desc: 菜单与图片关联表

**/

package models

import "summer/models/custom"

// MenuFile 菜单与图片关联表结构
type MenuFile struct {
	custom.BaseModel
	// 菜单ID
	MenuID uint64 `json:"menuID" gorm:"comment: 菜单ID"`
	Menu   Menu   `gorm:"foreignKey:MenuID"`
	// 文件ID
	FileID uint64 `json:"fileID" gorm:"文件ID"`
	File   File   `gorm:"foreignKey:FileID"`
	Sort   int    `json:"sort" gorm:"size: 10; comment: 排序字段"`
}
