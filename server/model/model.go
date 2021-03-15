package model

import "time"

type Question struct {
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	Choices   []Choice  `json:"choices"`
}

type Choice struct {
	Text string `json:"text"`
}
