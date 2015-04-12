package storages

// Storage .
type Storage struct {
	Driver   string
	Host     string
	Port     string
	Keyspace string
	Table    string
	ID       string
	Created  string
	Updated  string
}

// Driver interface
type Driver interface {
	Connect() error
	Debug()
	GetAll() ([]map[string]interface{}, error)
}
