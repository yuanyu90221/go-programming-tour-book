package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "Go", "info")
	flag.StringVar(&name, "n", "Go", "info")
	flag.Parse()

	log.Printf("name: %s", name)
}
