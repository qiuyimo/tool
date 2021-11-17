package main

import (
	"flag"
	"github.com/liudng/godump"
)

var name string

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "go", "set your name")
		_ = goCmd.Parse(args[1:])
	case "php":
		goCmd := flag.NewFlagSet("php", flag.ExitOnError)
		goCmd.StringVar(&name, "n", "php", "set your name")
		_ = goCmd.Parse(args[1:])
	}

	godump.Dump(name)
}
