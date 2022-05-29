package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type MonitorSchema struct {
	Monitor []Monitor `yaml:"monitor"`
}

type Monitor struct {
	Feed []Feed `yaml:"feed"`
}

type Feed struct {
	Name    string      `yaml:"name"`
	FeedUrl string      `yaml:"feed_url"`
	Item    interface{} `yaml:"item"`
}

func main(){

	file, err := ioutil.ReadFile("feed.yaml")
	if err != nil {
		fmt.Println("Error reading config file: ", err)
	}

	var monitor MonitorSchema

	err = yaml.Unmarshal(file, &monitor)
	if err != nil {
		fmt.Println("Error parsing config file: ", err)
	}

	// fmt.Println(monitor.Monitor[0].Feed[0].FeedUrl)
	for _,feed := range monitor.Monitor {
		for _, item := range feed.Feed{
			fmt.Println(item.FeedUrl)
		}
	}

	// for _,feed := range monitor.
}