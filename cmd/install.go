package cmd

import (
	"fmt"

	"github.com/MeenaAlfons/ez/core"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install",
	Long:  `install`,
	Run: func(cmd *cobra.Command, args []string) {

		driver := "minikube"

		fmt.Println("Installing necessary drivers: " + driver)

		infrastructure := core.MakeInfrastructure(driver)

		installed := infrastructure.Install()
		if !installed {
			fmt.Println("Failed to install necessary drivers: " + driver)
		}
	},
}
