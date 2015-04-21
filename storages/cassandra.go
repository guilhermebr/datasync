package storages

import (
	"fmt"

	"github.com/gocql/gocql"
)

// Cassandra .
type Cassandra struct {
	Session *gocql.Session
	Storage
}

// NewCassandraSession .
func NewCassandraSession(storage *Storage) *Cassandra {

	return &Cassandra{nil, *storage}
}

// Connect DB
func (d *Cassandra) Connect() error {
	cluster := gocql.NewCluster(d.Host)
	cluster.Keyspace = d.Keyspace
	session, err := cluster.CreateSession()

	if err != nil {
		return fmt.Errorf("cant create session")
	}

	d.Session = session

	return nil

}

// GetAll returns all data (ID, Created, Updated) in Keyspace/Table
func (d *Cassandra) GetAll() ([]map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT %s, %s, %s FROM %s",
		d.ID, d.Created, d.Updated, d.Table)

	allids, err := d.Session.Query(query).Iter().SliceMap()

	if err != nil {
		return nil, err
	}

	return allids, nil
}
