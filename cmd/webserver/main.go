package main

import (
	"net/http"
	"os"
	"time"

	"github.com/daniellcas/streaming-video-mp4/cmd/worker"
	"github.com/daniellcas/streaming-video-mp4/internal/config"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg := loadConfig()

	// Processar video mp4 em hls
	go worker.Execute(cfg)

	s := http.NewServeMux()
	s.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	err := http.ListenAndServe(":"+cfg.Port, s)
	if err != nil {
		log.Fatal().Err(err).Msg("Error em subir o servidor")
	}
}

func loadConfig() *config.Config {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime})

	// Carrega as variaveis de ambiente do .env
	if err := godotenv.Load(".env"); err != nil {
		log.Info().Err(err).Msg("Arquivo .env não carregado")
	}

	// COnfigura as variaveis para a struct
	cfg, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("Falhou em configurar o ambiente")
	}
	return cfg
}
