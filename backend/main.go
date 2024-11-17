package main

import (
	"backend/api"
	"backend/db"
)

func main() {
	db.InitDB()
	r := api.GetRouter()
	r.Run()
}
