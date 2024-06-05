package worker

import (
	"context"
	"log"
	"os"
	"os/exec"

	"github.com/daniellcas/streaming-video-mp4/internal/config"
	"github.com/daniellcas/streaming-video-mp4/internal/worker"
)

func CreateOutputDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func Execute(cfg *config.Config) {
	ctx := context.Background()
	outputPath := cfg.OutputDir
	if err := CreateOutputDir(outputPath); err != nil {
		log.Fatal("Failed to create output directory: " + cfg.OutputDir)
	}

	args := worker.BuildCommand(cfg)
	cmd := exec.CommandContext(ctx, "ffmpeg", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
