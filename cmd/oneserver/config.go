package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type config struct {
	MongoURI          string `yaml:"mongoURI"`
	MessageDB         string `yaml:"messageDB"`
	MessageColl       string `yaml:"messageColl"`
	NatsAddr          string `yaml:"natsAddr"`
	NatsSubjectPrefix string `yaml:"natsSubjectPrefix"`
}

func parseConfigFromFile(filename string) *config {
	var conf config
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(b, &conf)
	if err != nil {
		log.Fatalln(err)
	}
	return &conf
}
