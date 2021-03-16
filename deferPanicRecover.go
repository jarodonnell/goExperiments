package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	type State struct {
		number    int
		Name      string
		bird      string
		Cities    map[string]int
		Parks     []string
		neighbors []State
		Color     string
	}
	Alaska := State{}

	emptyStates := make(map[int]string)
	states := map[int]string{
		1: "Alaska",
		2: "Alabama",
		3: "Arkansas",
		4: "California",
	}
	allCities := map[int]map[string]int{
		1: map[string]int{
			"Juneau":    12345,
			"Fairbanks": 890,
			"homer":     897,
		},
		2: map[string]int{
			"Birmingham": 123434,
			"Mobile":     879,
		},
	}
	akCities := map[string]int{
		"Juneau":    12345,
		"Fairbanks": 890,
		"homer":     897,
	}
	akParks := [...]string{"Denali", "McKinley"}

	Alaska.Cities = akCities
	Alaska.Parks = []string{"Denali", "McKinley"}

	akCitiesA := [...]string{
		"Juneau",
		"Fairbanks",
		"homer",
	}

	moreStates := map[int]string{
		1: "Alaska",
		2: "Alabama",
		3: "Arkansas",
		4: "California",
	}
	moreStates = states
	type Docttor struct {
		number     int
		actorName  string
		episodes   []string
		companions []string
	}

	aDoctor := Docttor{
		actorName: "jon",
		number:    3,
		companions: []string{
			"liz",
			"jo	",
			"sarah",
		},
	}
	anotherDoc := aDoctor
	anotherDoc.actorName = "tom"
	fmt.Println("adoc", aDoctor)
	fmt.Println("anotherdoc", anotherDoc)
	fmt.Println("adoc", aDoctor)
	fmt.Println("anotherdoc", anotherDoc)
	fmt.Println(states[3])
	moreStates[3] = "Oklahoma"
	name, ok := states[5]
	fmt.Println("states 5", name, ok)
	name, ok = states[4]
	fmt.Println("states 4", name, ok)
	fmt.Println(states[3])
	fmt.Println(moreStates)
	fmt.Println(emptyStates)

	fmt.Println(allCities, ok)
	fmt.Println(akCities, ok)
	fmt.Println(akCitiesA, ok)
	fmt.Println(akParks, ok)
	k := 1
	v := ""
	for k, v := range moreStates {
		fmt.Printf("state %d is %s\n", k, v)

	}
	// k and v were shadowed in the loop
	fmt.Printf("state %d is %s\n", k, v)
	err := WriteToFile("result.txt", "Hello Readers of golangcode.com\n")
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range moreStates {
		
		if err != nil {
			log.Fatal(err)
		}
		mapEntry := fmt.Sprintf("state %d is %s\n", k, v)
		err = WriteToFile("result.txt", mapEntry)

	}
}

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end.
func WriteToFile(filename string, data string) error { //   rwx         r--       r--
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //  0110 (6)    0100 (4)  0100    (4)
	if err != nil {                                                               //   Owner Group Other
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
