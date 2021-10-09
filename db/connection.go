package db

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "hft/model"
)

func Connection() *gorm.DB {
    dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
    if err != nil {
        log.Fatal(err)
    }
    return db
}

func Migrate () {
    db := Connection()
    sqlDB, err := db.DB()

    if err != nil {
        log.Fatal(err)
    }

    defer sqlDB.Close()

    db.AutoMigrate(&model.Peer{})
}