package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type N struct {
	S []string `yaml:"surnames"`
	G []string `yaml:"givennames"`
}

func main() {
	fmt.Println("random name generator")
	file, err := os.Open("data.yaml")
	if err != nil {
		log.Fatal("ERROR: could not open data file: data.yaml: ", err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("ERROR: failed to read from file: ", err)
	}

	var nam N
	if err := yaml.Unmarshal(b, &nam); err != nil {
		log.Fatal("ERROR: could not unmarshal yaml: ", err)
	}

	rand.Seed(time.Now().UnixNano())
	lSur := len(nam.S)
	lGiv := len(nam.G)
	for i := 0; i < 20; i++ {
		fmt.Println(nam.S[rand.Intn(lSur)], nam.G[rand.Intn(lGiv)], nam.G[rand.Intn(lGiv)])
	}
}
