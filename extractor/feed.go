package extractor

import (
	"fmt"
	"gopkg.in/yaml.v3"
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

func GetFeedURL(data []byte) []string {

	var ParserURI []string

	var monitor MonitorSchema

	err := yaml.Unmarshal(data, &monitor)
	if err != nil {
		fmt.Println("Error parsing config file: ", err)
	}

	for _, feed := range monitor.Monitor {
		for _, item := range feed.Feed {
			ParserURI = append(ParserURI, item.FeedUrl)
		}
	}

	return ParserURI
}
