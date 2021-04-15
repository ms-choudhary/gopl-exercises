package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		r := bufio.NewReader(c)
		cmd, err := r.ReadString('\n')
		if err != nil {
			return
		}

		switch strings.Split(cmd, " ")[0] {
		case "PORT":
		case "CHDIR":
		}
	}
}

func writeData(c net.Conn, data []byte) error {
	defer c.Close()
	n, err := c.Write(data)
	if err != nil {
		return err
	}
	if n != len(data) {
		return fmt.Errorf("couldn't write data")
	}
}
