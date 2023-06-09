package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"html/template"
	"os"
)

type Attributes struct {
	Rel string `yaml:"rel,omitempty" json:"rel,omitempty"`
}

type Link struct {
	Title      string     `yaml:"title" json:"title"`
	Icon       string     `yaml:"icon" json:"icon"`
	URL        string     `yaml:"url" json:"url"`
	Attributes Attributes `yaml:"attributes" json:"attributes"`
}

type LinkConfig struct {
	Title    string `yaml:"title" json:"title"`
	Name     string `yaml:"name" json:"name"`
	Location string `yaml:"location" json:"location"`
	Links    []Link `yaml:"links" json:"links"`
}

func main() {
	yamlFile, err := os.ReadFile("links.yaml")
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}

	var linkConfig LinkConfig
	err = yaml.Unmarshal(yamlFile, &linkConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	paths := []string{"index.html.tmpl"}
	t := template.Must(template.New("index.html.tmpl").ParseFiles(paths...))
	err = t.Execute(os.Stdout, linkConfig)
	if err != nil {
		panic(err)
	}
}
