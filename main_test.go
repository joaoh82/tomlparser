package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestParser(t *testing.T) {
	//Content for temporary file
	content := []byte(`
    Age = 52
    Cats = [ "Bill", "Murray", "Otto" ]
    Pi = 3.14
    Perfection = [ 6, 28, 496, 8128, 23987 ]
    DOB = 1979-05-27T07:32:00Z
    `)
	//Creates a temporaty file in a specified directory and a prefix in the name
	tmpfile, err := ioutil.TempFile("", "temp")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	defer os.Remove(tmpfile.Name()) // Clean Up
	//Writes the content in the temporary file
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal("Error: ", err)
	} else {
		//Created an object the same as the content the should be in the file to be compared
		time, err := time.Parse(time.RFC3339, "1979-05-27T07:32:00Z")
		if err != nil {
			panic(err)
		}
		var answer = Config{
			Age:        52,
			Cats:       []string{"Bill", "Murray", "Otto"},
			Pi:         3.14,
			Perfection: []int{6, 28, 496, 8128, 23987},
			DOB:        time,
		}
		//Calls our main funcion in the main.go file
		config := parse(tmpfile.Name())
		//Actually compares boths contents
		if !reflect.DeepEqual(config, answer) {
			t.Fatalf("Expected\n-----\n%#v\n-----\nbut got\n-----\n%#v\n",
				answer, config)
		} else {
			fmt.Println("Passed! Result: ")
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
	}
	//Closes the temporary file
	if err := tmpfile.Close(); err != nil {
		log.Fatal("Error: ", err)
	}
}
