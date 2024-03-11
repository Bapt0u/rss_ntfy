package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	Notify string   `yaml:"notify"`
	Feeds  []string `yaml:"feeds"`
}

func (c *Conf) Config(path string) *Conf {
	CheckConfigFile(path)
	conf, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error loading file at: %s", path)
	}

	err = yaml.Unmarshal(conf, c)
	if err != nil {
		log.Printf("Error unmarshalling: %e", err)
	}

	return c
}

func CheckConfigFile(path string) error {
	file, err := os.Stat(path)
	log.Println(file)
	if err != nil {
		log.Fatalf("%e", err)
		os.Exit(1)
	}
	log.Printf("Loading config file %s", path)
	return nil
}
