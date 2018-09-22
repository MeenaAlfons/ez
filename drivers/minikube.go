package drivers

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/MeenaAlfons/ez/consts"
	"github.com/MeenaAlfons/ez/utils"
)

//
type Minikube interface {
	Start(profile string) bool
	Stop(profile string) bool
	Status(profile string) string
	StartService(profile, serviceName, imageName string, replicas int) bool
	ServiceInfo(profile, serviceName string) (exists bool, status, image string, replicas int)
}

type minikube struct {
	kubectl Kubectl
	test    int
}

//
func MakeMinikube() Minikube {
	m := new(minikube)
	m.kubectl = MakeKubectl()

	return m
}

//
func (m *minikube) Start(profile string) bool {
	// TODO Make validations on the profile string

	// out, err := exec.Command("minikube", "start", "--profile", profile).CombinedOutput()

	ok := utils.RunCommandWithStdout("minikube", "start", "--profile", profile)

	if !ok {
		fmt.Println("minikube failed to start a cluster with profile " + profile)
		return false
	}

	fmt.Println("minikube succeeded to start a cluster with profile " + profile)
	return true
}

//
func (m *minikube) Stop(profile string) bool {
	// TODO Make validations on the profile string

	out, err := exec.Command("minikube", "stop", "--profile", profile).CombinedOutput()
	if err != nil {
		fmt.Println("minikube failed to stop a cluster with profile " + profile)
		fmt.Println("The following is the output of minikube:")
		fmt.Println(string(out))

		return false
	}

	fmt.Println("minikube succeeded to stop a cluster with profile " + profile)
	fmt.Println("The following is the output of minikube:")
	fmt.Println(string(out))

	return true
}

//
func (m *minikube) Status(profile string) string {

	statusCmd := "minikube status --profile " + profile
	ok, out := utils.RunInShellWithStdoutAndCapture(statusCmd)
	if !ok {
		return "Error happend while retrieving status"
	}

	// TODO Parse output and return a meaningful status
	return out
}

//
func (m *minikube) Clean(profile string) bool {
	return false
}

//
func (m *minikube) StartService(profile, serviceName, imageName string, replicas int) bool {

	toHostIP := getIp(profile)

	fmt.Println("Uploading image " + imageName + " to cluster " + profile + " ...")
	ok := moveDockerImage(imageName, toHostIP)
	if ok {
		fmt.Println("Done uploading image to cluster")
	} else {
		fmt.Println("Error while uploading image to cluster")
		return false
	}

	fmt.Println("Deploying service " + serviceName + " using image " + imageName)
	ok = m.kubectl.Deploy(profile, serviceName, imageName, replicas)
	if ok {
		fmt.Println("Done deploying service " + serviceName + " using image " + imageName)
	} else {
		fmt.Println("Error while deploying service " + serviceName + " using image " + imageName)
		return false
	}

	return ok
}

//
func (m *minikube) ServiceInfo(profile, serviceName string) (exists bool, status, image string, replicas int) {
	exists = true
	status = ""
	image = ""
	replicas = 1

	return
}

func moveDockerImage(imageName, toHostIp string) bool {
	toHost := "tcp://" + toHostIp + ":2376"

	// TODO get Minikube home
	minikubeHome := "D:/minikube/home0"
	certsDir := minikubeHome + "/.minikube/certs"
	toDocker := MakeDocker(toHost, certsDir)

	uploadImageToClusterCmd := DockerSaveCmd(imageName) + " | " + toDocker.DockerLoadCmd()

	// fmt.Println("Command: " + uploadImageToClusterCmd)
	// TODO this part need progress indicator
	return utils.RunInShellWithStdout(uploadImageToClusterCmd)
}

func moveDockerImageOld(imageName, toProfile string) bool {

	prevDir := "E:/Temp/"
	// prevDir := ""

	tempDir, err := ioutil.TempDir(prevDir, consts.AppName)
	if err != nil {
		fmt.Println("Cannot create temporary directory " + tempDir)
		return false
	}
	tarPath := tempDir + "_" + imageName + ".tar"

	saveCmd := "docker save -o " + tarPath + " " + imageName
	loadCmd := "docker load -i " + tarPath
	changeEnvCmd := utils.GetEvalCmd("minikube docker-env --profile " + toProfile)
	uploadImageToClusterCmd := saveCmd + " && (" + changeEnvCmd + ") && " + loadCmd

	// fmt.Println("Command: " + uploadImageToClusterCmd)
	ok := utils.RunInShellWithStdout(uploadImageToClusterCmd)

	os.RemoveAll(tempDir)

	return ok
}

func getIp(profile string) string {
	ip := ""
	out, err := exec.Command("minikube", "ip", "--profile", profile).CombinedOutput()
	if err == nil {
		ip = strings.Trim(string(out), " \r\n\t")
	}
	return ip
}
