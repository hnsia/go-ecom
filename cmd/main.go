package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/hnsia/go-ecom/cmd/api"
	"github.com/hnsia/go-ecom/config"
	"github.com/hnsia/go-ecom/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}