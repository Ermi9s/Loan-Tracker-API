package domain

import "time"

type Logs struct {
	Event string `json:"event" bson:"event"`
	Time  time.Weekday `json:"time" bson:"time"`
}

