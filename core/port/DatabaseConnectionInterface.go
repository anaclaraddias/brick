package port

type DatabaseConnectionInterface interface {
	Open() error
	Close() error
	Raw(sql string, statement interface{}, values ...any) error
	Rows(sql string, values ...any) ([]map[string]interface{}, error)
}
