package core

import (
	"github.com/MeenaAlfons/ez/adapters/minikube"
)

//
type Infrastructure interface {
	Install() bool
}

//
func MakeInfrastructure(driver string) Infrastructure {
	infrastructureObj := minikube.MakeInfrastructure(driver)

	return infrastructureObj
}
