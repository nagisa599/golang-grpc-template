package repository

import (
	"context"

	"gorm.io/gorm"
)

type DatabaseHandler interface {
	// 接続の関数
	Conn(context.Context) *gorm.DB

	// 接続の終了
	Close()
}