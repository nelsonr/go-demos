package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	name string
	age  int
}

func main() {
	// Declare variables
	var people = make([]Person, 0)
	var oldestPerson Person

	// Read the contents of the text file
	contents, err := os.ReadFile("names.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Split the contents into a list (slice) of lines
	lines := strings.Split(string(contents), "\r\n")

	// Iterate over each line and extract the name and age into a Person struct
	for _, line := range lines {
		name, age, _ := strings.Cut(line, ", ")

		// Skip over empty lines
		if name == "" {
			continue
		}

		// Convert the age from a `string` to an `int`
		ageInt, err := strconv.Atoi(age)
		if err != nil {
			log.Fatal(err)
		}

		// Append the Person to the list
		people = append(people, Person{name: name, age: ageInt})
	}

	// Find out who's the oldest person
	for _, person := range people {
		if person.age > oldestPerson.age {
			oldestPerson = person
		}
	}

	fmt.Printf("The oldest of the bunch is %s with %d years old.", oldestPerson.name, oldestPerson.age)
}
