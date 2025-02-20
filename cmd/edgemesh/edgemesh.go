package main

import (
	"os"

	"k8s.io/component-base/logs"

	"github.com/kubeedge/edgemesh/cmd/edgemesh/app"
)

func main() {
	command := app.NewEdgeMeshCommand()
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
