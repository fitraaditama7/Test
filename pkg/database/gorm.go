package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"test/cmd/entity"
)

// DBConfig Database Config
type DBConfig struct {
	Host     string `envconfig:"HOST"`
	User     string `envconfig:"USER"`
	Port     int    `envconfig:"PORT"`
	Password string `envconfig:"PASSWORD"`
	Name     string `envconfig:"NAME"`
}

func Connect(cfg DBConfig) *gorm.DB {
	var err error

	dbConfig := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)

	db, err := gorm.Open(
		postgres.Open(dbConfig),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{})

	return db
}
