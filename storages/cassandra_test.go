package storages

import (
	"flag"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/gocql/gocql"
)

//Create a fake storage config
var casStorageConfig = Storage{
	ID:      "id",
	Created: "created_datetime",
	Updated: "updated_datetime",
	Table:   "post",
}

var (
	flagCasHost     = flag.String("cashost", "127.0.0.1", "a comma-separated list of host:port tuples")
	flagCasPort     = flag.String("casport", "9042", "the db port listening")
	flagCasKeyspace = flag.String("caskeyspace", "example", "keyspace")
)

func init() {
	flag.Parse()
	casStorageConfig.Host = *flagCasHost
	casStorageConfig.Port = *flagCasPort
	casStorageConfig.Keyspace = *flagCasKeyspace

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	initCassandra()
}

func initCassandra() {
	// Create cassandra initial test
	cluster := gocql.NewCluster(casStorageConfig.Host)
	session, err := cluster.CreateSession()

	if err != nil {
		fmt.Printf("createSession: %v", err)
	}

	// Drop keyspace if exist
	session.Query(`DROP KEYSPACE IF EXISTS ` + casStorageConfig.Keyspace).Exec()

	// Create keyspace
	err = session.Query(`CREATE KEYSPACE example
	WITH replication = {
		'class' : 'SimpleStrategy',
		'replication_factor' : 1
	}`).Consistency(gocql.One).Exec()

	if err != nil {
		fmt.Printf("error creating keyspace %s: %v", "example", err)
	}

	// Create Table to test
	err = session.Query(`CREATE TABLE example.post (
	id			  UUID,
	title         varchar,
	created_datetime timestamp,
	updated_datetime timestamp,
	PRIMARY KEY (title, id)
	)`).Consistency(gocql.One).Exec()

	// Insert data into table
	if err := session.Query(`INSERT INTO example.post (id, title, created_datetime, updated_datetime) values (?, ?,?,?)`, uuid(), "Title 1", time.Now().Truncate(time.Millisecond).UTC(), time.Now().Truncate(time.Millisecond).UTC()).Exec(); err != nil {
		fmt.Println("insert:", err)
	}

	if err := session.Query(`INSERT INTO example.post (id, title, created_datetime, updated_datetime) values (?, ?,?,?)`, uuid(), "Title 2", time.Now().Truncate(time.Millisecond).UTC(), time.Now().Truncate(time.Millisecond).UTC()).Exec(); err != nil {
		fmt.Println("insert:", err)
	}

	if err := session.Query(`INSERT INTO example.post (id, title, created_datetime, updated_datetime) values (?, ?,?,?)`, uuid(), "Title 3", time.Now().Truncate(time.Millisecond).UTC(), time.Now().Truncate(time.Millisecond).UTC()).Exec(); err != nil {
		fmt.Println("insert:", err)
	}

	session.Close()
}

func TestCreateSessionAndConnect(t *testing.T) {
	driver := NewCassandraSession(&casStorageConfig)

	if driver == nil {
		t.Fatalf("error creating driver: ", driver)
	}

	err := driver.Connect()

	if err != nil {
		t.Fatalf("cant connect to cassandra: %v", err)
	}
}

func TestGetAllCasData(t *testing.T) {
	driver := NewCassandraSession(&casStorageConfig)
	driver.Connect()

	allIds, err := driver.GetAll()

	if err != nil {
		t.Fatalf("error: ", err)
	}

	if len(allIds) != 3 {
		t.Fatalf("expected 3, got %v", len(allIds))
	}
}
