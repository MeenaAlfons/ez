package core

import (
	"github.com/MeenaAlfons/ez/adapters/minikube"
)

//
type Cluster interface {
	GetName() string
	GetStatus() string
	Start() bool
	Stop() bool
}

//
func MakeCluster(name string) Cluster {
	clusterObj := minikube.MakeCluster(name)

	return clusterObj
}
