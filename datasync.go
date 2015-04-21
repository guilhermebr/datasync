package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/guilhermebr/datasync/storages"
)

type configuration struct {
	App      string
	Timer    int
	Storages []*storages.Storage
}

// Projects found in config folder .
var Projects []*configuration

var quit = make(chan bool)

func main() {
	// Load config projects file
	LoadProjects()

	// Create a timer for each project
	for _, project := range Projects {
		scheduleProject(project)

	}
	<-quit

}

// LoadProjects run config folder and load projects into Project
func LoadProjects() {
	files, _ := ioutil.ReadDir("./config/")
	for _, f := range files {
		fmt.Println("Found config file: ", f.Name())
		file, err := os.Open("./config/" + f.Name())
		if err != nil {
			fmt.Println("error when open file:", err)

		}
		decoder := json.NewDecoder(file)
		config := configuration{}
		err = decoder.Decode(&config)
		if err != nil {
			fmt.Println("error decoding file: ", err)
		}
		fmt.Println(config.App)
		Projects = append(Projects, &config)
	}

}

func scheduleProject(config *configuration) {
	// Create a ticker based on timer option in config file
	t := time.NewTicker(time.Duration(config.Timer) * time.Second)

	// Get all storages in config file
	drivers := parseStorage(config)

	// run a goroutine that get data from storages and synchronize
	go func() {
		for {
			// Connect and get all data from first storage
			fmt.Println(config)
			drivers[0].Connect()
			d1Ids, _ := drivers[0].GetAll()

			for _, row := range d1Ids {
				fmt.Println(row)
			}

			// Connect and get all data from second storage
			drivers[1].Connect()
			d2Ids, _ := drivers[1].GetAll()
			for _, row := range d2Ids {
				fmt.Println(row)
			}

			// Wait for timer and run again
			select {
			case <-t.C:
			case <-quit:
				return
			}
		}
	}()
}

// parseStorage get storages in config file and get the Driver
func parseStorage(config *configuration) []storages.Driver {
	var storagesReturn []storages.Driver
	for _, storage := range config.Storages {
		if storage.Driver == "cassandra" {
			storagesReturn = append(storagesReturn, storages.NewCassandraSession(storage))
		} else if storage.Driver == "elasticsearch" {
			storagesReturn = append(storagesReturn, storages.NewElasticSearchSession(storage))
		}
	}

	return storagesReturn
}
