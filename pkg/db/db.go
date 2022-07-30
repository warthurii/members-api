package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"members-api/pkg/models"
)

func Init() *gorm.DB {
    dbURL := "postgres://prbc:iloverpw@localhost:5432/prbc_members"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Member{})

    return db
}