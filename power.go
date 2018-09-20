package main

import (
	"flag"
	"fmt"
)

// Person struct for name and battery
type Person struct {
	name    string
	battery *int
}

func main() {
	names := []string{"dani", "gellert", "gergo"}

	var people []Person
	for _, p := range names {
		people = append(people, Person{p, flag.Int(p, 100, fmt.Sprintf("%s laptopjanak toltottsegi szintje", p))})
	}

	flag.Parse()

	selectedPerson := personWithLowestBattery(people)
	fmt.Printf("%s laptopjat kell tolteni\n", selectedPerson.name)
}

func personWithLowestBattery(people []Person) Person {
	minBattery := 100
	selectedPerson := people[0]
	for _, person := range people {
		if *person.battery < minBattery {
			minBattery = *person.battery
			selectedPerson = person
		}
	}
	return selectedPerson
}
