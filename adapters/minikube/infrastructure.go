package minikube

import (
	"github.com/MeenaAlfons/ez/drivers"
)

//
type Infrastructure struct {
	driver   string
	minikube drivers.Minikube
}

//
func MakeInfrastructure(driver string) *Infrastructure {
	infrastructureObj := new(Infrastructure)
	infrastructureObj.driver = driver
	infrastructureObj.minikube = drivers.MakeMinikube()

	return infrastructureObj
}

//
func (i *Infrastructure) Install() bool {

	ok := i.minikube.Install()

	return ok
}
