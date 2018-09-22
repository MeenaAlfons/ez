package cmd

import (
	"fmt"
	"github.com/MeenaAlfons/ez/core"

	"github.com/spf13/cobra"
)

func init() {
	clusterCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a cluster",
	Long:  `This command is used to stop a running cluster`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO Make some validations

		fmt.Println("Stopping cluster " + name + " ...")

		cluster := core.MakeCluster(name)

		started := cluster.Stop()
		if !started {
			fmt.Println("Failed to stop cluster " + cluster.GetName() + ".")
		}

		// TODO Status feature is not implemented yet.
		// fmt.Println("Cluster " + cluster.GetName() + ", Status:" + cluster.GetStatus() + ".")
	},
}
