package cmd

import (
	"MatrixAI-Client/chain"
	"MatrixAI-Client/chain/pallets"
	"MatrixAI-Client/chain/subscribe"
	"MatrixAI-Client/config"
	"MatrixAI-Client/logs"
	"MatrixAI-Client/machine_info"
	"MatrixAI-Client/pattern"
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
	"github.com/urfave/cli"
	"time"
)

var ClientCommand = cli.Command{
	Name:  "client",
	Usage: "Starting or terminating a client.",
	Subcommands: []cli.Command{
		{
			Name:  "execute",
			Usage: "Upload hardware configuration and initiate listening events.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "mnemonic, m",
					Required: true,
					Usage:    "Mnemonics used to complete transactions",
				},
			},
			Action: func(c *cli.Context) error {
				matrixWrapper, hwInfo, chainInfo, err := getMatrix(c)
				if err != nil {
					logs.Error(err.Error())
					return err
				}

				machine, err := matrixWrapper.GetMachine(*hwInfo)
				if err != nil {
					logs.Error(fmt.Sprintf("Error: %v", err))
					return err
				}

				if machine.Metadata == "" {
					logs.Normal("Machine does not exist")
					hash, err := matrixWrapper.AddMachine(*hwInfo)
					if err != nil {
						logs.Error(fmt.Sprintf("Error block : %v, msg : %v\n", hash, err))
						return err
					}
				} else {
					logs.Normal("Machine already exists")
				}

				for {
					subscribeBlocks := subscribe.NewSubscribeWrapper(chainInfo)
					orderId, orderPlacedMetadata, err := subscribeBlocks.SubscribeEvents(hwInfo)
					if err != nil {
						logs.Error(err.Error())
						return err
					}
					logs.Normal("subscribe done")

					if codec.Eq(orderId, pattern.OrderId{}) {
						logs.Result("Stop the client.")
						return nil
					}

					// ------- Simulate AI model training -------
					time.Sleep(10 * time.Second)
					// ------- Simulate AI model training -------

					_, err = matrixWrapper.OrderCompleted(orderId, orderPlacedMetadata)
					if err != nil {
						logs.Error(err.Error())
						return err
					}

					logs.Normal("OrderCompleted done")

					time.Sleep(1 * time.Second)
				}
			},
		},
		{
			Name:  "stop",
			Usage: "Stop the client.",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "mnemonic, m",
					Required: true,
					Usage:    "Mnemonics used to complete transactions",
				},
			},
			Action: func(c *cli.Context) error {
				matrixWrapper, hwInfo, _, err := getMatrix(c)
				if err != nil {
					logs.Error(err.Error())
					return err
				}

				hash, err := matrixWrapper.RemoveMachine(*hwInfo)
				if err != nil {
					logs.Error(fmt.Sprintf("Error block : %v, msg : %v\n", hash, err))
					return err
				}
				return nil
			},
		},
	},
}

func getMatrix(c *cli.Context) (*pallets.WrapperMatrix, *machine_info.MachineInfo, *chain.InfoChain, error) {
	logs.Result("-------------------- start --------------------")

	mnemonic := c.String("mnemonic")

	hwInfo, err := machine_info.GetMachineInfo()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error getting hardware info: %v", err)
	}
	logs.Normal(fmt.Sprintf("Hardware Info:\n%+v\n", hwInfo))

	newConfig := config.NewConfig(
		mnemonic,
		pattern.RPC,
		1)

	chainInfo, err := chain.GetChainInfo(newConfig)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error getting chain info: %v", err)
	}

	return pallets.NewMatrixWrapper(chainInfo), &hwInfo, chainInfo, nil
}
