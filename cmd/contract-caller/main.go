package main

import (
	"context"
	"github.com/SlimeMutation/contract-caller/common/opio"
	"github.com/ethereum/go-ethereum/log"
	"os"
)

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))
	app := NewCli()
	ctx := opio.WithInterruptBlocker(context.Background())
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Error("Application failed", err)
		os.Exit(1)
	}
}
