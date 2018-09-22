package drivers

import (
	"bytes"
	"fmt"
	"html/template"
	"github.com/MeenaAlfons/ez/data"
	"github.com/MeenaAlfons/ez/utils"
)

//
type Kubectl interface {
	UseContext(profile string) bool
	Deploy(profile, serviceName, imageName string, replicas int) bool
}

type kubectl struct {
	test int
}

//
func MakeKubectl() Kubectl {
	k := new(kubectl)

	return k
}

//
func (k *kubectl) UseContext(profile string) bool {
	useContextCmd := "kubectl config use-context " + profile
	return utils.RunInShellWithStdout(useContextCmd)
}

//
func (k *kubectl) Deploy(profile, serviceName, imageName string, replicas int) bool {

	fmt.Println("Start Deploy")

	deployment := getDeploymentOfService(serviceName, imageName, replicas)

	return utils.RunInShellWithStdoutAndStdin("kubectl --context "+profile+" apply -f -", deployment)
}

func getDeploymentOfService(serviceName, imageName string, replicas int) string {

	// kubectl
	type ServiceT struct {
		ServiceName string
		Image       string
		Replicas    int
	}
	service := ServiceT{serviceName, imageName, replicas}

	fileString, err := data.ReadAsset("/deploymentTemplate.yaml")
	tmpl, err := template.New("deploymentTemplate").Parse(fileString)
	if err != nil {
		panic(err)
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, service)

	return b.String()
}
