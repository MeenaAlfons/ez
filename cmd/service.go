package cmd

import (
	"github.com/spf13/cobra"
)

var serviceName string
var clusterName string

func init() {
	rootCmd.AddCommand(serviceCmd)

	serviceCmd.PersistentFlags().StringVarP(&serviceName, "name", "n", "", "The name of the service. (required)")
	serviceCmd.PersistentFlags().StringVarP(&clusterName, "clusterName", "c", "", "The name of the cluster. (required)")
	serviceCmd.MarkPersistentFlagRequired("name")
	serviceCmd.MarkPersistentFlagRequired("clusterName")
}

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "service",
	Long:  `service`,
}
