package cmd

import (
	"github.com/spf13/cobra"
)

var name string

func init() {
	rootCmd.AddCommand(clusterCmd)

	clusterCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "The name of the cluster. (required)")
	clusterCmd.MarkPersistentFlagRequired("name")
}

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "cluster",
	Long:  `cluster`,
}
