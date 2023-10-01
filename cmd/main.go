package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"pet/internal/port"
	"pet/internal/repository"
	"pet/internal/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init the config: %s", err.Error())
	}
	fmt.Println(viper.GetString("db.sslmode"))
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password")})
	if err != nil {
		log.Fatalf("Can't connect to DB: %s", err.Error())
		return
	}
	newRepository := repository.NewRepository(db)
	newService := service.NewService(newRepository)
	server := port.NewServer(viper.GetString("port"), newService)

	err = server.Listen()
	if err != nil {
		log.Fatalf("Server ended with this error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
