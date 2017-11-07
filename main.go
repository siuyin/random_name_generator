package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"
)

// N is the names struct.
type N struct {
	S []string `yaml:"surnames"`
	G []string `yaml:"givennames"`
}

func getNames(r io.Reader) (N, error) {
	var nam N
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nam, err
	}

	if err := yaml.Unmarshal(b, &nam); err != nil {
		return nam, err
	}
	return nam, nil
}
func randSurName(n N) string {
	return n.S[rand.Intn(len(n.S))]
}
func randGivenName(n N) string {
	return n.G[rand.Intn(len(n.G))]
}
func main() {
	fmt.Println("random name generator")
	file, err := os.Open("data.yaml")
	if err != nil {
		log.Fatal("ERROR: could not open data file: data.yaml: ", err)
	}
	defer file.Close()

	nam, err := getNames(file)
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		fmt.Println(randSurName(nam), randGivenName(nam), randGivenName(nam))
	}
}
