package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DatabaseUsernameConst = "root"
	DatabasePasswordConst = "admin"
	DatabaseHostConst     = "postgres"
	DatabasePortConst     = "5432"
	DatabaseNameConst     = "brickdb"
)

type DatabaseFactory struct {
	dbUserName string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
}

func NewDatabaseFactory() *DatabaseFactory {
	return &DatabaseFactory{
		dbUserName: DatabaseUsernameConst,
		dbPassword: DatabasePasswordConst,
		dbHost:     DatabaseHostConst,
		dbPort:     DatabasePortConst,
		dbName:     DatabaseNameConst,
	}
}

func (databaseFactory *DatabaseFactory) CreatePostgresConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		databaseFactory.dbUserName,
		databaseFactory.dbPassword,
		databaseFactory.dbName,
		databaseFactory.dbHost,
		databaseFactory.dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
