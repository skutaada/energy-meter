package main

import (
	"backend/api"
	"backend/db"
	"backend/lib"
	"fmt"
	"os"
)

func main() {
	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	db.InitDB()
	s := lib.SetupCron()
	r := api.GetRouter()
	s.Start()
	defer s.Shutdown()
	r.Run(addr)
}