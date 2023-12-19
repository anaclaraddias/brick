package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/anaclaraddias/brick/core/port"
	"gorm.io/gorm"
)

const (
	ExecuteRowsQueryErrorConst    = "ExecuteRowsQueryError %s"
	ExecuteRawStatementErrorConst = "ExecuteRawStatementError %s"

	DATETIME_FORMAT = "2006-01-02 15:04:05"
)

type DatabaseConnection struct {
	connection *gorm.DB
}

func NewDatabaseConnection() port.DatabaseConnectionInterface {
	return &DatabaseConnection{}
}

func (db *DatabaseConnection) Open() error {
	var database *gorm.DB
	var err error

	maxRetries := 10
	retryInterval := 5 * time.Second

	dbFactory := NewDatabaseFactory()

	for retryCount := 1; retryCount <= maxRetries; retryCount++ {
		database, err = dbFactory.CreatePostgresConnection()
		if err == nil {
			break
		}

		log.Printf("failed to connect to the database. Retrying in %v...", retryInterval)
		time.Sleep(retryInterval)
	}

	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	sqlDB, ok := database.ConnPool.(*sql.DB)
	if !ok {
		return fmt.Errorf("unexpected connection pool type")
	}

	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	db.connection = database
	log.Printf("database connection successfully opened")
	return nil
}

func (db *DatabaseConnection) Close() error {
	sqlDB, err := db.connection.DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("falha ao fechar conexao com o banco: %w", err)
	}

	log.Println("database connection successfully closed")

	return nil
}

func (db *DatabaseConnection) Raw(query string, statment interface{}, values ...any) error {
	transaction := db.connection.Raw(query, values...).Scan(statment)

	if transaction.Error != nil {
		return fmt.Errorf(ExecuteRawStatementErrorConst, transaction.Error)
	}

	return nil
}
