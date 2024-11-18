package lib

import (
	"backend/db"
	"log"

	"github.com/go-co-op/gocron/v2"
)

func SetupCron() gocron.Scheduler {
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal("Error setting up scheduler")
	}
	_, err = s.NewJob(
		gocron.CronJob(
			"0 1 * * *",
			false,
		),
		gocron.NewTask(db.GetLatestEnergy),
	)
	if err != nil {
		log.Fatalf("Error setting up job: %s", err)
	}

	return s
}