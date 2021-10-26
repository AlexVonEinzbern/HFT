package model

import (
	"gorm.io/gorm"
	"time"
)

type PeerBase struct {
	Host string `json:"host"`
}

type Peer struct {
	gorm.Model
	PeerBase
	File string `json:"file"`
}

type PeerCreate struct {
	PeerBase
}

type PeerResponse struct {
	PeerBase
	ID        uint      `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`
}

type APIError struct {
	ErrorCode    int
	ErrorMessage string
}
