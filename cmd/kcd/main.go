package main

import (
	"github.com/tendermint/tendermint/libs/cli"

	app "github.com/cosmos/sdk-tutorials/hellochain"
	"github.com/cosmos/sdk-tutorials/hellochain/starter"
)

func main() {

	params := starter.NewServerCommandParams(
		"kcd",                  // name of the command
		"kennychain AppDaemon", // description
		starter.NewAppCreator(app.NewHelloChainApp),  // method for constructing an app
		starter.NewAppExporter(app.NewHelloChainApp), // method for exporting chain state
	)

	serverCmd := starter.NewServerCommand(params)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(serverCmd, "KC", starter.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
