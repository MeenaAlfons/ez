package minikube

import (
	"github.com/MeenaAlfons/ez/drivers"
)

//
type Cluster struct {
	name     string
	minikube drivers.Minikube
}

//
func MakeCluster(name string) *Cluster {
	clusterObj := new(Cluster)
	clusterObj.name = name
	clusterObj.minikube = drivers.MakeMinikube()

	return clusterObj
}

//
func (c *Cluster) GetName() string {
	return c.name
}

//
func (c *Cluster) GetStatus() string {
	return c.minikube.Status(c.name)
}

//
func (c *Cluster) Start() bool {

	ok := c.minikube.Start(c.name)

	return ok
}

//
func (c *Cluster) Stop() bool {

	ok := c.minikube.Stop(c.name)

	return ok
}
