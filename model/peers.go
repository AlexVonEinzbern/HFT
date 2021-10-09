package model

import "gorm.io/gorm"

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
