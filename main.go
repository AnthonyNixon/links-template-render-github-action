package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"html/template"
	"log"
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
	GitSha   string `yaml:"git_sha" json:"git_sha"`
}

func main() {
	templateFile := os.Args[1]
	valuesFile := os.Args[2]
	outputFile := os.Args[3]
	fmt.Printf("Parsing template '%s' with values in file '%s' to output file '%s'\n", templateFile, valuesFile, outputFile)

	yamlFile, err := os.ReadFile(valuesFile)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}

	fmt.Printf("values file contents:\n%s\n", string(yamlFile))

	var linkConfig LinkConfig
	err = yaml.Unmarshal(yamlFile, &linkConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	linkConfig.GitSha = os.Getenv("GIT_SHA")
	if linkConfig.GitSha != "" {
		fmt.Printf("Using Git SHA: %s\n", linkConfig.GitSha)
	}

	output, err := os.Create(outputFile)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	paths := []string{templateFile}
	t := template.Must(template.New(templateFile).ParseFiles(paths...))
	err = t.Execute(output, linkConfig)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully rendered %s\n", outputFile)
}
