package mysql

import (
	"database/sql"
	"time"

	"github.com/duykb2015/login-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(cfg *config.Config) (*gorm.DB, error) {
	db, err := sql.Open("mysql", cfg.MySQL.URI)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(cfg.MySQL.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MySQL.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.MySQL.ConnMaxLifeTime))

	return gorm.Open(mysql.New(mysql.Config{Conn: db}), &gorm.Config{})
}
