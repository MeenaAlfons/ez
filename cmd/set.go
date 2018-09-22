package cmd

import (
	"fmt"
	"github.com/MeenaAlfons/ez/core"

	"github.com/spf13/cobra"
)

var replicas int
var image string

func init() {
	serviceCmd.AddCommand(setCmd)

	setCmd.Flags().IntVarP(&replicas, "replicas", "r", 0, "The number of replicas.")
	setCmd.Flags().StringVarP(&image, "image", "i", "", "Docker image for the service")
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Setup a new service or update an existing service in a cluster",
	Long:  `Setup a new service or update an existing service in a cluster`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO Make some validations

		fmt.Println("Setting service " + serviceName + " in cluseter " + clusterName)

		service := core.MakeService(clusterName, serviceName, image, replicas)

		started := service.Start()
		if !started {
			fmt.Println("Failed to start service " + service.GetName() + " in cluster " + service.GetClusterName() + ".")
		}

		// TODO Status feature is not implemented yet.
		// fmt.Println("Service " + service.GetName() + " in cluster " + service.GetClusterName() + ", Status:" + service.GetStatus() + ".")
	},
}
