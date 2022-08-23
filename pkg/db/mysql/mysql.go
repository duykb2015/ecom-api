package gorm

import (
	"database/sql"
	"time"

	"github.com/duykb2015/ecom-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(cfg *config.MySQL) (*gorm.DB, error) {
	db, err := sql.Open("mysql", cfg.URI)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifeTime))

	return gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
}
