package main

import (
	"backend/api"
	"backend/db"
	"backend/lib"
	"os"
)

func main() {
	address := os.Args[1]
	db.InitDB()
	s := lib.SetupCron()
	r := api.GetRouter()
	s.Start()
	defer s.Shutdown()
	r.Run(address)
}