package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type conf struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
}

func (c *conf) getConfig() *conf {
	yamlFile, err := os.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c

}

func main() {
	var c conf
	c.getConfig()

	fmt.Println(c)
}
