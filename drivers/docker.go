package drivers

type Docker struct {
	host   string
	cacert string
	cert   string
	key    string
}

//
func MakeDocker(host, certsDir string) *Docker {
	d := new(Docker)
	d.host = host
	d.cacert = certsDir + "/ca.pem"
	d.cert = certsDir + "/cert.pem"
	d.key = certsDir + "/key.pem"

	return d
}

func (d *Docker) DockerSaveCmd(image string) string {
	return d.DockerCmd() + " save " + image
}

func (d *Docker) DockerLoadCmd() string {
	return d.DockerCmd() + " load"

}

/**
--tlscacert string   Trust certs signed only by this CA (default "/root/.docker/ca.pem")
--tlscert string     Path to TLS certificate file (default "/root/.docker/cert.pem")
--tlskey string      Path to TLS key file (default "/root/.docker/key.pem")
--tlsverify          Use TLS and verify the remote
*/
func (d *Docker) DockerCmd() string {
	return "docker" + " -H " + d.host + " --tlscacert " + d.cacert + " --tlscert " + d.cert + " --tlskey " + d.key + " --tlsverify"
}

func DockerCmd() string {
	return "docker"
}

func DockerLoadCmd() string {
	return DockerCmd() + " load"
}

func DockerSaveCmd(image string) string {
	return DockerCmd() + " save " + image
}
