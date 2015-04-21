package storages

import (
	"flag"
	"log"
	"testing"
	"time"
)

var esStorageConfig = Storage{
	Index: "example",
	Table: "tweet",
}

var (
	flagEsHost = flag.String("eshost", "127.0.0.1", "a comma-separated list of host:port tuples")
	flagEsPort = flag.String("esport", "9200", "the db port listening")
)

func createIndexAndData(t *testing.T, driver *ElasticSearch) {
	client := driver.client

	// Delete the index
	client.DeleteIndex(driver.Index).Do()

	// Create an index
	_, err := client.CreateIndex(driver.Index).Do()
	if err != nil {
		t.Fatal(err)
	}

	// Add a document to the index
	type Tweet struct {
		User    string    `json:"user"`
		Message string    `json:"message"`
		Created time.Time `json:"created,omitempty"`
		Updated time.Time `json:"updated,omitempty"`
	}

	_, err = client.Index().
		Index(driver.Index).
		Type(driver.Table).
		Id(uuid()).
		BodyJson(Tweet{
		User:    "raul",
		Message: "Rock baby",
		Created: time.Now().Truncate(time.Millisecond).UTC(),
		Updated: time.Now().Truncate(time.Millisecond).UTC()}).
		Do()

	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Flush().Index(driver.Index).Do()
	if err != nil {
		t.Fatal(err)
	}
}

func init() {
	flag.Parse()
	esStorageConfig.Host = *flagEsHost
	esStorageConfig.Port = *flagEsPort
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func TestSessionAndConnection(t *testing.T) {
	// Create ElasticSearch driver
	driver := NewElasticSearchSession(&esStorageConfig)

	if driver == nil {
		t.Fatalf("error creating driver: ", driver)
	}

	// Connect to the DB
	err := driver.Connect()

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAllESData(t *testing.T) {
	driver := NewElasticSearchSession(&esStorageConfig)
	driver.Connect()
	createIndexAndData(t, driver)

	allIds, err := driver.GetAll()

	if err != nil {
		t.Fatalf("error: ", err)
	}

	if len(allIds) != 1 {
		t.Fatalf("expected 1, got %v", len(allIds))
	}
}
