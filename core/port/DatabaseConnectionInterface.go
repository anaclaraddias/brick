package port

type DatabaseConnectionInterface interface {
	Open() error
	Close() error
	Raw(sql string, statement interface{}, values ...any) error
}
