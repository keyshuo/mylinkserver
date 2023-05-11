package community

import "time"

type UserComment struct {
	Username string
	Time     time.Time
	Content  string
}
