package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"io"
)

func main() {
	ln, err := net.Listen("tcp", ":1111")
	if err != nil {
		log.Fatalln("Unable to open listening port:", err)
	}
	
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Unable to accept connection:", err)
			continue
		}
		go debugPrint(conn)
	}
}

func debugPrint(conn net.Conn) {
	rd := bufio.NewReader(conn)
// 	wr := bufio.NewWriter(conn)
	for {
		p, err := ReadPacket(rd)
		if err != nil {
			log.Println("Error reading packet:", err)
			return
		}
		for {
			b, err := p.buf.ReadByte()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Println("Error reading from packet:", err)
				return
			}
			fmt.Printf("Got byte: %d %s\n", b, string(b))
		}
	}
}