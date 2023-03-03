/**
  @author: Zero
  @date: 2023/3/3 15:52:45
  @desc: 公共实体结构

**/

package custom

import (
	"gorm.io/gorm"
	"time"
)

// BaseModel 所有数据层实体的基层
type BaseModel struct {
	// ID主键
	ID uint64 `gorm:"primaryKey"`
	// 创建时间
	CreateTime time.Time `gorm:"autoCreateTime"`
	// 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime"`
	// 逻辑删除
	Deleted gorm.DeletedAt
}
