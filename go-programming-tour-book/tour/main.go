package main

import (
	"flag"
	"github.com/go-programming-tour-book/tour/cmd"
	"log"
)

var name string

func main() {

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v\n", err)
	}
	//flag01()
	//flag02()
}
func flag01() {
	var name string
	flag.StringVar(&name, "name", "go 语言编程之旅", "help")
	flag.StringVar(&name, "n", "go 语言编程之旅", "help")
	flag.Parse()
	log.Printf("name：%s\n", name)

}
func flag02() {
	flag.Parse()
	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "go 语言", "help")
	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "n", "php 语言", "help")

	args := flag.Args()
	switch args[0] {
	case "go":
		_ = goCmd.Parse(args[1:])
	case "php":
		_ = phpCmd.Parse(args[1:])
	}
	log.Printf("name: %s\n", name)
}
