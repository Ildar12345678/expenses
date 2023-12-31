package main

import (
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"expenses/internal/rest"
	"log"
)

type config struct {
	port int
	env  string
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|stagging|production)")
	flag.Parse()
	
	// logger, err := l.NewLog()
	// if err != nil {
	// 	log.Fatal("error while creating logger:", err)
	// }
	// logger := log.New(os.Stdout, "", log.LstdFlags|log.Llongfile)
	// logger.Printf("starting %s server on %d", cfg.env, cfg.port)
	// defer logger.Printf("stoped %s server on %d", cfg.env, cfg.port)
	
	err := rest.RunAPI(fmt.Sprintf(":%d", cfg.port))
	if err != nil {
		log.Fatal("fatal error", err)
	}
}
