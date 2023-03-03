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
	// ID 主键,禁止使用自动递增
	ID uint64 `gorm:"primaryKey;autoIncrement:false; comment: 主键ID"`
	// 创建时间
	CreateTime time.Time `gorm:"autoCreateTime; comment: 创建时间"`
	// 更新时间
	UpdateTime time.Time `gorm:"autoUpdateTime; comment: 更新时间"`
	// 逻辑删除 建立普通索引
	Deleted gorm.DeletedAt `gorm:"index:idx_del; comment: 逻辑删除"`
}
