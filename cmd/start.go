package cmd

import (
	"fmt"

	"github.com/MeenaAlfons/ez/core"

	"github.com/spf13/cobra"
)

func init() {
	clusterCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a cluster",
	Long:  `This command is used to start a new cluster or a stopped cluster`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO Make some validations

		fmt.Println("Starting cluster " + name + " ...")

		cluster := core.MakeCluster(name)

		started := cluster.Start()
		if !started {
			fmt.Println("Failed to start cluster " + cluster.GetName() + ".")
		}

		// TODO Status feature is not implemented yet.
		// fmt.Println("Cluster " + cluster.GetName() + ", Status:" + cluster.GetStatus() + ".")
	},
}
