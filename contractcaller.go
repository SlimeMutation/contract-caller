package contract_caller

import (
	"context"
	"math/big"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	common2 "github.com/SlimeMutation/contract-caller/common"
	"github.com/SlimeMutation/contract-caller/config"
	"github.com/SlimeMutation/contract-caller/driver"
	"github.com/SlimeMutation/contract-caller/worker"
)

type ContractCaller struct {
	worker *worker.Worker

	shutdown context.CancelCauseFunc
	stopped  atomic.Bool
}

func NewContractCaller(ctx context.Context, cfg *config.Config, shutdown context.CancelCauseFunc) (*ContractCaller, error) {
	ethcli, err := driver.EthClientWithTimeout(ctx, cfg.Chain.ChainRpcUrl)
	if err != nil {
		log.Error("new eth cli fail", "err", err)
		return nil, err
	}

	callerPrivateKey, _, err := common2.ParseWalletPrivKeyAndContractAddr(
		"ContractCaller", cfg.Chain.Mnemonic, cfg.Chain.CallerHDPath,
		cfg.Chain.PrivateKey, cfg.Chain.DappLinkVrfContractAddress, cfg.Chain.Passphrase,
	)

	decg := &driver.DriverEingineConfig{
		ChainClient:               ethcli,
		ChainId:                   big.NewInt(int64(cfg.Chain.ChainId)),
		DappLinkVrfAddress:        common.HexToAddress(cfg.Chain.DappLinkVrfContractAddress),
		CallerAddress:             common.HexToAddress(cfg.Chain.CallerAddress),
		PrivateKey:                callerPrivateKey,
		NumConfirmations:          cfg.Chain.Confirmations,
		SafeAbortNonceTooLowCount: cfg.Chain.SafeAbortNonceTooLowCount,
	}
	eingine, err := driver.NewDriverEingine(ctx, decg)
	if err != nil {
		log.Error("new driver eingine fail", "err", err)
		return nil, err
	}

	workerConfig := &worker.WorkerConfig{
		LoopInterval: cfg.Chain.CallInterval,
	}

	workerProcessor, err := worker.NewWorker(eingine, workerConfig, shutdown)
	if err != nil {
		log.Error("new event processor fail", "err", err)
		return nil, err
	}

	return &ContractCaller{
		worker:   workerProcessor,
		shutdown: shutdown,
	}, nil
}

func (dvrf *ContractCaller) Start(ctx context.Context) error {
	err := dvrf.worker.Start()
	if err != nil {
		return err
	}
	return nil
}

func (dvrf *ContractCaller) Stop(ctx context.Context) error {
	err := dvrf.worker.Close()
	if err != nil {
		return err
	}
	return nil
}

func (dvrf *ContractCaller) Stopped() bool {
	return dvrf.stopped.Load()
}
