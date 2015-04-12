package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/guilhermebr/datasync/datasync/storages"
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
	t := time.NewTicker(time.Duration(config.Timer) * time.Second)
	drivers := parseStorage(config)

	go func() {
		for {
			fmt.Println(config)
			drivers[0].Connect()
			d1Ids, _ := drivers[0].GetAll()

			// if err != nil {
			// 	fmt.Println("error: ", err)
			// }

			for _, row := range d1Ids {
				fmt.Println(row)
			}
			// driver.Debug()
			select {
			case <-t.C:
			case <-quit:
				return
			}
		}
	}()
}

func parseStorage(config *configuration) []storages.Driver {
	var storagesReturn []storages.Driver
	for _, storage := range config.Storages {
		if storage.Driver == "cassandra" {
			storagesReturn = append(storagesReturn, storages.NewCassandraSession(storage))

		} else if storage.Driver == "elasticsearch" {
			storagesReturn = append(storagesReturn, storages.NewElasticSearchSession(storage))

		}
	}

	return nil
}
