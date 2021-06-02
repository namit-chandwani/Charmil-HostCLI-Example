package main

import (
	"log"

	subcmd "github.com/namit-chandwani/Charmil-HostCLI-Example/pkg/cmd"
	"github.com/namit-chandwani/Charmil-SDK-POC/pkg/core"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

const hostKey = "host"

func main() {
	cmd := &cobra.Command{
		Use:          "host",
		Short:        "Host commands",
		SilenceUsage: true,
	}

	hostCLI := core.GetCLI(hostKey)
	err := hostCLI.AddCommands(cmd, subcmd.FooCmd, subcmd.BarCmd)
	if err != nil {
		log.Fatal(err)
	}

	err = core.LoadCommands(cmd)
	if err != nil {
		log.Fatal(err)
	}

	err = doc.GenMarkdownTree(cmd, "./docs/commands")
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
