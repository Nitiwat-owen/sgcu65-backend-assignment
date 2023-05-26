package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sgcu65-backend-assignment/src/config"
	"strconv"
)

func InitPostgresDatabase(conf *config.Postgres) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", conf.Host, strconv.Itoa(conf.Port), conf.Username, conf.Password, conf.Name, conf.SSL)

	gormConf := &gorm.Config{}

	db, err = gorm.Open(postgres.Open(dsn), gormConf)
	if err != nil {
		return nil, err
	}

	return
}