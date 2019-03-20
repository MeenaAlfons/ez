# ez

Never been easier to deploy and test your microservices application localling. 

# Usage

To be continued ...

# Build

`go build`

# To Be Done

- installation
- install specific version of minikube and kubernetes

# In Progress

The following are in progress work.

```
go run main.go cluster stop -n tell
go run main.go service set -n myservice -i myimage -r 5 -c mycluster

minikube stop --profile dev

@FOR /f "tokens=*" %i IN ('minikube docker-env --profile dodo') DO @%i

  

  -----------------------------------------

  docker save -o E:\Temp\MyAppName179497923imageName.tar mymain:1.0 && @FOR /f "to
kens=*" %i IN ('minikube docker-env --profile dody') DO @%i && docker load -i E:
\Temp\MyAppName179497923imageName.tar

cmd /c docker save -o E:\Temp\MyAppName179497923imageName.tar mymain:1.0 && @FOR /f "tokens=*" %i IN ('minikube docker-env --profile dody') DO @%i && docker load -i E:\Temp\MyAppName179497923imageName.tar



docker --host tcp://192.168.99.104:2376 save mymain:2.0 | docker --host tcp://192.168.99.105:2376 load


docker --tlscacert D:\minikube\home0\.minikube\certs\ca.pem --tlscert D:\minikube\home0\.minikube\certs\cert.pem --tlskey D:\minikube\home0\.minikube\certs\key.pem --tlsverify -H tcp://192.168.99.104:2376 save mymain | pv -N "Transfering..." | docker --tlscacert D:\minikube\home0\.minikube\certs\ca.pem --tlscert D:\minikube\home0\.minikube\certs\cert.pem --tlskey D:\minikube\home0\.minikube\certs\key.pem --tlsverify -H tcp://192.168.99.105:2376 load


docker --tlscacert D:\minikube\home0\.minikube\certs\ca.pem --tlscert D:\minikube\home0\.minikube\certs\cert.pem --tlskey D:\minikube\home0\.minikube\certs\key.pem --tlsverify -H tcp://192.168.99.104:2376 save mymain | pv -N "Transfering..." > E:/Temp/test.tar

main.go docker -c cluster build <...>
main.go docker -c cluster upload image
main.go service set -n .. -c .. -i .. -r
// Check if the image exists in the cluster. If not move it to the cluster.




minikube dashboard --url --profile dody
kubectl config use-context dody
```