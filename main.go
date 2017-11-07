package main

import (
	"log"
	"os"
	"strconv"

	flag "github.com/spf13/pflag"
)

func main() {
	// port to listen
	port := getPort()
	// create a new server
	server := newServer()
	// listen
	server.Run(":" + port)
}

func getPort() string {
	defaultPort := "8080"
	// get port from env
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	// get port from flag
	pPort := flag.StringP("port", "p", defaultPort, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	// validate
	if _, err := strconv.ParseUint(port, 10, 16); err != nil {
		log.Fatalln("Invalid port", port)
	}
	return port
}
