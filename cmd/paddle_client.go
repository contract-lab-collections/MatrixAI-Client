package cmd

import (
	client2 "MatrixAI-Client/deep_learning_model/paddlepaddle/client"
	"MatrixAI-Client/logs"
	"context"
	"fmt"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

var PaddleCommand = cli.Command{
	Name:  "paddle",
	Usage: "Paddlepaddle client.",
	Subcommands: []cli.Command{
		{
			Name:  "start",
			Usage: "Start paddlepaddle client.",
			Action: func(c *cli.Context) error {

				logs.Result("start paddle client")

				conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
				if err != nil {
					return err
				}
				defer func(conn *grpc.ClientConn) {
					err := conn.Close()
					if err != nil {

					}
				}(conn)

				client := client2.NewTrainServiceClient(conn)

				req := &client2.Empty{}
				res, err := client.TrainAndPredict(context.Background(), req)
				if err != nil {
					return err
				}

				imgData := res.GetImageData()
				trueLabel := res.GetTrueLabel()
				predLabel := res.GetPredictedLabel()

				logs.Normal(fmt.Sprintf("imgData: %v", imgData))
				logs.Normal(fmt.Sprintf("trueLabel: %v", trueLabel))
				logs.Normal(fmt.Sprintf("predLabel: %v", predLabel))

				return nil
			},
		},
	},
}
