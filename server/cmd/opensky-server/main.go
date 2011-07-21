package main

import (
	"fmt"
	"opensky/people"
	"opensky/rooms"
)

func main() {
	fmt.Println("OpenSkyScraper beta server version 0.0.0.1")
	fmt.Println("Loading...")
	
	// Avoid an error about unused packages
	foo := new(people.Person)
	bar := new(rooms.Room)
	_, _ = foo, bar
	
	
}