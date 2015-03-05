package resource

import "time"

type Show struct {
	Id              int32     `json:"id"`
	Name            string    `json:"name"`
	TotalEpisodes   uint8     `json:"total_episodes"`
	WatchedEpisodes uint8     `json:"watched_episodes"`
	CreatedAt       time.Time `json:"created"`
}
