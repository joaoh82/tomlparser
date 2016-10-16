package main

import (
	"fmt"
	"log"
	"time"

	"github.com/BurntSushi/toml"
)

// Config is a struct defining the content of the TOML file
type Config struct {
	Age        int
	Cats       []string
	Pi         float64
	Perfection []int
	DOB        time.Time
}

func main() {
	config := parse("mytoml.toml")

	fmt.Println("Age: ", config.Age)
	for index, element := range config.Cats {
		fmt.Printf("Cat %d: %s\n", index, element)
	}
	fmt.Println("Pi: ", config.Pi)
	for index, element := range config.Perfection {
		fmt.Printf("Perfection %d: %d\n", index, element)
	}
	fmt.Println("DOB: ", config.DOB)
}

func parse(path string) Config {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
