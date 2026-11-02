package main

import (
	"flag"
	"fmt"
	"log"

	"smart-pharmacy/internal/config"
	"smart-pharmacy/internal/node"
	"smart-pharmacy/internal/server"
)

func main() {
	cfgPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*cfgPath)
	if err != nil {
		log.Fatalf("load config %q: %v", *cfgPath, err)
	}

	fmt.Printf("Starting node %s as %s\n", cfg.Node.ID, cfg.Node.Role)

	switch cfg.Node.Role {
	case "central":
		server.Start(cfg)
	case "edge":
		node.Start(cfg)
	default:
		log.Fatalf("unknown node role %q — must be 'central' or 'edge'", cfg.Node.Role)
	}
}
