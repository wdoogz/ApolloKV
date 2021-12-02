package main

import (
	"log"

	"github.com/wdoogz/ApolloKV/api"
)

func main() {
	log.Print("[INFO] Starting ApolloKV...")
	log.Print("[INFO] ApolloKV started")
	api.HandleRequests()
}
