package resource

import "time"

type Show struct {
	ID            uint32    `json:"id"`
	Name          string    `json:"name"`
	TotalEpisodes uint8     `json:"totalEpisodes"`
	EpisodeTime   uint8     `json:"episodeTime"`
	StoppedAt     uint8     `json:"stoppedAt"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type ShowJSON struct {
	Show Show `json:"show"`
}
