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


	for _,feed := range monitor. {
		fmt.Print(feed)
	}
}