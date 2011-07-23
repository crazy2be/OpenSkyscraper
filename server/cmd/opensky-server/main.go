package main

import (
	"fmt"
	"log"
	"net"
	"opensky/people"
	"opensky/rooms"
	"opensky/packets"
)

func main() {
	fmt.Println("OpenSkyScraper beta server version 0.0.0.1")
	fmt.Println("Loading...")
	
	// Avoid an error about unused packages
	foo := new(people.Person)
	bar := new(rooms.Room)
	_, _ = foo, bar
	
	listener, err := net.Listen("tcp", ":2898")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Loaded and accepting connections on tcp port 2898!")
	for {
		rawconn, err := listener.Accept()
		if err != nil {
			// TODO: Better error handling
			log.Fatal(err)
		}
		
		go func() {
			pc, err := packets.NewPlayerConn(rawconn)
			if err != nil {
				log.Fatal(err)
			}
			for {
				pc.NextPacket()
			}
		}()
	}
}