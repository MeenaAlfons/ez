package cmd

import (
	"fmt"

	"github.com/MeenaAlfons/ez/consts"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + consts.AppName,
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(consts.AppName + " " + consts.Version)
	},
}
