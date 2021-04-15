package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	results, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d results\n", results.TotalCount)
	for _, item := range results.Items {
		fmt.Printf("#%-5d %9.9s %55.55s %10.10s\n", item.Number, item.User.Login, item.Title, item.Age())
	}
}
