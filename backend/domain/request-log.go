package domain

import (
	"net"
)

type RequestLog struct {
	PlayerId  int    `json:"player_id" binding:"required"`
	Ip        net.IP `json:"ip" binding:"required"`
	UserAgent string `json:"user_agent" binding:"required"`
	CreatedAt string `json:"created_at"`
	IsBot     bool   `json:"is_bot"`
}
