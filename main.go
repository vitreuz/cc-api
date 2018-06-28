package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vitreuz/hackday/api"
)

const bq = "https://api.beque.cli.fun"

func main() {

	var token string
	if args := os.Args; len(args) > 1 {
		token = args[1]
	}

	client := api.NewClient(bq, token)
	apps, err := client.ListApps()
	if err != nil {
		log.Fatalln(err)
	}

	for _, app := range apps {
		fmt.Printf("%+v", app)
		fmt.Println()
	}
}
