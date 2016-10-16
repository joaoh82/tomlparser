package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	content := []byte(`
Age = asd
Cats = [ "Bill", "Murray", "Otto" ]
Pi = 3.14
Perfection = [ 6, 28, 496, 8128, 23987 ]
DOB = 1982-11-25T07:42:00Z
`)

	tmpfile, err := ioutil.TempFile("", "temp")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	defer os.Remove(tmpfile.Name()) // Clean Up
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal("Error: ", err)
	} else {
		fmt.Println(tmpfile.Name())
		config := parse(tmpfile.Name())
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

	if err := tmpfile.Close(); err != nil {
		log.Fatal("Error: ", err)
	}
}
