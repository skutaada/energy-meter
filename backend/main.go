package main

import (
	"backend/api"
	"backend/db"
	"backend/lib"
	"fmt"
	"os"
)

func main() {
	addr := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
	db.InitDB()
	s := lib.SetupCron()
	r := api.GetRouter()
	s.Start()
	defer s.Shutdown()
	r.Run(addr)
}