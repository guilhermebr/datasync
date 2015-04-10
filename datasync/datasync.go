package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// Storage .
type storage struct {
	Driver   string
	Host     string
	Port     string
	Keyspace string
	Table    string
	ID       string
	Created  string
	Updated  string
}

type configuration struct {
	App      string
	Timer    int
	Storages []*storage
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
	go func() {
		for {
			fmt.Println(config)
			select {
			case <-t.C:
			case <-quit:
				return
			}
		}
	}()
}
