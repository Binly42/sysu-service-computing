package main

import (
	"cloudgo"
	"flag"
	"log"
)

const (
	DefaultPort = cloudgo.DefaultPort
)

var (
	port string
	// ...
)

func init() {
	flag.StringVar(&port, "p", DefaultPort, "The PORT to be listened by cloudgo.")
}

func main() {
	flag.Parse()
	// TODO: validate port ?

	cloudgo.LoadAll()
	defer cloudgo.SaveAll()

	server := cloudgo.New()
	err := server.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
