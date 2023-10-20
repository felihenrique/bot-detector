package app

import (
	"net"
	"time"
)

type RequestLog struct {
	PlayerId  int       `json:"player_id" binding:"required"`
	Ip        net.IP    `json:"ip" binding:"required"`
	UserAgent string    `json:"user_agent" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	IsBot     bool      `json:"is_bot"`
}
