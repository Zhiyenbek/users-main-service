package app

import (
	"log"

	"github.com/Zhiyenbek/users-auth-service/config"
)

func Run() {
	cfg, err := config.New()
	if err != nil {
		log.Println(err)
	}

	log.Println(cfg.App.Host)
	log.Println(cfg.App.Port)
}
