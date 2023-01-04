package main

import (
	"log"

	"app/config"
	"app/pkg/db"
)

func main() {

	cfg := config.Load()

	conn, err := db.NewConnectPostgres(cfg)
	if err != nil {
		log.Fatal("error databse connection", err.Error())
	}

	query := `
		insert into movie (
			id,
			title,
			duration,
			description
		) values ($1, $2, $3, $4)
	`

	_, err = conn.Exec(query,
		"f904d5c9-3e24-4637-b0d7-b3d103d2011c",
		"Avatar 2",
		"03:00:00",
		"Best movie",
	)

	if err != nil {
		log.Println("error whiling insert movie: ", err)
		return
	}

}
