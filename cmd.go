// Package cmd defines the cobra command structs and an execution method for adding an improved CLI to
// KrakenD based api gateways
package cmd

import (
	"fmt"
	"os"

	"github.com/devopsfaith/krakend/config"
)

// Executor defines the function that requires a service description
type Executor func(config.ServiceConfig)

// Execute sets up the cmd package with the received configuration parser and executor and delegates
// the CLI execution to the cobra lib
func Execute(configParser config.Parser, f Executor) {
	parser = configParser
	run = f
	eulas = make([]bool, len(EulaFlags))
	for i, eula := range EulaFlags {
		runCmd.PersistentFlags().BoolVarP(&eulas[i], eula.FlagLong, eula.FlagShort, false, eula.Description)
	}
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
