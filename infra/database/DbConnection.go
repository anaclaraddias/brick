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

func (db *DatabaseConnection) Rows(query string, values ...any) ([]map[string]interface{}, error) {
	rows, err := db.connection.Raw(query, values...).Rows()

	if err != nil {
		return nil, fmt.Errorf(ExecuteRowsQueryErrorConst, err)
	}

	columns, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	count := len(columns)
	statementPtrs := make([]interface{}, count)

	for i := range statementPtrs {
		var v interface{}
		statementPtrs[i] = &v
	}

	result := make([]map[string]interface{}, 0)

	for rows.Next() {
		err := rows.Scan(statementPtrs...)
		if err != nil {
			continue
		}
		row := make(map[string]interface{})
		for i, col := range columns {
			val := *(statementPtrs[i].(*interface{}))
			row[col] = handleScanReturn(val)
		}
		result = append(result, row)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

func handleScanReturn(result any) any {
	switch value := result.(type) {
	case bool:
		return value
	case int64:
		return value
	case float32:
		return value
	case float64:
		return value
	case int:
		return value
	case string:
		return value
	case []byte:
		return string(value)
	case time.Time:
		return value.Format(DATETIME_FORMAT)
	case nil:
		return nil
	default:
		fmt.Println("tipo de resultado desconhecido")
		return nil
	}
}
