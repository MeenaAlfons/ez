package minikube

import (
	"github.com/MeenaAlfons/ez/drivers"
)

//
type Service struct {
	name        string
	clusterName string
	image       string
	replicas    int
	minikube    drivers.Minikube
}

//
func MakeService(clusterName, serviceName, image string, replicas int) *Service {
	serviceObj := new(Service)
	serviceObj.name = serviceName
	serviceObj.clusterName = clusterName
	serviceObj.image = image
	serviceObj.replicas = replicas
	serviceObj.minikube = drivers.MakeMinikube()

	return serviceObj
}

//
func (s *Service) GetName() string {
	return s.name
}

//
func (s *Service) GetClusterName() string {
	return s.clusterName
}

//
func (s *Service) GetStatus() string {
	_, status, _, _ := s.minikube.ServiceInfo(s.clusterName, s.name)

	return status
}

//
func (s *Service) Start() bool {
	// Get current service info
	exists, _, currentImage, currentReplicas := s.minikube.ServiceInfo(s.clusterName, s.name)
	if exists {
		if len(s.image) == 0 {
			s.image = currentImage
		}

		if s.replicas == 0 {
			s.replicas = currentReplicas
		}
	}

	// Set default replicas
	if s.replicas == 0 {
		s.replicas = 1
	}

	ok := s.minikube.StartService(s.clusterName, s.name, s.image, s.replicas)

	return ok
}
