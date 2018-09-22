package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	serviceCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing service in a cluster",
	Long:  `Delete an existing service in a cluster`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO Make some validations

		fmt.Println("Deleting service " + serviceName + " in cluseter " + clusterName + " ...")

		fmt.Println("Deleting service is not implemented yet")
	},
}
