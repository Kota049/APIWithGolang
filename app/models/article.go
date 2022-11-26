package models

import "gorm.io/gorm"

// Articleテーブルのモデル
type Article struct {
	gorm.Model
	User_id uint
	Title   string `gorm:"size:255"`
}
