package main

import (
	"fmt"
	"os"

	"github.com/allora-network/allora-chain/app"
	"github.com/allora-network/allora-chain/app/params"
	"github.com/allora-network/allora-chain/cmd/allorad/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	params.InitSDKConfig()

	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
