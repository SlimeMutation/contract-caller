package main

import (
	"context"

	"github.com/urfave/cli/v2"

	"github.com/ethereum/go-ethereum/log"

	contract_caller "github.com/SlimeMutation/contract-caller"
	"github.com/SlimeMutation/contract-caller/common/cliapp"
	"github.com/SlimeMutation/contract-caller/config"
	flag2 "github.com/SlimeMutation/contract-caller/flags"
)

func runContractCaller(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Info("run dapplink vrf")
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		log.Error("failed to load config", "err", err)
		return nil, err
	}
	return contract_caller.NewContractCaller(ctx.Context, &cfg, shutdown)
}

func NewCli() *cli.App {
	flags := flag2.Flags
	return &cli.App{
		Version:              "v0.0.1",
		Description:          "An indexer of all optimism events with a serving api layer",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "caller",
				Flags:       flags,
				Description: "Run ContractCaller",
				Action:      cliapp.LifecycleCmd(runContractCaller),
			},
			{
				Name:        "version",
				Description: "print version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}
