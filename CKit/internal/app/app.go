package app

import (
	"CKit/internal/app/config"
	"CKit/internal/app/connections"
	"CKit/internal/app/start"
	"fmt"
)

func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("error during config setup")
	}
	fmt.Printf("HTTP server config: %v\n", cfg.HTTPServer)
	fmt.Printf("DB server config: %v\n", cfg.DB)

	conn, err := connections.New(cfg)
	if err != nil {
		fmt.Printf("error during connections setup")
	}

	start.HTTP(conn, &cfg.HTTPServer)
}
