package main

import (
	"flag"
	"log"
)

func main() {
	var name string

	flag.StringVar(&name, "name", "QY", "名称")

	flag.StringVar(&name, "n", "QY", "名称")

	flag.Parse()

	log.Printf("name: %s", name)
}
