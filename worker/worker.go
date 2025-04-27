package worker

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/SlimeMutation/contract-caller/common/tasks"
	"github.com/SlimeMutation/contract-caller/driver"
)

type WorkerConfig struct {
	LoopInterval time.Duration
}

type Worker struct {
	workerConfig *WorkerConfig

	deg *driver.DriverEingine

	resourceCtx    context.Context
	resourceCancel context.CancelFunc
	tasks          tasks.Group
}

func NewWorker(deg *driver.DriverEingine, workerConfig *WorkerConfig, shutdown context.CancelCauseFunc) (*Worker, error) {
	resCtx, resCancel := context.WithCancel(context.Background())
	return &Worker{
		deg:            deg,
		workerConfig:   workerConfig,
		resourceCtx:    resCtx,
		resourceCancel: resCancel,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in bridge processor: %w", err))
		}},
	}, nil
}

func (wk *Worker) Start() error {
	log.Info("starting worker processor...")
	tickerEventWorker := time.NewTicker(wk.workerConfig.LoopInterval)
	wk.tasks.Go(func() error {
		for range tickerEventWorker.C {
			log.Info("start handler random for vrf")
			err := wk.ProcessCallerVrf()
			if err != nil {
				log.Error("process caller vrf fail", "err", err)
				return err
			}
		}
		return nil
	})
	return nil
}

func (wk *Worker) Close() error {
	wk.resourceCancel()
	return wk.tasks.Wait()
}

func (wk *Worker) ProcessCallerVrf() error {
	randomList := make([]*big.Int, 0)
	for i := 0; i < 3; i++ {
		randomList = append(randomList, big.NewInt(int64(i)))
	}
	txReceipt, err := wk.deg.FulfillRandomWords(big.NewInt(1), randomList)
	if err != nil {
		log.Error("failed to fulfill random words", "err", err)
		return err
	}
	if txReceipt.Status == 1 {
		log.Info("successfully fulfill random words")
	}
	return nil
}
