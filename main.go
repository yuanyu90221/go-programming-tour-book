package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

// var name string

type Name string

func (i *Name) String() string {
	return fmt.Sprint(*i)
}

func (i *Name) Set(value string) error {
	if len(*i) > 0 {
		return errors.New("name flag already set")
	}
	*i = Name("json:" + value)
	return nil
}
func main() {
	var name Name
	flag.Var(&name, "name", "說明資訊")
	flag.Parse()
	// goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	// goCmd.StringVar(&name, "name", "Go語言", "說明資訊")
	// phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	// phpCmd.StringVar(&name, "n", "PHP語言", "說明資訊")

	// args := flag.Args()
	// if len(args) <= 0 {
	// 	return
	// }

	// switch args[0] {
	// case "go":
	// 	_ = goCmd.Parse(args[1:])
	// case "php":
	// 	_ = phpCmd.Parse(args[1:])
	// }

	log.Printf("name: %s", name)
}
