package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"google.golang.org/grpc"

	"github.com/iotexproject/iotex-core/pkg/log"
)

func main() {
	conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
	if err != nil {
		log.L().Fatal("failed to connect GRPC!")
	}
	defer conn.Close()

	height, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		log.L().Fatal("failed to parse height!")
	}

	res, err := iotexapi.NewAPIServiceClient(conn).GetTransactionLogByBlockHeight(
		context.Background(),
		&iotexapi.GetTransactionLogByBlockHeightRequest{
			BlockHeight: height,
		},
	)
	if err != nil {
		log.L().Fatal("failed to get log!")
	}

	fmt.Printf("%v\n", res.TransactionLogs)
	return
}
