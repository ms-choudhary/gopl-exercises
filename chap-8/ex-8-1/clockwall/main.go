package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	names := []string{}
	conns := []net.Conn{}
	for _, arg := range os.Args[1:] {
		fields := strings.Split(arg, "=")
		names = append(names, fields[0])

		conn, err := net.Dial("tcp", fields[1])
		if err != nil {
			log.Fatal(err)
		}

		conns = append(conns, conn)
	}

	readers := []*bufio.Reader{}
	for _, c := range conns {
		readers = append(readers, bufio.NewReader(c))
	}

	defer func(cs []net.Conn) {
		for _, c := range cs {
			fmt.Println("closing conn")
			c.Close()
		}
	}(conns)

	for _, n := range names {
		fmt.Printf("%10.10s", n)
	}
	fmt.Println()

	for {
		for _, r := range readers {
			clock, err := r.ReadString('\n')
			if err != nil {
				clock = "error"
			}
			fmt.Printf("%10.10s", clock[:len(clock)-1])
		}
		fmt.Println()
	}
}
