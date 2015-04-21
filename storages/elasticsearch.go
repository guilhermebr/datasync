package storages

import (
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic"
)

// ElasticSearch .
type ElasticSearch struct {
	client *elastic.Client
	Storage
}

// NewElasticSearchSession .
func NewElasticSearchSession(storage *Storage) *ElasticSearch {
	return &ElasticSearch{nil, *storage}
}

// Connect DB
func (d *ElasticSearch) Connect() error {
	dburl := fmt.Sprintf("http://%s:%s", d.Host, d.Port)
	client, err := elastic.NewClient(
		elastic.SetURL(dburl),
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	d.client = client

	return nil
}

// GetAll returns all data (ID, Created, Updated) in Index/Type
func (d *ElasticSearch) GetAll() ([]map[string]interface{}, error) {
	var returnIds []map[string]interface{}

	searchResult, err := d.client.Search().
		Index(d.Index). // search in index d.Index
		Type(d.Table).
		Sort("updated", true).
		Pretty(true). // pretty print request and response JSON
		Do()          // execute
	if err != nil {
		// Handle error
		return nil, err
	}

	// fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	if searchResult.Hits != nil {
		// fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// Deserialize hit.Source into a map[string]interface{}.
			var t map[string]interface{}
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				return nil, err
			}

			t[d.ID] = hit.Id

			returnIds = append(returnIds, t)

		}
	}

	return returnIds, nil
}

// func get() {
// 	/ Get document
// 	getResult, err := client.Get().
// 		Index(testIndexName).
// 		Type("tweet").
// 		Id("1").
// 		Do()
// }
