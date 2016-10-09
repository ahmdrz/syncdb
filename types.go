package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Home      string  `json:"home"`
	Away      string  `json:"away"`
	Tables    []Table `json:"tables"`
	Direction bool    `json:"direction"`
}

type Table struct {
	Name   string `json:"name"`
	Column string `json:"column"`
}

func Read() Config {
	config := Config{}
	bytes, err := ioutil.ReadFile("config.json")
	Panic(err)
	err = json.Unmarshal(bytes, &config)
	Panic(err)
	return config
}
