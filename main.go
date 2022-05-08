package main

import (
	"PDDiktiProxyAPI/Modules/ServerCaches"
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

	// Gocron example
	//location, _ := time.LoadLocation("Asia/Jakarta")
	//x := gocron.NewScheduler(location)
	//
	//x.Every(1).Day().At("08:54").Do(func() {
	//	log.Print("Hello!")
	//})
	//
	//x.StartAsync()

	test := ServerCaches.FetchUniversities()

	log.Print(test)

	log.Fatal(RunServer(address))
}
