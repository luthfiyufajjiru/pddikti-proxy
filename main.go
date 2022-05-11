package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var (
		ipaddress string = os.Getenv("IP_ADDRESS")
		port      string = os.Getenv("PORT")
		address   string = fmt.Sprintf("%v:%v", ipaddress, port)
	)

	log.Fatal(RunServer(address))
}
