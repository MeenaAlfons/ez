package core

import (
	"github.com/MeenaAlfons/ez/adapters/minikube"
)

//
type Service interface {
	GetName() string
	GetClusterName() string
	GetStatus() string
	Start() bool
}

//
func MakeService(clusterName, serviceName, image string, replicas int) Service {
	service := minikube.MakeService(clusterName, serviceName, image, replicas)

	return service
}
