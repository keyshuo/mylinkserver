package leaderboard

import "time"

type UserComment struct {
	Username string
	Time     time.Time
	Rank     string
}
