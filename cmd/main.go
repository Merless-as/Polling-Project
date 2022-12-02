package main

import (
	"RestApiProj"
	"RestApiProj/pkg/handler"
	"RestApiProj/pkg/repository"
	"RestApiProj/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalln(err)
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalln(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(RestApiProj.Server)
	if err = srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}