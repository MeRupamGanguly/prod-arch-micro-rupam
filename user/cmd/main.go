package main

import (
	"log"
	"rupx/user/repository"
	"rupx/user/usecase"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	url := "host=localhost user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewUserRepo(db)
	usecase.NewUserSvc(repo)
}
