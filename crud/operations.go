package crud

import (
	"hft/db"
	"hft/model"
	"log"
)

func CreatePeer(cPeer model.PeerCreate) model.Peer {
	peer := model.Peer{PeerBase: model.PeerBase{Host: cPeer.Host}}
	conn := db.Connection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	result := conn.Create(&peer)
	if result.Error != nil{
		log.Fatal(result.Error)
	}

	return peer
}

func SearchAllPeers(name string) []model.Peer {
	var peers []model.Peer
	conditions := make(map[string]interface{})
	conn := db.Connection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	if name != "" {
		conditions["name"] = name
	}

	result := conn.Where(conditions).Find(&peers)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return peers
}

func GetPeers(file string) []model.Peer {
	var peers []model.Peer
	conn := db.Connection()
	sqlconn, err := conn.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlconn.Close()

	result := conn.Where("file <> ?", file).Find(&peers)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return peers

}
